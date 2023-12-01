package pay

import (
	"context"
	"fmt"
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
	wxConfig     = buildWxClientConfig()
	channelIdMap = make(map[string]uint64)
)

func TestMain(m *testing.M) {
	channelIdMap[model.WxPub] = 1
	_, err := factory.CreateOrUpdatePayClient(1, model.WxPub, wxConfig)
	if err != nil {
		logx.Error(err.Error())
	}
	os.Exit(m.Run())
}

func TestCreateWxPayClient(t *testing.T) {
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
	privateKeyFile, err := os.ReadFile("D:\\Users\\DESK-0010\\Downloads\\WXCertUtil\\cert\\cert\\apiclient_key.pem")
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
