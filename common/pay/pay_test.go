package pay

import (
	"context"
	"fmt"
	"github.com/agui-coder/simple-admin-pay-api/common/pay/ali"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/alipay/cert"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/agui-coder/simple-admin-pay-api/common/pay/model"
	"github.com/agui-coder/simple-admin-pay-api/common/pay/weixin"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	factory      = NewFactory()
	channelIdMap = make(map[string]uint64)
)

func TestCreateWxPayClient(t *testing.T) {
	channelIdMap[model.WxPub] = 1
	_, err := factory.CreateOrUpdatePayClient(1, model.WxPub, buildWxClientConfig())
	if err != nil {
		logx.Error(err.Error())
	}
	client, err := factory.GetClient(channelIdMap[model.WxPub])
	if err != nil {
		logx.Error(err.Error())
	}
	order, err := client.UnifiedOrder(context.Background(), buildPayOrderUnifiedReq())
	if err != nil {
		logx.Error(err.Error())
	}
	fmt.Print(&order)
}

func TestCreateAliPayClient(t *testing.T) {
	channelIdMap[model.AlipayPc] = 2
	_, err := factory.CreateOrUpdatePayClient(2, model.AlipayPc, buildAliClientConfig())
	if err != nil {
		logx.Error(err.Error())
	}
	client, err := factory.GetClient(channelIdMap[model.AlipayPc])
	if err != nil {
		logx.Error(err.Error())
	}
	resp, err := client.UnifiedOrder(context.Background(), buildPayOrderUnifiedReq())
	if err != nil {
		logx.Error(err.Error())
	}
	fmt.Print(resp)
}

func buildPayOrderUnifiedReq() model.OrderUnifiedReq {
	m := make(map[string]string)
	// TODO openid
	m["openid"] = ""
	unifiedReq := model.OrderUnifiedReq{
		Price:         123,
		Subject:       "IPhone 13",
		Body:          "biubiubiu",
		OutTradeNo:    strconv.FormatInt(time.Now().UnixNano(), 10),
		UserIp:        "127.0.0.1",
		NotifyUrl:     "http://127.0.0.1:9107",
		ChannelExtras: m,
	}
	return unifiedReq
}

func buildWxClientConfig() weixin.ClientConfig {
	privateKeyFile, err := os.ReadFile("apiclient_key.pem")
	if err != nil {
		logx.Error(err.Error())
	}
	//TODO
	return weixin.ClientConfig{
		AppId:             "",
		MchId:             "",
		PrivateKeyContent: string(privateKeyFile),
		SerialNumber:      "",
		ApiV3Key:          "",
	}
}

func buildAliClientConfig() ali.ClientConfig {
	return ali.ClientConfig{
		AppId:                   "代填",
		SignType:                alipay.RSA2,
		PrivateKey:              cert.PrivateKey,
		AppPublicContent:        string(cert.AppPublicContent),
		AlipayPublicContentRSA2: string(cert.AlipayPublicContentRSA2),
		AlipayRootContent:       string(cert.AlipayRootContent),
	}
}
