package product

import (
	"context"
	"server/internal/dao"
	dao_product "server/internal/type/product/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IProduct.
func (s sProduct) GetDetail(ctx context.Context, id int64) (res *dao_product.Detail, err error) {
	product, err := dao.SysProduct.Ctx(ctx).
		WherePri(id).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if product.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}

	// 2. 使用结构体转换代替手动映射
	var detail *dao_product.Detail
	if err := gconv.Scan(product.Map(), &detail); err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	game, err := dao.SysGame.Ctx(ctx).
		Where(dao.SysGame.Columns().Id, product.GMap().Get(dao.SysProduct.Columns().GameId)).
		Value(dao.SysGame.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Game = game.String()

	category, err := dao.SysCategory.Ctx(ctx).
		Where(dao.SysCategory.Columns().Id, product.GMap().Get(dao.SysProduct.Columns().CategoryId)).
		Value(dao.SysCategory.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Category = category.String()

	return detail, nil
}
