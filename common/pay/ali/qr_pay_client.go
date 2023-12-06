package ali

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-api/common/pay/model"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

// QrPayClient
type QrPayClient struct {
	Client
}

// 编译时接口实现的检查
var _ model.Client = (*QrPayClient)(nil)

func NewAliQrPayClient(channelId uint64, config model.ClientConfig) model.Client {
	aliConfig, ok := config.(ClientConfig)
	if !ok {
		logx.Error("config is not of type ali.ClientConfig")
		return nil
	}
	return &QrPayClient{
		Client{Config: &aliConfig, ChannelId: channelId},
	}
}

func (a *QrPayClient) UnifiedOrder(ctx context.Context, req model.OrderUnifiedReq) (*model.OrderResp, error) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", req.OutTradeNo)
	bm.Set("body", req.Body)
	bm.Set("subject", req.Subject)
	bm.Set("total_amount", formatAmount(req.Price))
	bm.Set("ProductCode", "FACE_TO_FACE_PAYMENT") // 销售产品码. 目前扫码支付场景下仅支持 FACE_TO_FACE_PAYMENT
	a.client.SetReturnUrl(req.ReturnUrl)
	a.client.SetNotifyUrl(req.NotifyUrl)
	resp, err := a.client.TradePrecreate(ctx, bm)
	if err != nil {
		if bizErr, ok := alipay.IsBizError(err); ok {
			logx.Errorf("%+v", bizErr)
			return model.CloseOf(resp.Response.SubCode, resp.Response.SubMsg, req.OutTradeNo, resp.Response), nil
		}
		return nil, err
	}
	return model.WaitingOf(pointy.GetPointer(model.QrCode), &resp.Response.QrCode, req.OutTradeNo, resp.Response), nil
}
