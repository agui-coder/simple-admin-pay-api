package notify

import (
	"context"
	"encoding/json"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotifyOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyOrderLogic {
	return &NotifyOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *NotifyOrderLogic) NotifyOrder(req *types.NotifyRep) (resp string, err error) {
	client, err := l.svcCtx.GetPayClient(l.ctx, req.ChannelId)
	if err != nil {
		return "error", nil
	}
	unifiedOrderResp, err := client.ParseOrderNotify(req.Header, req.Body)
	if err != nil {
		return "error", nil
	}
	channelNotifyData, err := json.Marshal(unifiedOrderResp)
	if err != nil {
		return "error", nil
	}
	_, err = l.svcCtx.PayRpc.NotifyOrder(l.ctx, &payclient.NotifyOrderReq{
		ChannelId:         req.ChannelId,
		Status:            uint32(unifiedOrderResp.Status),
		OutTradeNo:        unifiedOrderResp.OutTradeNo,
		ChannelNotifyData: string(channelNotifyData),
		SuccessTime:       unifiedOrderResp.SuccessTime.Unix(),
		ChannelOrderNo:    unifiedOrderResp.ChannelOrderNo,
		ChannelUserId:     unifiedOrderResp.ChannelUserId,
	})
	if err != nil {
		return "error", err
	}
	return "success", nil
}
