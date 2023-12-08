package refund

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRefundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRefundLogic {
	return &GetRefundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetRefundLogic) GetRefund(req *types.IDReq) (resp *types.RefundInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
