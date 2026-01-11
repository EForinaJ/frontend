package service

import (
	"context"
	dao_auth "server/internal/type/auth/dao"
	dto_auth "server/internal/type/auth/dto"
)

// 定义显示接口
type IAuth interface {
	// 小程序快捷登录
	ProgramCodeLogin(ctx context.Context, req *dto_auth.ProgramCodeLogin) (res *dao_auth.LoginRes, err error)
	ProgramPhoneLogin(ctx context.Context, req *dto_auth.ProgramPhoneLogin) (res *dao_auth.LoginRes, err error)
}

// 定义接口变量
var localAuth IAuth

// 定义一个获取接口的方法
func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

// 定义一个接口实现的注册方法
func RegisterAuth(i IAuth) {
	localAuth = i
}
