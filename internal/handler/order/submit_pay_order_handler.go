package order

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/order"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /order/submit order SubmitPayOrder
//
// SubmitPayOrder order information | 提交退款Order
//
// SubmitPayOrder order information | 提交退款Order
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: OrderSubmitReq
//
// Responses:
//  200: OrderSubmitResp

func SubmitPayOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderSubmitReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := order.NewSubmitPayOrderLogic(r.Context(), svcCtx)
		resp, err := l.SubmitPayOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
