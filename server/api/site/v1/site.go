package v1

import (
	dao_site "server/internal/type/site/dao"

	"github.com/gogf/gf/v2/frame/g"
)

// ---------------------- 基础配置
type GetSiteReq struct {
	g.Meta `path:"/site" method:"get" tags:"站点" summary:"获取系统站点信息"`
}
type GetSiteRes struct {
	*dao_site.Detail
}
