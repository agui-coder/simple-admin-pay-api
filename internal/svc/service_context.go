package svc

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/payclient"

	"github.com/agui-coder/simple-admin-pay-api/common/pay"
	"github.com/agui-coder/simple-admin-pay-api/common/pay/model"

	"github.com/agui-coder/simple-admin-pay-api/internal/config"
	i18n2 "github.com/agui-coder/simple-admin-pay-api/internal/i18n"
	"github.com/agui-coder/simple-admin-pay-api/internal/middleware"
	"github.com/casbin/casbin/v2"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	Authority        rest.Middleware
	Casbin           *casbin.Enforcer
	PayRpc           payclient.Pay
	Redis            *redis.Redis
	Trans            *i18n.Translator
	PayClientFactory *pay.Factory
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(c.RedisConf)
	cbn := c.CasbinConf.MustNewCasbinWithRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(), c.RedisConf)
	trans := i18n.NewTranslator(i18n2.LocaleFS)
	return &ServiceContext{
		Config:           c,
		Authority:        middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		Redis:            rds,
		PayRpc:           payclient.NewPay(zrpc.NewClientIfEnable(c.PayRpc)),
		Trans:            trans,
		PayClientFactory: pay.NewFactory(),
	}
}

func (s *ServiceContext) GetPayClient(ctx context.Context, id uint64) (model.Client, error) {
	var client model.Client
	client, err := s.PayClientFactory.GetClient(id)
	if err == nil {
		return client, nil
	}
	channel, err := s.PayRpc.GetChannelById(ctx, &payclient.IDReq{Id: id})
	if err != nil {
		return nil, err
	}
	config, err := pay.GetClientConfig(*channel.Code, *channel.Config)
	if err != nil {
		return nil, err
	}
	client, err = s.PayClientFactory.CreateOrUpdatePayClient(*channel.Id, *channel.Code, config)
	if err != nil {
		return nil, err
	}
	return client, nil
}
