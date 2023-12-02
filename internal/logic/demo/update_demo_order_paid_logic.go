package demo

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-api/common/consts"
	"strconv"

	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDemoOrderPaidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDemoOrderPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDemoOrderPaidLogic {
	return &UpdateDemoOrderPaidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateDemoOrderPaidLogic) UpdateDemoOrderPaid(req *types.PayOrderNotifyReq) (resp *types.BaseMsgResp, err error) {
	id, err := strconv.ParseUint(req.MerchantOrderId, 10, 64)
	if err != nil {
		return nil, err
	}
	data, err := l.svcCtx.PayRpc.UpdateDemoOrderPaid(l.ctx, &pay.UpdateDemoOrderPaidReq{Id: id, PayOrderId: req.PayOrderId})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Code: int(consts.SUCCESS), Msg: data.Msg}, nil
}
