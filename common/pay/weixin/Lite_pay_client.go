package weixin

import (
	"github.com/agui-coder/simple-admin-pay-api/common/pay/model"
	"github.com/zeromicro/go-zero/core/logx"
)

// LitePayClient 结构体继承了 PubPayClient 接口 由于公众号和小程序的微信支付逻辑一致，所以直接进行继承
type LitePayClient struct {
	PubPayClient
}

// 编译时接口实现的检查
var _ model.Client = (*LitePayClient)(nil)

func NewWxLitePayClient(channelId uint64, config model.ClientConfig) model.Client {
	wxConfig, ok := config.(ClientConfig)
	if !ok {
		logx.Error("config is not of type weixin.ClientConfig")
		return nil
	}
	return &LitePayClient{
		PubPayClient{Client{Config: &wxConfig, ChannelId: channelId}},
	}
}
