// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysWitkeyOnboardingDao is the data access object for the table sys_witkey_onboarding.
type SysWitkeyOnboardingDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of the current DAO.
	columns SysWitkeyOnboardingColumns // columns contains all the column names of Table for convenient usage.
}

// SysWitkeyOnboardingColumns defines and stores column names for the table sys_witkey_onboarding.
type SysWitkeyOnboardingColumns struct {
	Id         string // 用户ID
	ManageId   string //
	Name       string //
	TitleId    string //
	GameId     string //
	Phone      string //
	Salt       string //
	Password   string //
	Birthday   string //
	Address    string //
	Sex        string //
	Status     string //
	CreateTime string // 创建时间
}

// sysWitkeyOnboardingColumns holds the columns for the table sys_witkey_onboarding.
var sysWitkeyOnboardingColumns = SysWitkeyOnboardingColumns{
	Id:         "id",
	ManageId:   "manage_id",
	Name:       "name",
	TitleId:    "title_id",
	GameId:     "game_id",
	Phone:      "phone",
	Salt:       "salt",
	Password:   "password",
	Birthday:   "birthday",
	Address:    "address",
	Sex:        "sex",
	Status:     "status",
	CreateTime: "create_time",
}

// NewSysWitkeyOnboardingDao creates and returns a new DAO object for table data access.
func NewSysWitkeyOnboardingDao() *SysWitkeyOnboardingDao {
	return &SysWitkeyOnboardingDao{
		group:   "default",
		table:   "sys_witkey_onboarding",
		columns: sysWitkeyOnboardingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysWitkeyOnboardingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysWitkeyOnboardingDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysWitkeyOnboardingDao) Columns() SysWitkeyOnboardingColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysWitkeyOnboardingDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysWitkeyOnboardingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysWitkeyOnboardingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
