package order

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/agui-coder/simple-admin-pay-common/payno"
	"strconv"
	"time"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-common/payment/model"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitPayOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitPayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitPayOrderLogic {
	return &SubmitPayOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *SubmitPayOrderLogic) SubmitPayOrder(req *types.OrderSubmitReq) (resp *types.OrderSubmitResp, err error) {
	order, err := l.svcCtx.PayRpc.ValidateOrderCanSubmit(l.ctx, &payclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	channel, err := l.svcCtx.PayRpc.ValidateChannelCanSubmit(l.ctx, &payclient.ValidateChannelReq{AppId: *order.AppId, Code: req.ChannelCode})
	if err != nil {
		return nil, err
	}
	no, err := payno.Generate(l.svcCtx.Redis, payno.OrderNoPrefix)
	if err != nil {
		return nil, err
	}
	userIP, ok := l.ctx.Value("userIp").(string)
	if !ok {
		return nil, errorx.NewInvalidArgumentError("userIP error")
	}
	_, err = l.svcCtx.PayRpc.CreateOrderExtension(l.ctx, &payclient.OrderCreateExtensionReq{
		OrderID:       *order.Id,
		ChannelCode:   *channel.Code,
		ChannelExtras: req.ChannelExtras,
		No:            no,
		ChannelID:     *channel.Id,
		Status:        0,
		UserIP:        userIP,
	})
	if err != nil {
		return nil, err
	}
	client, err := l.svcCtx.GetPayClient(l.ctx, *channel.Id)
	if err != nil {
		logx.Errorf("[validatePayChannelCanSubmit][渠道编号(%d) 找不到对应的支付客户端]", channel.Id)
		return nil, err
	}
	unifiedOrderResp, err := client.UnifiedOrder(l.ctx, model.OrderUnifiedReq{
		DisplayMode:   req.DisplayMode,
		UserIp:        userIP,
		OutTradeNo:    no,
		Subject:       *order.Subject,
		Body:          *order.Body,
		NotifyUrl:     l.svcCtx.Config.PayProperties.OrderNotifyUrl + "/" + strconv.FormatUint(*channel.Id, 10),
		ReturnUrl:     req.ReturnUrl,
		Price:         *order.Price,
		ExpireTime:    time.UnixMilli(*order.ExpireTime),
		ChannelExtras: req.ChannelExtras,
	})
	if err != nil {
		return nil, err
	}
	if unifiedOrderResp != nil {
		channelNotifyData, err := json.Marshal(unifiedOrderResp)
		if err != nil {
			return nil, err
		}
		go func() {
			_, err = l.svcCtx.PayRpc.NotifyOrder(context.Background(), &payclient.NotifyOrderReq{
				Status:            uint32(unifiedOrderResp.Status),
				OutTradeNo:        unifiedOrderResp.OutTradeNo,
				ChannelNotifyData: string(channelNotifyData),
				SuccessTime:       unifiedOrderResp.SuccessTime.Unix(),
				ChannelOrderNo:    unifiedOrderResp.ChannelOrderNo,
				ChannelUserId:     unifiedOrderResp.ChannelUserId,
				ChannelId:         *channel.Id,
			})
			if err != nil {
				logx.Error(err)
			}
		}()
		if unifiedOrderResp.ChannelErrorCode != nil {
			return nil, errorx.NewInvalidArgumentError(fmt.Sprintf("发起支付报错，错误码：%d，错误提示：%d",
				unifiedOrderResp.ChannelErrorCode, unifiedOrderResp.ChannelErrorMsg))
		}
	}
	order, err = l.svcCtx.PayRpc.GetOrder(l.ctx, &payclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.OrderSubmitResp{Status: order.Status, DisplayMode: unifiedOrderResp.DisplayMode, DisplayContent: unifiedOrderResp.DisplayContent}, nil
}
