package pay

import (
	"encoding/json"
	"fmt"

	"github.com/agui-coder/simple-admin-pay-api/common/pay/ali"
	"github.com/agui-coder/simple-admin-pay-api/common/pay/mock"
	"github.com/agui-coder/simple-admin-pay-api/common/pay/model"
	"github.com/agui-coder/simple-admin-pay-api/common/pay/weixin"
	"github.com/zeromicro/go-zero/core/errorx"
)

type Factory struct {
	clients map[uint64]model.Client
}

// NewFactory 工厂方法创建支付客户端
func NewFactory() *Factory {
	factory := &Factory{
		clients: make(map[uint64]model.Client),
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

func (f *Factory) CreateOrUpdatePayClient(channelId uint64, channelCode string, config model.ClientConfig) (model.Client, error) {
	client, exit := f.clients[channelId]
	if !exit {
		newClient, err := getNewClient(channelId, channelCode, config)
		if err != nil {
			return nil, err
		}
		err = newClient.Init()
		if err != nil {
			return nil, err
		}
		f.clients[channelId] = newClient
		client = newClient
	} else {
		err := client.Refresh(config)
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}

func getNewClient(channelId uint64, channelCode string, config model.ClientConfig) (model.Client, error) {
	var client model.Client

	switch channelCode {
	case model.WxApp, model.WxPub, model.WxNative, model.WxLite:
		if c, ok := config.(weixin.ClientConfig); ok {
			switch channelCode {
			case model.WxApp:
				client = weixin.NewWxAppPayClient(channelId, c)
			case model.WxPub:
				client = weixin.NewWxPubPayClient(channelId, c)
			case model.WxNative:
				client = weixin.NewWxNativePayClient(channelId, c)
			case model.WxLite:
				client = weixin.NewWxLitePayClient(channelId, c)
			}
		}
	case model.AlipayPc, model.AlipayWap, model.AlipayApp, model.AlipayBar, model.AlipayQr:
		if c, ok := config.(ali.ClientConfig); ok {
			switch channelCode {
			case model.AlipayPc:
				client = ali.NewAliPcPayClient(channelId, c)
			case model.AlipayWap:
				client = ali.NewAliWapPayClient(channelId, c)
			case model.AlipayApp:
				return nil, errorx.NewInvalidArgumentError("not support alipay app")
			case model.AlipayBar:
				client = ali.NewAliBarPayClient(channelId, c)
			case model.AlipayQr:
				client = ali.NewAliQrPayClient(channelId, c)
			}

		}
	case model.Mock:
		if c, ok := config.(mock.ClientConfig); ok {
			client = mock.NewMockPayClient(channelId, c)
		}
	default:
		return nil, errorx.NewInvalidArgumentError(fmt.Sprintf("unsupported channel code: %s", channelCode))
	}

	if client == nil {
		return nil, errorx.NewInvalidArgumentError(fmt.Sprintf("unsupported channel code: %s", channelCode))
	}

	return client, nil
}

func GetClientConfig(channelCode string, config string) (clientConfig model.ClientConfig, err error) {
	switch channelCode {
	case model.WxApp, model.WxPub, model.WxNative, model.WxLite:
		var wxClientConfig weixin.ClientConfig
		err = json.Unmarshal([]byte(config), &wxClientConfig)
		if err != nil {
			return nil, err
		}
		clientConfig = wxClientConfig
	case model.AlipayBar, model.AlipayApp, model.AlipayPc, model.AlipayWap, model.AlipayQr:
		var aliClientConfig ali.ClientConfig
		err = json.Unmarshal([]byte(config), &aliClientConfig)
		if err != nil {
			return nil, err
		}
		clientConfig = aliClientConfig
	case model.Mock:
		var mockClientConfig mock.ClientConfig
		clientConfig = mockClientConfig
		return clientConfig, nil
	default:
		return nil, errorx.NewInvalidArgumentError(fmt.Sprintf("unsupported channel code: %s", channelCode))
	}
	if clientConfig == nil {
		return nil, errorx.NewInvalidArgumentError(fmt.Sprintf("unsupported channel code: %s", channelCode))
	}
	return clientConfig, nil
}
