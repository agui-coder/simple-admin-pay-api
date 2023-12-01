package demo

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDemoOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDemoOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDemoOrderLogic {
	return &GetDemoOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetDemoOrderLogic) GetDemoOrder(req *types.IDAtPathReq) (resp *types.DemoOrderInfo, err error) {
	order, err := l.svcCtx.PayRpc.GetDemoOrder(l.ctx, &payclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.DemoOrderInfo{
		Id:             order.Id,
		CreateAt:       order.CreatedAt,
		UpdateAt:       order.UpdatedAt,
		SpuId:          order.SpuId,
		SpuName:        order.SpuName,
		Price:          order.Price,
		PayStatus:      order.PayStatus,
		PayOrderId:     order.PayOrderId,
		PayTime:        order.PayTime,
		PayChannelCode: order.PayChannelCode,
		PayRefundId:    order.PayRefundId,
		RefundPrice:    order.RefundPrice,
		RefundTime:     order.RefundTime,
	}, nil
}
