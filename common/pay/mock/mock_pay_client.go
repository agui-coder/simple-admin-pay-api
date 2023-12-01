package mock

import (
	"context"
	"time"

	"github.com/agui-coder/simple-admin-pay-api/common/pay/model"
	"github.com/zeromicro/go-zero/core/errorx"
)

// ClientConfig 实现了 model.ClientConfig 接口 模拟支付方便调试
type ClientConfig struct {
	Name string
}

func (p ClientConfig) Validate() error {
	return nil
}

const respSuccessData = "MOCK_SUCCESS"

// Client 结构体实现了 model.Client 接口
type Client struct {
	ChannelId uint64
	Config    *ClientConfig
}

func (c *Client) ParseOrderNotify(map[string][]string, []byte) (*model.OrderResp, error) {
	return nil, errorx.NewInvalidArgumentError("mock no parse order notify")
}

func NewMockPayClient(channelId uint64, config ClientConfig) *Client {
	return &Client{
		Config: &config, ChannelId: channelId,
	}
}

func (c *Client) Init() error {
	return nil
}

func (c *Client) UnifiedOrder(ctx context.Context, req model.OrderUnifiedReq) (*model.OrderResp, error) {
	return model.SuccessOf("MOCK-P-"+req.OutTradeNo, "", time.Now(), req.OutTradeNo, respSuccessData), nil
}

func (c *Client) GetId() uint64 {
	return c.ChannelId
}

func (c *Client) GetOrder(ctx context.Context, outTradeNo string) (*model.OrderResp, error) {
	return model.SuccessOf("MOCK-P-"+outTradeNo, "", time.Now(),
		outTradeNo, respSuccessData), nil
}

func (c *Client) Refresh(config model.ClientConfig) error {
	return nil
}

func (c *Client) UnifiedRefund(ctx context.Context, req model.RefundUnifiedReq) (*model.RefundResp, error) {
	return nil, nil
}
