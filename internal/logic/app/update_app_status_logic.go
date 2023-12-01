package app

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAppStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAppStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAppStatusLogic {
	return &UpdateAppStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateAppStatusLogic) UpdateAppStatus(req *types.AppUpdateStatusReq) (resp *types.BaseMsgResp, err error) {
	baseResp, err := l.svcCtx.PayRpc.UpdateAppStatus(l.ctx, &payclient.AppUpdateStatusReq{Status: req.Status, Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: baseResp.Msg}, nil
}
