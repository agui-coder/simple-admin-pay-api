package channel

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/channel"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /pay/channel/get channel GetChannel
//
// Get channel  | 获得支付渠道
//
// Get channel  | 获得支付渠道
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ChannelReq
//
// Responses:
//  200: ChannelResp

func GetChannelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChannelReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := channel.NewGetChannelLogic(r.Context(), svcCtx)
		resp, err := l.GetChannel(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
