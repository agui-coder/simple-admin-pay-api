package middleware

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthorityMiddleware struct {
	Cbn   *casbin.Enforcer
	Rds   *redis.Redis
	Trans *i18n.Translator
}

func NewAuthorityMiddleware(cbn *casbin.Enforcer, rds *redis.Redis, trans *i18n.Translator) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		Cbn:   cbn,
		Rds:   rds,
		Trans: trans,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "userId", "test1")
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		if host == "::1" {
			host = "127.0.0.1"
		}
		ipAddr, err := net.ResolveIPAddr("ip", host)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		ipv4 := ipAddr.IP.To4()
		if ipv4 == nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("unable to get IPv4 address"))
			return
		}
		ctx = context.WithValue(ctx, "userIp", ipv4.String())
		next(w, r.WithContext(ctx))
		return
	}
}
