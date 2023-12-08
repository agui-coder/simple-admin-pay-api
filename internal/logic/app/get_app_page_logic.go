package app

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-common/convert"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAppPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppPageLogic {
	return &GetAppPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetAppPageLogic) GetAppPage(req *types.AppListReq) (resp *types.AppListResp, err error) {
	appListResp, err := l.svcCtx.PayRpc.GetAppPage(l.ctx, &payclient.AppPageReq{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		return nil, err
	}
	data := make([]types.AppPageItemResp, len(appListResp.AppList))
	appIds := convert.List(appListResp.AppList, func(t *payclient.AppInfo) uint64 {
		return *t.Id
	})
	channelPage, err := l.svcCtx.PayRpc.GetChannelListByAppIds(l.ctx, &payclient.IDsReq{Ids: appIds})
	if err != nil {
		return nil, err
	}
	appIdChannelMap := make(map[uint64][]*string, len(channelPage.Data))
	for _, datum := range channelPage.Data {
		appIdChannelMap[*datum.AppId] = append(appIdChannelMap[*datum.AppId], datum.Code)
	}
	for i, info := range appListResp.AppList {
		data[i] = types.AppPageItemResp{
			AppBaseResp: types.AppBaseResp{
				Name:            info.Name,
				Status:          info.Status,
				Remark:          info.Remark,
				OrderNotifyUrl:  info.OrderNotifyUrl,
				RefundNotifyUrl: info.RefundNotifyUrl,
			},
			Id:           info.Id,
			CreateAt:     info.CreatedAt,
			ChannelCodes: appIdChannelMap[*info.Id],
		}
	}
	return &types.AppListResp{
		BaseListInfo: types.BaseListInfo{
			Total: appListResp.Total,
		},
		Data: data,
	}, nil
}
