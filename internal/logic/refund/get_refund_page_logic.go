package refund

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRefundPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRefundPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRefundPageLogic {
	return &GetRefundPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetRefundPageLogic) GetRefundPage(req *types.RefundPageReq) (resp *types.RefundPageResp, err error) {
	pageResp, err := l.svcCtx.PayRpc.GetRefundPage(l.ctx, &payclient.RefundPageReq{
		Page:            req.Page,
		PageSize:        req.PageSize,
		ChannelCode:     req.ChannelCode,
		MerchantOrderId: req.MerchantOrderId,
		ChannelOrderNo:  req.ChannelOrderNo,
		Status:          req.Status,
		CreateTime:      req.CreateAt,
	})
	if err != nil {
		return nil, err
	}
	dataResp := make([]*types.RefundInfo, len(pageResp.Data))
	for i, data := range pageResp.Data {
		dataResp[i] = &types.RefundInfo{
			Id:                data.Id,
			CreatedAt:         data.CreatedAt,
			UpdatedAt:         data.UpdatedAt,
			Status:            data.Status,
			No:                data.No,
			ChannelCode:       data.ChannelCode,
			OrderId:           data.OrderId,
			OrderNo:           data.OrderNo,
			MerchantOrderId:   data.MerchantOrderId,
			MerchantRefundId:  data.MerchantRefundId,
			PayPrice:          data.PayPrice,
			RefundPrice:       data.RefundPrice,
			Reason:            data.Reason,
			UserIp:            data.UserIp,
			ChannelOrderNo:    data.ChannelOrderNo,
			ChannelRefundNo:   data.ChannelRefundNo,
			SuccessTime:       data.SuccessTime,
			ChannelErrorCode:  data.ChannelErrorCode,
			ChannelErrorMsg:   data.ChannelErrorMsg,
			ChannelNotifyData: data.ChannelNotifyData,
		}
	}

	return &types.RefundPageResp{
		BaseListInfo: types.BaseListInfo{
			Total: pageResp.Total,
		},
		Data: dataResp,
	}, nil
}
