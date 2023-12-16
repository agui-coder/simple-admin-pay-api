package order

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailLogic {
	return &GetOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetOrderDetailLogic) GetOrderDetail(req *types.IDReq) (resp *types.OrderDetailResp, err error) {
	order, err := l.svcCtx.PayRpc.GetOrder(l.ctx, &payclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	extension, err := l.svcCtx.PayRpc.GetOrderExtension(l.ctx, &payclient.IDReq{Id: *order.ExtensionId})
	if err != nil {
		return nil, err
	}
	return &types.OrderDetailResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: &types.OrderDetail{
			OrderInfo: types.OrderInfo{
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
			},
			Extension: &types.OrderExtension{
				No:                extension.No,
				ChannelNotifyData: extension.ChannelNotifyData,
			},
		},
	}, nil
}
