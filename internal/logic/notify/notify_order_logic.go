package notify

import (
	"context"
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
	_, err = l.svcCtx.PayRpc.NotifyOrder(l.ctx, &payclient.NotifyOrderReq{
		Code: req.ChannelCode,
		R:    req.R,
	})
	if err != nil {
		return "error", err
	}
	return "success", nil
}
