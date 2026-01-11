package router

import (
	"server/internal/controller/account"
	"server/internal/controller/auth"
	"server/internal/controller/site"
	"server/internal/controller/upload"
	"server/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func LoadRouter(s *ghttp.Server) {
	s.Group("/api/v1/frontend", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			site.NewV1(),
			auth.NewV1(),
		).Middleware(middleware.Response)
		group.Bind(
			account.NewV1(),
			upload.NewV1(),
		).Middleware(middleware.Auth).Middleware(middleware.Response)
	})
}
