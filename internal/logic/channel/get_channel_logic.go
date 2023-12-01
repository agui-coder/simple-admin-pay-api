package channel

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChannelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChannelLogic {
	return &GetChannelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetChannelLogic) GetChannel(req *types.ChannelReq) (resp *types.ChannelInfoResp, err error) {
	var channel *payclient.ChannelInfo
	if req.Id != nil {
		channel, err = l.svcCtx.PayRpc.GetChannelById(l.ctx, &payclient.IDReq{Id: *req.Id})
		if err != nil {
			return nil, err
		}
	} else if req.Code != nil && req.AppId != nil {
		channel, err = l.svcCtx.PayRpc.GetChannelListByAppIdAndCode(l.ctx, &payclient.ByAppIdAndCodeReq{Code: *req.Code, AppId: *req.AppId})
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	return &types.ChannelInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.ChannelInfo{
			ChannelBaseResp: types.ChannelBaseResp{
				Status:  channel.Status,
				Remark:  channel.Remark,
				FeeRate: channel.FeeRate,
				AppId:   channel.AppId,
			},
			Id:       channel.Id,
			Code:     channel.Code,
			Config:   channel.Config,
			CreateAt: channel.CreatedAt,
		},
	}, nil
}
