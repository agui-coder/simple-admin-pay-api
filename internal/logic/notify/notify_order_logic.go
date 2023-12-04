package notify

import (
	"context"
	"encoding/json"
	"github.com/agui-coder/simple-admin-pay-api/common/pay/model"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyOrderLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	OrderResp *model.OrderResp
}

func NewNotifyOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyOrderLogic {
	return &NotifyOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *NotifyOrderLogic) NotifyOrder(req *types.NotifyRep) (resp string, err error) {
	channelNotifyData, err := json.Marshal(l.OrderResp)
	if err != nil {
		return "error", nil
	}
	_, err = l.svcCtx.PayRpc.NotifyOrder(l.ctx, &payclient.NotifyOrderReq{
		ChannelId:         req.ChannelId,
		Status:            uint32(l.OrderResp.Status),
		OutTradeNo:        l.OrderResp.OutTradeNo,
		ChannelNotifyData: string(channelNotifyData),
		SuccessTime:       l.OrderResp.SuccessTime.Unix(),
		ChannelOrderNo:    l.OrderResp.ChannelOrderNo,
		ChannelUserId:     l.OrderResp.ChannelUserId,
	})
	if err != nil {
		return "error", err
	}
	return "success", nil
}
