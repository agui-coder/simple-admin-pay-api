package channel

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/channel"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /pay/channel/create channel CreateChannel
//
// Create channel information | 创建Channel
//
// Create channel information | 创建Channel
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ChannelCreateReq
//
// Responses:
//  200: BaseMsgResp

func CreateChannelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChannelCreateReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := channel.NewCreateChannelLogic(r.Context(), svcCtx)
		resp, err := l.CreateChannel(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
