package ali

import (
	"context"
	"errors"

	"github.com/agui-coder/simple-admin-pay-api/common/pay/model"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/zeromicro/go-zero/core/logx"
)

// ClientConfig 没支付宝商户号就没写
// ClientConfig 实现了 pay.ClientConfig 接口
type ClientConfig struct {
	ServerUrl               string
	AppId                   string
	SignType                string
	Mode                    int32
	PrivateKey              string
	AlipayPublicKey         string
	AppCertContent          string
	AlipayPublicCertContent string
	RootCertContent         string
}

func (p ClientConfig) Validate() error {
	return nil
}

// Client 结构体实现了 PayClient 接口
type Client struct {
	ChannelId uint64
	Config    *ClientConfig
	client    *alipay.Client
}

func (a *Client) Init() error {
	// 初始化支付宝客户端
	// appid：应用ID
	// privateKey：应用私钥，支持PKCS1和PKCS8
	// isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
	client, err := alipay.NewClient(a.Config.AppId, a.Config.PrivateKey, true)
	if err != nil {
		logx.Error(err)
		return errors.New("支付客户端初始化异常")
	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//client.SetBodySize() // 没有特殊需求，可忽略此配置

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetCharset(alipay.UTF8). // 设置字符编码，不设置默认 utf-8
					SetSignType(a.Config.SignType) // 设置签名类型，不设置默认 RSA2
	//SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
	//SetReturnUrl("https://www.fmm.ink"). // 设置返回URL
	//SetNotifyUrl("https://www.fmm.ink"). // 设置异步通知URL
	//SetAppAuthToken() // TODO 设置第三方应用授权

	// 设置biz_content加密KEY，设置此参数默认开启加密
	// client.SetAESKey("1234567890123456")
	// 自动同步验签（只支持证书模式）
	// 传入 alipayPublicCert.crt 内容
	client.AutoVerifySign([]byte(a.Config.AlipayPublicKey))
	// 证书内容
	err = client.SetCertSnByContent([]byte(a.Config.AppCertContent), []byte(a.Config.RootCertContent), []byte(a.Config.AlipayPublicCertContent))
	if err != nil {
		logx.Error(err)
		return err
	}
	a.client = client
	return nil
}

func (a PcPayClient) ParseOrderNotify(map[string][]string, []byte) (*model.OrderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (a Client) UnifiedOrder(ctx context.Context, req model.OrderUnifiedReq) (*model.OrderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (a Client) GetId() uint64 {
	return a.ChannelId
}

func (a Client) GetOrder(c context.Context, s string) (*model.OrderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (a Client) Refresh(config model.ClientConfig) error {
	//TODO implement me
	panic("implement me")
}

func (a PcPayClient) UnifiedRefund(ctx context.Context, req model.RefundUnifiedReq) (*model.RefundResp, error) {
	//TODO implement me
	panic("implement me")
}
