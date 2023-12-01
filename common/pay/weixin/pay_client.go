package weixin

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/textproto"
	"strconv"
	"time"

	"github.com/agui-coder/simple-admin-pay-api/common/pay/model"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xpem"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

// ClientConfig API 版本 - V3协议说明  https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay-1.shtml
// ClientConfig 实现了 pay.ClientConfig 接口
type ClientConfig struct {
	AppId             string `json:"appId" validate:"required"`             //appId
	MchId             string `json:"mchId" validate:"required"`             //商户ID 或者服务商模式的 sp_mchid
	PrivateKeyContent string `json:"privateKeyContent" validate:"required"` //apiclient_key.pem 证书文件的对应字符串
	SerialNumber      string `json:"serialNumber" validate:"required"`      //apiclient_cert.pem 证书文件的证书号
	ApiV3Key          string `json:"apiV3Key" validate:"required"`          //apiV3Key，商户平台获取
}

func (p ClientConfig) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(p)
	return err
}

type Client struct {
	ChannelId uint64
	Config    *ClientConfig
	client    *wechat.ClientV3
}

func (w *Client) Init() error {
	client, err := wechat.NewClientV3(w.Config.MchId, w.Config.SerialNumber, w.Config.ApiV3Key, w.Config.PrivateKeyContent)
	if err != nil {
		logx.Error(err)
		return err
	}
	//certs, err := wechat.GetPlatformCerts(context.Background(), w.Config.MchId, w.Config.ApiV3Key, w.Config.SerialNumber, w.Config.PrivateKeyContent, wechat.CertTypeALL)
	//if certs.Code == wechat.Success && len(certs.Certs) > 0 {
	//	client.SetPlatformCert([]byte(certs.Certs[0].PublicKey), w.Config.SerialNumber)
	//} else {
	//	logx.Error("certs:%s", certs.Error)
	//	return errorx.NewInvalidArgumentError(certs.Error)
	//}
	// 启用自动同步返回验签，并定时更新微信平台API证书（开启自动验签时，无需单独设置微信平台API证书和序列号）
	err = client.AutoVerifySign()
	if err != nil {
		logx.Error(err)
		return err
	}
	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//client.SetBodySize() // 没有特殊需求，可忽略此配置
	// 打开Debug开关，输出日志，默认是关闭的
	client.DebugSwitch = gopay.DebugOn
	w.client = client
	logx.Infof("[init][客户端 %d 初始化完成]\n", w.ChannelId)
	return nil
}

func (w *Client) UnifiedOrder(context.Context, model.OrderUnifiedReq) (*model.OrderResp, error) {
	return nil, errorx.NewInvalidArgumentError("该客户端不能下单哦")
}

func (w *Client) GetId() uint64 {
	return w.ChannelId
}

func (w *Client) GetOrder(ctx context.Context, no string) (*model.OrderResp, error) {
	rsp, err := w.client.V3TransactionQueryOrder(ctx, wechat.OutTradeNo, no)
	if err != nil && rsp.Code != 0 {
		return nil, err
	}
	status, err := parseStatus(rsp.Response.TradeState)
	if err != nil {
		return nil, err
	}
	var openid *string
	if rsp.Response.Payer != nil {
		openid = &rsp.Response.Payer.Openid
	}
	successTime, err := parseDate(rsp.Response.SuccessTime)
	if err != nil {
		return nil, err
	}
	return model.Of(status, rsp.Response.TransactionId, openid, successTime, no, rsp.Response), nil
}

func (w *Client) Refresh(config model.ClientConfig) error {
	if *w.Config == config {
		return nil
	}
	logx.Infof("[refresh][客户端 %d 发生变化，重新初始化]", w.ChannelId)
	if clientConfig, ok := config.(ClientConfig); ok {
		w.Config = &clientConfig
	} else {
		return errorx.NewInvalidArgumentError(fmt.Sprintf("客户端%d重新初始化失败", w.ChannelId))
	}
	return nil
}

// buildPayUnifiedOrderBm 通用返回
func (w *Client) buildPayUnifiedOrderBm(req model.OrderUnifiedReq) gopay.BodyMap {
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("appid", w.Config.AppId).
		Set("description", req.Body).
		Set("out_trade_no", req.OutTradeNo).
		Set("time_expire", req.ExpireTime.Format(time.RFC3339)).
		Set("notify_url", req.NotifyUrl).
		SetBodyMap("scene_info", func(bm gopay.BodyMap) {
			bm.Set("payer_client_ip", req.UserIp)
		}).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", req.Price).
				Set("currency", "CNY") //CNY：人民币，境内商户号仅支持人民币。
		})
	return bm
}

func (w *Client) ParseOrderNotify(header map[string][]string, body []byte) (*model.OrderResp, error) {
	notifyReq, err := V3ParseNotify(header, body)
	if err != nil {
		return nil, err
	}
	cert := w.client.WxPublicKey()
	err = notifyReq.VerifySignByPK(cert)
	if err != nil {
		return nil, err
	}
	result, err := notifyReq.DecryptCipherText(string(w.client.ApiV3Key))
	if err != nil {
		return nil, err
	}
	status, err := parseStatus(result.TradeState)
	if err != nil {
		return nil, err
	}
	var openid *string
	if result.Payer != nil {
		openid = &result.Payer.Openid
	}
	successTime, err := parseDate(result.SuccessTime)
	if err != nil {
		return nil, err
	}
	return model.Of(status, result.TransactionId, openid, successTime, result.OutTradeNo, body), nil
}

