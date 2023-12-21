package notify

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyRefundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotifyRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyRefundLogic {
	return &NotifyRefundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *NotifyRefundLogic) NotifyRefund(req *types.NotifyRep) (resp string, err error) {
	_, err = l.svcCtx.PayRpc.NotifyRefund(l.ctx, &payclient.NotifyRefundReq{ChannelCode: req.ChannelCode, R: req.R})
	if err != nil {
		return "error", err
	}
	return "success", nil
}
