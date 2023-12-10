package demo

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-common/consts"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"strconv"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDemoRefundPaidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDemoRefundPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDemoRefundPaidLogic {
	return &UpdateDemoRefundPaidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateDemoRefundPaidLogic) UpdateDemoRefundPaid(req *types.PayRefundNotifyReq) (resp *types.BaseMsgResp, err error) {
	id, err := strconv.ParseUint(req.MerchantOrderId, 10, 64)
	if err != nil {
		return nil, err
	}
	data, err := l.svcCtx.PayRpc.UpdateDemoRefundPaid(l.ctx, &pay.UpdateDemoRefundPaidReq{Id: id, PayRefundId: req.PayRefundId})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Code: int(consts.SUCCESS), Msg: data.Msg}, nil
}
