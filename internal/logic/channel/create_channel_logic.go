package channel

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateChannelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateChannelLogic {
	return &CreateChannelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *CreateChannelLogic) CreateChannel(req *types.ChannelCreateReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.PayRpc.CreateChannel(l.ctx, &payclient.ChannelCreateReq{
		Code:    req.Code,
		Config:  req.Config,
		Status:  req.Status,
		Remark:  req.Remark,
		FeeRate: req.FeeRate,
		AppId:   req.AppId,
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
