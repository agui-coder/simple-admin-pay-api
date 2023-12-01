package channel

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/channel"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /pay/channel/get-enable-code-list channel GetEnableChannelCodeList
//
// Get channel by ID | 通过ID获得指定应用的开启的支付渠道编码列表
//
// Get channel by ID | 通过ID获得指定应用的开启的支付渠道编码列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: ChannelListResp

func GetEnableChannelCodeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := channel.NewGetEnableChannelCodeListLogic(r.Context(), svcCtx)
		resp, err := l.GetEnableChannelCodeList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
