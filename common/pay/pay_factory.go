package pay

import (
	"fmt"

	"github.com/agui-coder/simple-admin-pay-api/common/pay/ali"
	"github.com/agui-coder/simple-admin-pay-api/common/pay/mock"
	"github.com/agui-coder/simple-admin-pay-api/common/pay/model"
	"github.com/agui-coder/simple-admin-pay-api/common/pay/weixin"
	"github.com/zeromicro/go-zero/core/errorx"
)

type Factory struct {
	clients            map[uint64]model.Client
	clientConstructors map[string]func(uint64, model.ClientConfig) model.Client
}

// NewFactory 工厂方法创建支付客户端
func NewFactory() *Factory {
	factory := &Factory{
		clients: make(map[uint64]model.Client),
		clientConstructors: map[string]func(uint64, model.ClientConfig) model.Client{
			model.WxApp:     weixin.NewWxAppPayClient,
			model.WxPub:     weixin.NewWxPubPayClient,
			model.WxNative:  weixin.NewWxNativePayClient,
			model.WxLite:    weixin.NewWxLitePayClient,
			model.AlipayPc:  ali.NewAliPcPayClient,
			model.AlipayWap: ali.NewAliWapPayClient,
			model.AlipayBar: ali.NewAliBarPayClient,
			model.AlipayQr:  ali.NewAliQrPayClient,
			model.Mock:      mock.NewMockPayClient,
		},
	}
	return factory
}

// GetClient 获取支付客户端
func (f *Factory) GetClient(clientID uint64) (model.Client, error) {
	client, ok := f.clients[clientID]
	if !ok {
		return nil, errorx.NewInvalidArgumentError("invalid client ID")
	}
	return client, nil
}

func (f *Factory) ClearClient(clientID uint64) {
	delete(f.clients, clientID)
}

func (f *Factory) CreateOrUpdatePayClient(channelId uint64, channelCode string, config model.ClientConfig) error {
	client, exit := f.clients[channelId]
	if !exit {
		newClient, err := f.getNewClient(channelId, channelCode, config)
		if err != nil {
			return err
		}
		err = newClient.Init()
		if err != nil {
			return err
		}
		f.clients[channelId] = newClient
	} else {
		err := client.Refresh(config)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Factory) getNewClient(channelId uint64, channelCode string, config model.ClientConfig) (model.Client, error) {
	constructor, ok := f.clientConstructors[channelCode]
	if !ok {
		return nil, errorx.NewInvalidArgumentError(fmt.Sprintf("unsupported channel code: %s", channelCode))
	}
	client := constructor(channelId, config)
	if client == nil {
		return nil, errorx.NewInvalidArgumentError(fmt.Sprintf("unsupported channel code: %s", channelCode))
	}
	return client, nil
}

func ParseClientConfig(channelCode string, config string) (clientConfig model.ClientConfig, err error) {
	switch channelCode {
	case model.WxApp, model.WxPub, model.WxNative, model.WxLite:
		clientConfig, err = weixin.ParseWxClientConfig(config)
	case model.AlipayBar, model.AlipayApp, model.AlipayPc, model.AlipayWap, model.AlipayQr:
		clientConfig, err = ali.ParseAliClientConfig(config)
	case model.Mock:
		clientConfig, err = mock.ParseMockClientConfig(config)
	default:
		return nil, errorx.NewInvalidArgumentError(fmt.Sprintf("unsupported channel code: %s", channelCode))
	}
	return clientConfig, err
}
