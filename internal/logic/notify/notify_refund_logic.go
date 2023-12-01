package notify

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
