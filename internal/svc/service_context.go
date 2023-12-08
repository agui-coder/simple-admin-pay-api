package svc

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-common/payment"
	"github.com/agui-coder/simple-admin-pay-common/payment/model"
	"github.com/agui-coder/simple-admin-pay-rpc/payclient"
	"github.com/zeromicro/go-zero/core/collection"
	"log"
	"strconv"
	"time"

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
	PayClientFactory *payment.Factory
	PayClientCache   *collection.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(c.RedisConf)
	cbn := c.CasbinConf.MustNewCasbinWithRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(), c.RedisConf)
	trans := i18n.NewTranslator(i18n2.LocaleFS)
	cache, err := collection.NewCache(time.Second*10, collection.WithName("payClientCache"))
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:           c,
		Authority:        middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		Redis:            rds,
		PayRpc:           payclient.NewPay(zrpc.NewClientIfEnable(c.PayRpc)),
		Trans:            trans,
		PayClientFactory: payment.NewFactory(),
		PayClientCache:   cache,
	}
}

func (s *ServiceContext) GetPayClient(ctx context.Context, id uint64) (model.Client, error) {
	take, err := s.PayClientCache.Take("pay_client:"+strconv.FormatUint(id, 10), func() (any, error) {
		channel, err := s.PayRpc.GetChannelById(ctx, &payclient.IDReq{Id: id})
		if err == nil {
			payConfig, err := payment.ParseClientConfig(*channel.Code, *channel.Config)
			if err != nil {
				return nil, err
			}
			err = s.PayClientFactory.CreateOrUpdatePayClient(*channel.Id, *channel.Code, payConfig)
			if err != nil {
				return nil, err
			}
		}
		client, err := s.PayClientFactory.GetClient(*channel.Id)
		if err != nil {
			return nil, err
		}
		return client, nil
	})
	if err != nil {
		return nil, err
	}
	return take.(model.Client), nil
}
