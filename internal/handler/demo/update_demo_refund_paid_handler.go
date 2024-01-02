package demo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/demo"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /demo-refund/update-paid demo UpdateDemoRefundPaid
//
// Update demoRefund status | 更新退款订单支付状态
//
// Update demoRefund status | 更新退款订单支付状态
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PayRefundNotifyReq
//
// Responses:
//  200: BaseMsgResp

func UpdateDemoRefundPaidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PayRefundNotifyReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := demo.NewUpdateDemoRefundPaidLogic(r.Context(), svcCtx)
		resp, err := l.UpdateDemoRefundPaid(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
