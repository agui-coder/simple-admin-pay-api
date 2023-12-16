package order

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderPageLogic {
	return &GetOrderPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetOrderPageLogic) GetOrderPage(req *types.OrderPageReq) (resp *types.OrderPageResp, err error) {
	orderPage, err := l.svcCtx.PayRpc.GetOrderPage(l.ctx, &payclient.OrderPageReq{
		Page:            req.Page,
		PageSize:        req.PageSize,
		ChannelCode:     req.ChannelCode,
		MerchantOrderId: req.MerchantOrderId,
		ChannelOrderNo:  req.ChannelOrderNo,
		No:              req.No,
		Status:          req.Status,
		CreateTime:      req.CreateAt,
	})
	if err != nil {
		return nil, err
	}
	orderList := make([]*types.OrderInfo, len(orderPage.Data))

	for i, order := range orderPage.Data {
		orderList[i] = &types.OrderInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        order.Id,
				CreatedAt: order.CreatedAt,
				UpdatedAt: order.UpdatedAt,
			},
			Status:          order.Status,
			ChannelCode:     order.ChannelCode,
			MerchantOrderId: order.MerchantOrderId,
			Subject:         order.Subject,
			Body:            order.Body,
			Price:           order.Price,
			ChannelFeeRate:  order.ChannelFeeRate,
			ChannelFeePrice: order.ChannelFeePrice,
			UserIp:          order.UserIp,
			ExpireTime:      order.ExpireTime,
			SuccessTime:     order.SuccessTime,
			NotifyTime:      order.NotifyTime,
			ExtensionId:     order.ExtensionId,
			No:              order.No,
			RefundPrice:     order.RefundPrice,
			ChannelUserId:   order.ChannelUserId,
			ChannelOrderNo:  order.ChannelOrderNo,
		}
	}

	return &types.OrderPageResp{
		BaseListInfo: types.BaseListInfo{
			Total: orderPage.Total,
		},
		Data: orderList}, nil
}
