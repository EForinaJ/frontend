package category

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_category "server/internal/type/category/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.ICategory.
func (s *sCategory) GetList(ctx context.Context, gameId int64) (res []*dao_category.List, err error) {
	m := dao.SysCategory.Ctx(ctx).
		Where(dao.SysCategory.Columns().GameId, gameId).
		OrderDesc(dao.SysOrder.Columns().CreateTime)
	var list []*entity.SysCategory
	err = m.Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_category.List, len(list))
	for i, v := range list {
		var entity *dao_category.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		res[i] = entity
	}
	return
}
