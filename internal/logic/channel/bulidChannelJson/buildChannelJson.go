package main

import (
	"encoding/json"
	"fmt"

	"github.com/agui-coder/simple-admin-pay-api/internal/types"
	"github.com/agui-coder/simple-admin-pay-common/payment/ali"
	"github.com/agui-coder/simple-admin-pay-common/payment/model"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/alipay/cert"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

// 打印 channel.http 需要的数据
func main() {
	config := buildAliClientConfig()
	marshal, err := json.Marshal(config)
	if err != nil {
		logx.Error(err.Error())
	}
	aliConfig := types.ChannelCreateReq{
		Code:   model.AlipayPc,
		Config: string(marshal),
		ChannelBase: types.ChannelBase{
			Remark:  pointy.GetPointer("支付宝电脑网站支付"),
			Status:  1,
			FeeRate: 0.006,
			AppId:   2,
		},
	}
	bytes, err := json.Marshal(&aliConfig)
	if err != nil {
		logx.Error(err.Error())
	}
	fmt.Println(string(bytes))
}

func buildAliClientConfig() ali.ClientConfig {
	return ali.ClientConfig{
		AppId:                   "待填",
		SignType:                alipay.RSA2,
		PrivateKey:              cert.PrivateKey,
		AppPublicContent:        string(cert.AppPublicContent),
		AlipayPublicContentRSA2: string(cert.AlipayPublicContentRSA2),
		AlipayRootContent:       string(cert.AlipayRootContent),
	}
}
