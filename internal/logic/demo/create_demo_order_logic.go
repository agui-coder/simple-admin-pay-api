package demo

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-api/internal/consts"
	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDemoOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDemoOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDemoOrderLogic {
	return &CreateDemoOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *CreateDemoOrderLogic) CreateDemoOrder(req *types.CreateDemoOrderReq) (resp *types.BaseMsgResp, err error) {
	userId := l.ctx.Value(consts.UserId).(string)
	userIp := l.ctx.Value(consts.UserIp).(string)
	data, err := l.svcCtx.PayRpc.CreateDemoOrder(l.ctx, &payclient.PayDemoOrderCreateReq{UserId: userId, SpuId: req.SpuId, UserIp: userIp})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
