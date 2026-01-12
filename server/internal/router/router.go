package router

import (
	"server/internal/controller/account"
	"server/internal/controller/auth"
	"server/internal/controller/order"
	"server/internal/controller/site"
	"server/internal/controller/upload"
	"server/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func LoadRouter(s *ghttp.Server) {
	s.Group("/api/v1/frontend", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Middleware(ghttp.MiddlewareCORS)
		group.Middleware(middleware.FixedWindowMiddleware)
		group.Bind(
			site.NewV1(),
			auth.NewV1(),
		).Middleware(middleware.Response)

		// 读取操作
		group.Bind(
			account.NewV1().GetDetail,

			order.NewV1().GetList,
			order.NewV1().GetDetail,
		).Middleware(middleware.Auth).Middleware(middleware.Response)

		// 写入操作
		group.Bind(
			upload.NewV1(),
			account.NewV1().Edit,
		).Middleware(middleware.FixedWindowActionMiddleware).
			Middleware(middleware.Auth).
			Middleware(middleware.Response)
	})
}
