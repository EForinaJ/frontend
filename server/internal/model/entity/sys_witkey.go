// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWitkey is the golang structure for table sys_witkey.
type SysWitkey struct {
	Id          int64       `json:"id"          orm:"id"          description:"用户ID"` // 用户ID
	Name        string      `json:"name"        orm:"name"        description:""`     //
	TitleId     int64       `json:"titleId"     orm:"title_id"    description:""`     //
	GameId      int64       `json:"gameId"      orm:"game_id"     description:""`     //
	Phone       string      `json:"phone"       orm:"phone"       description:""`     //
	Avatar      string      `json:"avatar"      orm:"avatar"      description:""`     //
	Address     string      `json:"address"     orm:"address"     description:""`     //
	Salt        string      `json:"salt"        orm:"salt"        description:""`     //
	Password    string      `json:"password"    orm:"password"    description:""`     //
	Album       string      `json:"album"       orm:"album"       description:""`     //
	Sex         int         `json:"sex"         orm:"sex"         description:""`     //
	Birthday    *gtime.Time `json:"birthday"    orm:"birthday"    description:""`     //
	Rate        int         `json:"rate"        orm:"rate"        description:""`     //
	Commission  float64     `json:"commission"  orm:"commission"  description:""`     //
	Description string      `json:"description" orm:"description" description:""`     //
	Status      int         `json:"status"      orm:"status"      description:""`     //
	LoginIp     string      `json:"loginIp"     orm:"login_ip"    description:""`     //
	LoginTime   *gtime.Time `json:"loginTime"   orm:"login_time"  description:""`     //
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time" description:"创建时间"` // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time" description:"更新时间"` // 更新时间
}
