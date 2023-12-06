package ali

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-api/common/pay/model"
	"github.com/go-pay/gopay"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

// WapPayClient
type WapPayClient struct {
	Client
}

// 编译时接口实现的检查
var _ model.Client = (*WapPayClient)(nil)

func NewAliWapPayClient(channelId uint64, config model.ClientConfig) model.Client {
	aliConfig, ok := config.(ClientConfig)
	if !ok {
		logx.Error("config is not of type ali.ClientConfig")
		return nil
	}
	return &WapPayClient{
		Client{Config: &aliConfig, ChannelId: channelId},
	}
}

func (a *WapPayClient) UnifiedOrder(ctx context.Context, req model.OrderUnifiedReq) (*model.OrderResp, error) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", req.OutTradeNo)
	bm.Set("body", req.Body)
	bm.Set("subject", req.Subject)
	bm.Set("total_amount", formatAmount(req.Price))
	bm.Set("ProductCode", "QUICK_WAP_WAY") // 销售产品码. 目前扫码支付场景下仅支持 QUICK_WAP_WAY
	a.client.SetReturnUrl(req.ReturnUrl)
	a.client.SetNotifyUrl(req.NotifyUrl)
	bm.Set("quit_url", req.ReturnUrl)
	payUrl, err := a.client.TradeWapPay(ctx, bm)
	if err != nil {
		logx.Errorf("client.TradeWapPay(%+v),error:%+v", bm, err)
		return nil, err
	}
	return model.WaitingOf(pointy.GetPointer(model.Url), pointy.GetPointer(payUrl), req.OutTradeNo, payUrl), nil
}
