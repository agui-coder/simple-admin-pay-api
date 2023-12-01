package channel

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateChannelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateChannelLogic {
	return &UpdateChannelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateChannelLogic) UpdateChannel(req *types.ChannelUpdateReq) (resp *types.BaseMsgResp, err error) {
	baseResp, err := l.svcCtx.PayRpc.UpdateChannel(l.ctx, &payclient.ChannelUpdateReq{
		Id:      req.Id,
		Config:  req.Config,
		Status:  req.Status,
		Remark:  req.Remark,
		FeeRate: req.FeeRate,
		AppId:   req.AppId,
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: baseResp.Msg}, nil
}
