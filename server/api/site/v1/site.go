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

type GetAboutUsReq struct {
	g.Meta `path:"/site/about/us" method:"get" tags:"站点" summary:"关于我们"`
}
type GetAboutUsRes struct {
	*dao_site.Protocol
}

type GetPrivacyPolicyReq struct {
	g.Meta `path:"/site/privacy/policy" method:"get" tags:"站点" summary:"隐私政策"`
}
type GetPrivacyPolicyRes struct {
	*dao_site.Protocol
}

type GetUserAgreementReq struct {
	g.Meta `path:"/site/user/agreement" method:"get" tags:"站点" summary:"用户协议"`
}
type GetUserAgreementRes struct {
	*dao_site.Protocol
}

type GetOrderInstructionsReq struct {
	g.Meta `path:"/site/order/instructions" method:"get" tags:"站点" summary:"下单须知"`
}
type GetOrderInstructionsRes struct {
	*dao_site.Protocol
}