func (w *Client) UnifiedRefund(ctx context.Context, req model.RefundUnifiedReq) (*model.RefundResp, error) {
	//TODO implement me
	panic("implement me")
}

// V3ParseNotify 解析微信回调请求的参数到 V3NotifyReq 结构体
func V3ParseNotify(header map[string][]string, body []byte) (notifyReq *wechat.V3NotifyReq, err error) {
	mimeHeader := textproto.MIMEHeader(header)
	si := &wechat.SignInfo{
		HeaderTimestamp: mimeHeader.Get(wechat.HeaderTimestamp),
		HeaderNonce:     mimeHeader.Get(wechat.HeaderNonce),
		HeaderSignature: mimeHeader.Get(wechat.HeaderSignature),
		HeaderSerial:    mimeHeader.Get(wechat.HeaderSerial),
		SignBody:        string(body),
	}
	notifyReq = &wechat.V3NotifyReq{SignInfo: si}
	if err = json.Unmarshal(body, notifyReq); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s, %+v)：%w", body, notifyReq, err)
	}
	return notifyReq, nil
}

func parseStatus(tradeState string) (uint8, error) {
	switch tradeState {
	case "NOTPAY":
	case "USERPAYING": // 支付中，等待用户输入密码（条码支付独有）
		return model.WAITING, nil
	case "SUCCESS":
		return model.SUCCESS, nil
	case "REFUND":
		return model.REFUND, nil
	case "CLOSED":
	case "REVOKED": // 已撤销（刷卡支付独有）
	case "PAYERROR": // 支付失败（其它原因，如银行返回失败）
		return model.CLOSED, nil
	default:
		return model.ERROR, errorx.NewInvalidArgumentError(fmt.Sprintf("未知的支付状态%s", tradeState))
	}
	return model.ERROR, errorx.NewInvalidArgumentError(fmt.Sprintf("未知的支付状态%s", tradeState))
}

const UTCWithXXXOffsetPattern = "2006-01-02T15:04:05.999Z07:00"

func parseDate(timeStr string) (time.Time, error) {
	parsedTime, err := time.Parse(UTCWithXXXOffsetPattern, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

type JsapiResult struct {
	Appid        string
	TimeStamp    string
	NonceStr     string
	PackageValue string
	SignType     string
	PaySign      string
}

func (j JsapiResult) getSignStr() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n", j.Appid, j.TimeStamp, j.NonceStr, j.PackageValue)
}

func (w *Client) getJsapiResult(prepayId string) (*JsapiResult, error) {
	timestampStr := strconv.FormatInt(time.Now().Unix(), 10)
	nonceStr, err := genRandomStr(32)
	if err != nil {
		return nil, err
	}
	var jsapiResult JsapiResult
	jsapiResult.Appid = w.Config.AppId
	jsapiResult.TimeStamp = timestampStr
	jsapiResult.NonceStr = nonceStr
	jsapiResult.PackageValue = "prepay_id=" + prepayId
	jsapiResult.SignType = "RSA"
	paySign, err := sign(jsapiResult.getSignStr(), w.Config.PrivateKeyContent)
	if err != nil {
		return nil, err
	}
	jsapiResult.PaySign = paySign
	return &jsapiResult, nil
}

type AppResult struct {
	Appid        string
	PartnerId    string
	PrepayId     string
	PackageValue string
	Noncestr     string
	Timestamp    string
	Sign         string
}

func (a *AppResult) getSignStr() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n", a.Appid, a.Timestamp, a.Noncestr, a.PrepayId)
}

func (w *Client) getAppResult(prepayId string) (*AppResult, error) {
	timestampStr := strconv.FormatInt(time.Now().Unix(), 10)
	nonceStr, err := genRandomStr(32)
	if err != nil {
		return nil, err
	}
	var appResult AppResult
	appResult.Appid = w.Config.AppId
	appResult.PrepayId = prepayId
	appResult.PartnerId = w.Config.MchId
	appResult.Timestamp = timestampStr
	appResult.Noncestr = nonceStr
	appResult.PackageValue = "Sign=WXPay"
	paySign, err := sign(appResult.getSignStr(), w.Config.PrivateKeyContent)
	if err != nil {
		return nil, err
	}
	appResult.Sign = paySign
	return &appResult, nil
}

func genRandomStr(length int) (string, error) {
	const base = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(base))))
		if err != nil {
			return "", err
		}
		result[i] = base[index.Int64()]
	}
	return string(result), nil
}

func sign(message string, privateKeyContent string) (string, error) {
	priKey, err := xpem.DecodePrivateKey([]byte(privateKeyContent))
	if err != nil {
		return "", err
	}
	hashed := sha256.Sum256([]byte(message))
	// 使用RSA私钥对待签名数据进行签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	// 对签名结果进行Base64编码
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)
	return signatureBase64, nil
}
