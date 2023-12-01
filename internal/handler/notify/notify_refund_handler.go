package notify

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/notify"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /pay/order/notify/refund notify NotifyRefund
//
// Update order information | 更新Order
//
// Update order information | 更新Order
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: notifyRep
//
// Responses:
//  200: string

func NotifyRefundHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NotifyRep
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := notify.NewNotifyRefundLogic(r.Context(), svcCtx)
		resp, err := l.NotifyRefund(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
