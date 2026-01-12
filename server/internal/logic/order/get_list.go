package order

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_order "server/internal/type/order/dao"
	dto_order "server/internal/type/order/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IOrder.
func (s sOrder) GetList(ctx context.Context, req *dto_order.Query) (total int, res []*dao_order.List, err error) {
	m := dao.SysOrder.Ctx(ctx).
		Page(req.Page, req.Limit).
		Where(dao.SysOrder.Columns().UserId, ctx.Value("userId")).
		OrderDesc(dao.SysOrder.Columns().CreateTime)

	if req.Code != "" {
		m = m.Where(dao.SysOrder.Columns().Code, req.Code)
	}
	if req.Status != 0 {
		m = m.Where(dao.SysOrder.Columns().Status, req.Status)
	}

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var list []*entity.SysOrder
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_order.List, len(list))
	for i, v := range list {
		var entity *dao_order.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		// 商品内容
		product, err := dao.SysProduct.Ctx(ctx).
			Where(dao.SysProduct.Columns().Id, v.ProductId).
			Fields(
				dao.SysProduct.Columns().Name,
				dao.SysProduct.Columns().Pic,
				dao.SysProduct.Columns().GameId,
				dao.SysProduct.Columns().CategoryId,
			).
			One()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Product = &dao_order.Product{
			Id:   v.ProductId,
			Name: gconv.String(product.GMap().Get(dao.SysProduct.Columns().Name)),
			Pic:  gconv.String(product.GMap().Get(dao.SysProduct.Columns().Pic)),
		}

		game, err := dao.SysGame.Ctx(ctx).
			Where(dao.SysGame.Columns().Id, product.GMap().Get(dao.SysProduct.Columns().GameId)).
			Value(dao.SysGame.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Product.Game = game.String()

		category, err := dao.SysCategory.Ctx(ctx).
			Where(dao.SysCategory.Columns().Id, product.GMap().Get(dao.SysProduct.Columns().CategoryId)).
			Value(dao.SysCategory.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Product.Category = category.String()

		res[i] = entity
	}

	return
}
