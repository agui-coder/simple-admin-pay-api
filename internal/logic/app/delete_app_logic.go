package app

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAppLogic {
	return &DeleteAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeleteAppLogic) DeleteApp(req *types.IDReq) (resp *types.BaseMsgResp, err error) {
	baseResp, err := l.svcCtx.PayRpc.DeleteApp(l.ctx, &payclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: baseResp.Msg}, nil
}
