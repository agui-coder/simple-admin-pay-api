package base

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
	"github.com/agui-coder/simple-admin-pay-rpc/payclient"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	if l.svcCtx.Config.CoreRpc.Enabled {
		err = l.insertApiData()
		if err != nil {
			if strings.Contains(err.Error(), "common.createFailed") {
				return nil, errorx.NewInvalidArgumentError(i18n.AlreadyInit)
			}
			return nil, err
		}

		err = l.insertMenuData()
		if err != nil {
			return nil, err
		}

	}
	data, err := l.svcCtx.PayRpc.InitDatabase(l.ctx, &payclient.Empty{})
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.Casbin.LoadPolicy()
	if err != nil {
		logx.Errorw("failed to load Casbin Policy", logx.Field("detail", err))
		return nil, errorx.NewCodeInternalError(i18n.DatabaseError)
	}
	return &types.BaseMsgResp{
		Code: 0,
		Msg:  l.svcCtx.Trans.Trans(l.ctx, data.Msg),
	}, nil
}
