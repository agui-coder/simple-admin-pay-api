package app

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/app"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /pay/app/update app UpdateApp
//
// Update app information | 更新App
//
// Update app information | 更新App
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: AppUpdateReq
//
// Responses:
//  200: BaseMsgResp

func UpdateAppHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppUpdateReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := app.NewUpdateAppLogic(r.Context(), svcCtx)
		resp, err := l.UpdateApp(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
