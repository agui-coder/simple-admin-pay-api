package order

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/order"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /order/page order GetOrderPage
//
// Get Order Page | 获取Order分页列表
//
// Get Order Page | 获取Order分页列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: OrderPageReq
//
// Responses:
//  200: OrderPageResp

func GetOrderPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderPageReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewGetOrderPageLogic(r.Context(), svcCtx)
		resp, err := l.GetOrderPage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
