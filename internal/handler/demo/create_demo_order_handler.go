package demo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/demo"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /pay/order/demo/create demo CreateDemoOrder
//
// createDemoOrder demoOrder information | 创建demoOrder
//
// createDemoOrder demoOrder information | 创建demoOrder
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: createDemoOrderReq
//
// Responses:
//  200: BaseMsgResp

func CreateDemoOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateDemoOrderReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := demo.NewCreateDemoOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateDemoOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
