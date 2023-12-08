package config

import (
	"github.com/agui-coder/simple-admin-pay-common/payment/model"
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/plugins/casbin"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth          rest.AuthConf
	CROSConf      config.CROSConf
	PayRpc        zrpc.RpcClientConf
	DatabaseConf  config.DatabaseConf
	RedisConf     redis.RedisConf
	CasbinConf    casbin.CasbinConf
	PayProperties model.Properties
}
