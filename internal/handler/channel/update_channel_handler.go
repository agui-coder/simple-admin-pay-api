package channel

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/channel"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /pay/channel/update channel UpdateChannel
//
// Update channel information | 更新Channel
//
// Update channel information | 更新Channel
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ChannelUpdateReq
//
// Responses:
//  200: BaseMsgResp

func UpdateChannelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChannelUpdateReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := channel.NewUpdateChannelLogic(r.Context(), svcCtx)
		resp, err := l.UpdateChannel(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
