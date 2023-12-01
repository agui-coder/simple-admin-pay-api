package app

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAppLogic {
	return &UpdateAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateAppLogic) UpdateApp(req *types.AppUpdateReq) (resp *types.BaseMsgResp, err error) {
	baseResp, err := l.svcCtx.PayRpc.UpdateApp(l.ctx, &payclient.AppUpdateReq{
		Id:              req.Id,
		Name:            req.Name,
		Status:          req.Status,
		Remark:          req.Remark,
		OrderNotifyUrl:  req.OrderNotifyUrl,
		RefundNotifyUrl: req.RefundNotifyUrl,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: baseResp.Msg}, nil
}
