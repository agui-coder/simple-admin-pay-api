package weixin

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-api/common/pay/model"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

// AppPayClient 结构体继承了 Client 接口
type AppPayClient struct {
	Client
}

func NewWxAppPayClient(channelId uint64, config ClientConfig) *AppPayClient {
	return &AppPayClient{
		Client{Config: &config, ChannelId: channelId},
	}
}

func (w *AppPayClient) UnifiedOrder(ctx context.Context, req model.OrderUnifiedReq) (*model.OrderResp, error) {
	bm := w.buildPayUnifiedOrderBm(req)
	wxRsp, err := w.client.V3TransactionApp(ctx, bm)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	if wxRsp.Code == wechat.Success {
		result, err := w.getAppResult(wxRsp.Response.PrepayId)
		if err != nil {
			return nil, err
		}
		orderResp, err := model.WaitingOf(pointy.GetPointer(model.App), req.OutTradeNo, result)
		if err != nil {
			return nil, err
		}
		return orderResp, nil
	}
	logx.Errorf("wxRsp:%s", wxRsp.Error)
	return nil, errorx.NewInvalidArgumentError(wxRsp.Error)
}
