package refund

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/payclient"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRefundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRefundLogic {
	return &GetRefundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetRefundLogic) GetRefund(req *types.IDReq) (resp *types.RefundInfoResp, err error) {
	data, err := l.svcCtx.PayRpc.GetRefundById(l.ctx, &payclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.RefundInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: &types.RefundInfo{
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
		},
	}, nil
}
