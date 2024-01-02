package demo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/demo"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route get /demo-order/get/{id} demo GetDemoOrder
//
// Get demoOrder information | 获得示例订单
//
// Get demoOrder information | 获得示例订单
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDAtPathReq
//
// Responses:
//  200: DemoOrderInfo

func GetDemoOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDAtPathReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := demo.NewGetDemoOrderLogic(r.Context(), svcCtx)
		resp, err := l.GetDemoOrder(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
