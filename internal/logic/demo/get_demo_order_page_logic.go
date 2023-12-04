package demo

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDemoOrderPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDemoOrderPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDemoOrderPageLogic {
	return &GetDemoOrderPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetDemoOrderPageLogic) GetDemoOrderPage(req *types.PageInfo) (resp *types.DemoOrderListResp, err error) {
	data, err := l.svcCtx.PayRpc.GetListDemoOrder(l.ctx, &payclient.DemoOrderPageReq{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		return nil, err
	}
	orderInfos := make([]*types.DemoOrderInfo, len(data.DemoOrderList))
	for i, order := range data.DemoOrderList {
		orderInfos[i] = &types.DemoOrderInfo{
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
		}
	}
	return &types.DemoOrderListResp{BaseListInfo: types.BaseListInfo{
		Total: data.Total,
	}, Data: orderInfos}, nil
}
