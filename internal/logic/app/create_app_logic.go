package app

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAppLogic {
	return &CreateAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *CreateAppLogic) CreateApp(req *types.AppCreateReq) (resp *types.BaseMsgResp, err error) {
	createAppResp, err := l.svcCtx.PayRpc.CreateApp(l.ctx, &payclient.AppCreateReq{
		Name:            req.Name,
		Status:          req.Status,
		Remark:          req.Remark,
		OrderNotifyUrl:  req.OrderNotifyUrl,
		RefundNotifyUrl: req.RefundNotifyUrl,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: createAppResp.Msg}, nil
}
