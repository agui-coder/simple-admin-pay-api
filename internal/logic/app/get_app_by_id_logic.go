package app

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAppByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppByIdLogic {
	return &GetAppByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetAppByIdLogic) GetAppById(req *types.IDReq) (resp *types.AppInfoResp, err error) {
	app, err := l.svcCtx.PayRpc.GetApp(l.ctx, &payclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.AppInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.AppInfo{
			Name:            app.Name,
			Status:          app.Status,
			Remark:          app.Remark,
			OrderNotifyUrl:  app.OrderNotifyUrl,
			RefundNotifyUrl: app.RefundNotifyUrl,
			BaseIDInfo: types.BaseIDInfo{
				Id:        app.Id,
				CreatedAt: app.CreatedAt,
				UpdatedAt: app.UpdatedAt,
			},
		},
	}, nil
}
