package account

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_account "server/internal/type/account/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Edit implements service.IAccount.
func (s *sAccount) Edit(ctx context.Context, req *dto_account.Edit) (err error) {
	entity := g.Map{
		dao.SysUser.Columns().Name:        req.Name,
		dao.SysUser.Columns().Address:     req.Address,
		dao.SysUser.Columns().Birthday:    gtime.NewFromTimeStamp(req.Birthday / 1000).UTC(),
		dao.SysUser.Columns().Avatar:      req.Avatar,
		dao.SysUser.Columns().Sex:         req.Sex,
		dao.SysUser.Columns().Description: req.Description,
		dao.SysUser.Columns().UpdateTime:  gtime.Now(),
	}

	_, err = dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id, ctx.Value("userId")).
		Data(&entity).Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}

	_, err = g.Redis().Del(ctx, consts.Account+gconv.String(ctx.Value("userId")))
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR, response.CodeMsg(response.DB_SAVE_ERROR))
	}
	return
}
