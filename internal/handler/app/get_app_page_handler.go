package app

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/app"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /pay/app/page app GetAppPage
//
// Get app page | 获取App列表分页
//
// Get app page | 获取App列表分页
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: AppListReq
//
// Responses:
//  200: AppListResp

func GetAppPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := app.NewGetAppPageLogic(r.Context(), svcCtx)
		resp, err := l.GetAppPage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
