package demo

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefundDemoOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefundDemoOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefundDemoOrderLogic {
	return &RefundDemoOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *RefundDemoOrderLogic) RefundDemoOrder(req *types.IDReq) (resp *types.BaseMsgResp, err error) {
	userIp := l.ctx.Value("userIp").(string)
	data, err := l.svcCtx.PayRpc.RefundDemoOrder(l.ctx, &payclient.RefundDemoOrderReq{
		Id:     req.Id,
		UserIp: userIp,
	})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
