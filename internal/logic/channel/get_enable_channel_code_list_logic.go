package channel

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEnableChannelCodeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEnableChannelCodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEnableChannelCodeListLogic {
	return &GetEnableChannelCodeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetEnableChannelCodeListLogic) GetEnableChannelCodeList(req *types.IDReq) (resp *types.ChannelListResp, err error) {
	channelList, err := l.svcCtx.PayRpc.GetEnableChannelList(l.ctx, &payclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	channelCodes := make([]*string, len(channelList.Data))
	for i, datum := range channelList.Data {
		channelCodes[i] = datum.Code
	}
	return &types.ChannelListResp{Channels: channelCodes}, nil
}
