package dto_auth

type ProgramCodeLogin struct {
	Provider string `v:"required#请输入登录服务商" p:"provider" dc:"登录服务商"`
	Code     string `v:"required#请输入用户登录码" p:"code" dc:"用户登录码"`
}

type ProgramPhoneLogin struct {
	Provider  string `v:"required#请输入登录服务商" p:"provider" dc:"登录服务商"`
	PhoneCode string `v:"required#请输入手机号授权码" p:"phoneCode" dc:"手机号授权"`
	Code      string `v:"required#请输入用户登录码" p:"code" dc:"用户登录码"`
}
