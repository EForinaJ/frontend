// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWitkeyChangeTitle is the golang structure for table sys_witkey_change_title.
type SysWitkeyChangeTitle struct {
	Id          int64       `json:"id"          orm:"id"           description:""` //
	ManageId    int64       `json:"manageId"    orm:"manage_id"    description:""` //
	WitkeyId    int64       `json:"witkeyId"    orm:"witkey_id"    description:""` //
	AfterGame   int64       `json:"afterGame"   orm:"after_game"   description:""` //
	AfterTitle  int64       `json:"afterTitle"  orm:"after_title"  description:""` //
	BeforeTitle int64       `json:"beforeTitle" orm:"before_title" description:""` //
	BeforeGame  int64       `json:"beforeGame"  orm:"before_game"  description:""` //
	Status      int         `json:"status"      orm:"status"       description:""` //
	Reason      string      `json:"reason"      orm:"reason"       description:""` //
	Description string      `json:"description" orm:"description"  description:""` //
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"  description:""` //
}
