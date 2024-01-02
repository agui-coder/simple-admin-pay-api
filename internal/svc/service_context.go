package svc

import (
	"github.com/agui-coder/simple-admin-pay-api/internal/config"
	i18n2 "github.com/agui-coder/simple-admin-pay-api/internal/i18n"
	"github.com/agui-coder/simple-admin-pay-api/internal/middleware"
	"github.com/agui-coder/simple-admin-pay-rpc/payclient"
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/coreclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware
	UserIp    rest.Middleware
	Casbin    *casbin.Enforcer
	PayRpc    payclient.Pay
	CoreRpc   coreclient.Core
	Redis     redis.UniversalClient
	Trans     *i18n.Translator
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := c.RedisConf.MustNewUniversalRedis()
	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(), c.RedisConf)
	trans := i18n.NewTranslator(i18n2.LocaleFS)
	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		Redis:     rds,
		PayRpc:    payclient.NewPay(zrpc.NewClientIfEnable(c.PayRpc)),
		CoreRpc:   coreclient.NewCore(zrpc.NewClientIfEnable(c.CoreRpc)),
		Trans:     trans,
		UserIp:    middleware.NewUserIpMiddleware().Handle,
	}
}
