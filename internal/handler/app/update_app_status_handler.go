package app

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/app"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /pay/app/update-status app UpdateAppStatus
//
// Update app information | 更新App状态
//
// Update app information | 更新App状态
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: AppUpdateStatusReq
//
// Responses:
//  200: BaseMsgResp

func UpdateAppStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppUpdateStatusReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := app.NewUpdateAppStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateAppStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
