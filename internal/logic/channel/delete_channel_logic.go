package channel

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteChannelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteChannelLogic {
	return &DeleteChannelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeleteChannelLogic) DeleteChannel(req *types.IDReq) (resp *types.BaseMsgResp, err error) {
	baseResp, err := l.svcCtx.PayRpc.DeleteChannel(l.ctx, &payclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	// 清除 pay client
	l.svcCtx.PayClientFactory.ClearClient(req.Id)
	return &types.BaseMsgResp{Msg: baseResp.Msg}, nil
}
