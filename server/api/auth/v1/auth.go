package v1

import (
	dao_auth "server/internal/type/auth/dao"
	dto_auth "server/internal/type/auth/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type ProgramCodeLoginReq struct {
	g.Meta `path:"/program/code/login" method:"post" tags:"登录" summary:"小程序登录快捷登录"`
	*dto_auth.ProgramCodeLogin
}
type ProgramCodeLoginRes struct {
	*dao_auth.LoginRes
}

type ProgramPhoneLoginReq struct {
	g.Meta `path:"/program/phone/login" method:"post" tags:"登录" summary:"小程序登录手机号快捷登录"`
	*dto_auth.ProgramPhoneLogin
}
type ProgramPhoneLoginRes struct {
	*dao_auth.LoginRes
}
