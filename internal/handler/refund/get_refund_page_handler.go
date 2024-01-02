package refund

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/refund"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /refund/page refund GetRefundPage
//
// Get Refund Page | 获取Refund分页列表
//
// Get Refund Page | 获取Refund分页列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: RefundPageReq
//
// Responses:
//  200: RefundPageResp

func GetRefundPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RefundPageReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := refund.NewGetRefundPageLogic(r.Context(), svcCtx)
		resp, err := l.GetRefundPage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
