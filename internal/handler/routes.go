// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	app "github.com/agui-coder/simple-admin-pay-api/internal/handler/app"
	base "github.com/agui-coder/simple-admin-pay-api/internal/handler/base"
	channel "github.com/agui-coder/simple-admin-pay-api/internal/handler/channel"
	demo "github.com/agui-coder/simple-admin-pay-api/internal/handler/demo"
	notify "github.com/agui-coder/simple-admin-pay-api/internal/handler/notify"
	order "github.com/agui-coder/simple-admin-pay-api/internal/handler/order"
	refund "github.com/agui-coder/simple-admin-pay-api/internal/handler/refund"
	"github.com/agui-coder/simple-admin-pay-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/pay/init/database",
				Handler: base.InitDatabaseHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/app/create",
					Handler: app.CreateAppHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/app/update",
					Handler: app.UpdateAppHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/app/update-status",
					Handler: app.UpdateAppStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/app/delete",
					Handler: app.DeleteAppHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/app/page",
					Handler: app.GetAppPageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/app",
					Handler: app.GetAppByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/pay"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/order/submit",
					Handler: order.SubmitPayOrderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/order/update",
					Handler: order.UpdateOrderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/order/get",
					Handler: order.GetOrderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/order/page",
					Handler: order.GetOrderPageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/order/get-detail",
					Handler: order.GetOrderDetailHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/pay"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/channel/create",
					Handler: channel.CreateChannelHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/channel/update",
					Handler: channel.UpdateChannelHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/channel/delete",
					Handler: channel.DeleteChannelHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/channel/get",
					Handler: channel.GetChannelHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/channel/get-enable-code-list",
					Handler: channel.GetEnableChannelCodeListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/pay"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/notify/order/:channelId",
					Handler: notify.NotifyOrderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/notify/refund",
					Handler: notify.NotifyRefundHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/pay"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/demo-order/create",
					Handler: demo.CreateDemoOrderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/demo-order/page",
					Handler: demo.GetDemoOrderPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/demo-order/get/:id",
					Handler: demo.GetDemoOrderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/demo-order/update-paid",
					Handler: demo.UpdateDemoOrderPaidHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/pay"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/refund/get",
					Handler: refund.GetRefundHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/refund/page",
					Handler: refund.GetRefundPageHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/pay"),
	)
}
