package refund

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRefundPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRefundPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRefundPageLogic {
	return &GetRefundPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetRefundPageLogic) GetRefundPage(req *types.RefundPageReq) (resp *types.RefundPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
