package demo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/demo"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /demo-order/page demo GetDemoOrderPage
//
// Get demoOrder Page information | 获得示例订单列表
//
// Get demoOrder Page information | 获得示例订单列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PageInfo
//
// Responses:
//  200: DemoOrderListResp

func GetDemoOrderPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := demo.NewGetDemoOrderPageLogic(r.Context(), svcCtx)
		resp, err := l.GetDemoOrderPage(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
