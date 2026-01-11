// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysWitkeyChangeTitleDao is the data access object for the table sys_witkey_change_title.
type SysWitkeyChangeTitleDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of the current DAO.
	columns SysWitkeyChangeTitleColumns // columns contains all the column names of Table for convenient usage.
}

// SysWitkeyChangeTitleColumns defines and stores column names for the table sys_witkey_change_title.
type SysWitkeyChangeTitleColumns struct {
	Id          string //
	ManageId    string //
	WitkeyId    string //
	AfterGame   string //
	AfterTitle  string //
	BeforeTitle string //
	BeforeGame  string //
	Status      string //
	Reason      string //
	Description string //
	CreateTime  string //
}

// sysWitkeyChangeTitleColumns holds the columns for the table sys_witkey_change_title.
var sysWitkeyChangeTitleColumns = SysWitkeyChangeTitleColumns{
	Id:          "id",
	ManageId:    "manage_id",
	WitkeyId:    "witkey_id",
	AfterGame:   "after_game",
	AfterTitle:  "after_title",
	BeforeTitle: "before_title",
	BeforeGame:  "before_game",
	Status:      "status",
	Reason:      "reason",
	Description: "description",
	CreateTime:  "create_time",
}

// NewSysWitkeyChangeTitleDao creates and returns a new DAO object for table data access.
func NewSysWitkeyChangeTitleDao() *SysWitkeyChangeTitleDao {
	return &SysWitkeyChangeTitleDao{
		group:   "default",
		table:   "sys_witkey_change_title",
		columns: sysWitkeyChangeTitleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysWitkeyChangeTitleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysWitkeyChangeTitleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysWitkeyChangeTitleDao) Columns() SysWitkeyChangeTitleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysWitkeyChangeTitleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysWitkeyChangeTitleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysWitkeyChangeTitleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
