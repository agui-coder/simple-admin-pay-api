package middleware

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-api/internal/consts"
	"github.com/duke-git/lancet/v2/netutil"
	"net/http"
)

type UserIpMiddleware struct {
}

func NewUserIpMiddleware() *UserIpMiddleware {
	return &UserIpMiddleware{}
}

func (m *UserIpMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), consts.UserIp, netutil.GetRequestPublicIp(r))
		next(w, r.WithContext(ctx))
	}
}
