package demo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/demo"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /demo-order/update-paid demo UpdateDemoOrderPaid
//
// Update demoOrder status | 更新示例订单支付状态
//
// Update demoOrder status | 更新示例订单支付状态
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PayOrderNotifyReq
//
// Responses:
//  200: BaseMsgResp

func UpdateDemoOrderPaidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PayOrderNotifyReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := demo.NewUpdateDemoOrderPaidLogic(r.Context(), svcCtx)
		resp, err := l.UpdateDemoOrderPaid(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
