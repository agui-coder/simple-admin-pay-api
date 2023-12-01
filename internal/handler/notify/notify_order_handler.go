package notify

import (
	"bytes"
	"io"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/notify"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /pay/order/notify/order notify NotifyOrder
//
// SubmitPayOrder order information | 更新Order
//
// SubmitPayOrder order information | 更新Order
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: notifyRep
//
// Responses:
//  200: string

func NotifyOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NotifyRep
		bs, err := io.ReadAll(io.LimitReader(r.Body, int64(5<<20))) // default 5MB change the size you want;
		defer r.Body.Close()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bs))
		if err = httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		req.Body = bs
		l := notify.NewNotifyOrderLogic(r.Context(), svcCtx)
		req.Header = r.Header
		resp, err := l.NotifyOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
