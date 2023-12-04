package ali

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-api/common/pay/model"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// BarPayClient
type BarPayClient struct {
	Client
}

func NewAliBarPayClient(channelId uint64, config ClientConfig) *BarPayClient {
	return &BarPayClient{
		Client{Config: &config, ChannelId: channelId},
	}
}

func (a *BarPayClient) UnifiedOrder(ctx context.Context, req model.OrderUnifiedReq) (*model.OrderResp, error) {
	authCode, ok := req.ChannelExtras["auth_code"]
	if !ok {
		return nil, errorx.NewApiInternalError("auth_code not found")
	}
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", req.OutTradeNo)
	bm.Set("subject", req.Subject)
	bm.Set("body", req.Body)
	bm.Set("total_amount", formatAmount(req.Price))
	bm.Set("scene", "bar_code")   // 支付场景 条码支付，取值：bar_code 声波支付，取值：wave_code
	bm.Set("auth_code", authCode) // 支付授权码
	a.client.SetReturnUrl(req.ReturnUrl)
	a.client.SetNotifyUrl(req.NotifyUrl)
	resp, err := a.client.TradePay(ctx, bm)
	if err != nil {
		if bizErr, ok := alipay.IsBizError(err); ok {
			logx.Errorf("%+v", bizErr)
			return nil, bizErr
		}
		return nil, err
	}
	if "10000" == resp.Response.Code { // 免密支付
		// TODO go pay 包缺少 gmt_payment 字段 暂时用 当前时间代替 支付成功时间
		return model.SuccessOf(resp.Response.TradeNo, resp.Response.BuyerUserId, time.Now(), resp.Response.OutTradeNo, resp.Response), nil
	}
	return model.WaitingOf(pointy.GetPointer(model.BarCode), pointy.GetPointer(""), req.OutTradeNo, resp.Response), nil
}
