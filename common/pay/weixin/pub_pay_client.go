package weixin

import (
	"context"
	"errors"

	"github.com/agui-coder/simple-admin-pay-api/common/pay/model"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

// PubPayClient 微信支付（公众号）
type PubPayClient struct {
	Client
}

func NewWxPubPayClient(channelId uint64, config ClientConfig) *PubPayClient {
	return &PubPayClient{
		Client{Config: &config, ChannelId: channelId},
	}
}

func (w *PubPayClient) UnifiedOrder(ctx context.Context, req model.OrderUnifiedReq) (*model.OrderResp, error) {
	openid, err := getOpenid(req)
	if err != nil {
		return nil, err
	}
	bm := w.buildPayUnifiedOrderBm(req).SetBodyMap("payer", func(bm gopay.BodyMap) {
		bm.Set("openid", openid)
	})
	wxRsp, err := w.client.V3TransactionJsapi(ctx, bm)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	if wxRsp.Code == wechat.Success {
		result, err := w.getJsapiResult(wxRsp.Response.PrepayId)
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
	return nil, errors.New(wxRsp.Error)
}

func getOpenid(req model.OrderUnifiedReq) (string, error) {
	openid, exit := req.ChannelExtras["openid"]
	if !exit {
		return "", errors.New("支付请求的 openid 不能为空！")
	}
	return openid, nil
}
