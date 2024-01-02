package notify

import (
	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-pay-api/internal/logic/notify"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"
	"github.com/agui-coder/simple-admin-pay-api/internal/types"
)

// swagger:route post /notify/refund/{channelCode} notify NotifyRefund
//
// Notify order information | 退款回调
//
// Notify order information | 退款回调
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
		err := httpx.ParsePath(r, &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		ctx := r.Context()
		req.R, err = model.ParseRefundNotify(req.ChannelCode, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		l := notify.NewNotifyRefundLogic(ctx, svcCtx)
		resp, err := l.NotifyRefund(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
