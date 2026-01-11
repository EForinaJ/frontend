package v1

import (
	dao_account "server/internal/type/account/dao"
	dto_account "server/internal/type/account/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDetailReq struct {
	g.Meta `path:"/account" method:"get" tags:"账户" summary:"获取登录账户信息"`
}
type GetDetailRes struct {
	*dao_account.Detail
}

type EditReq struct {
	g.Meta `path:"/account/edit" method:"post" tags:"账户" summary:"修改用户"`
	*dto_account.Edit
}
type EditRes struct{}
