package order

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-api/internal/middleware"
	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitPayOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitPayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitPayOrderLogic {
	return &SubmitPayOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *SubmitPayOrderLogic) SubmitPayOrder(req *types.OrderSubmitReq) (resp *types.OrderSubmitResp, err error) {
	UserIp := l.ctx.Value(middleware.UserIp).(string)
	data, err := l.svcCtx.PayRpc.SubmitPayOrder(l.ctx, &payclient.OrderSubmitReq{
		Id:            req.Id,
		ChannelCode:   req.ChannelCode,
		ChannelExtras: req.ChannelExtras,
		ReturnUrl:     req.ReturnUrl,
		UserIP:        UserIp,
	})
	if err != nil {
		return nil, err
	}
	return &types.OrderSubmitResp{Status: &data.Status, DisplayMode: data.DisplayMode, DisplayContent: data.DisplayContent}, nil
}
