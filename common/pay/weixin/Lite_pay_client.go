package weixin

// LitePayClient 结构体继承了 PubPayClient 接口 由于公众号和小程序的微信支付逻辑一致，所以直接进行继承
type LitePayClient struct {
	PubPayClient
}

func NewWxLitePayClient(channelId uint64, config ClientConfig) *AppPayClient {
	return &AppPayClient{
		Client{Config: &config, ChannelId: channelId},
	}
}
