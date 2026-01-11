// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWitkeyOnboarding is the golang structure of table sys_witkey_onboarding for DAO operations like Where/Data.
type SysWitkeyOnboarding struct {
	g.Meta     `orm:"table:sys_witkey_onboarding, do:true"`
	Id         interface{} // 用户ID
	ManageId   interface{} //
	Name       interface{} //
	TitleId    interface{} //
	GameId     interface{} //
	Phone      interface{} //
	Salt       interface{} //
	Password   interface{} //
	Birthday   *gtime.Time //
	Address    interface{} //
	Sex        interface{} //
	Status     interface{} //
	CreateTime *gtime.Time // 创建时间
}
