package site

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_site "server/internal/type/site/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetHome implements service.ISite.
func (s *sSite) GetHome(ctx context.Context) (res *dao_site.Home, err error) {
	options, err := g.Redis().Get(ctx, "frontent_home")
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
	}
	if !options.IsEmpty() {
		err = options.Scan(&res)
		if err != nil {
			return nil, utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
		}
		return
	}

	var list []*entity.SysGame
	err = dao.SysGame.Ctx(ctx).Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	gameList := make([]*dao_site.Game, len(list))
	productList := make([]*dao_site.ProductList, len(list))
	for i, v := range list {
		var gameEntity *dao_site.Game
		err = gconv.Scan(v, &gameEntity)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		gameList[i] = gameEntity

		var productObjList []*entity.SysProduct
		err = dao.SysProduct.Ctx(ctx).
			Page(1, 10).
			Where(dao.SysProduct.Columns().GameId, v.Id).
			OrderDesc(dao.SysProduct.Columns().SalesCount).
			Scan(&productObjList)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		productxList := make([]*dao_site.Product, len(productObjList))
		for pi, pv := range productObjList {
			var productEntity *dao_site.Product
			err = gconv.Scan(pv, &productEntity)
			if err != nil {
				return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
			}
			productxList[pi] = productEntity
		}
		productListEntity := dao_site.ProductList{
			Game:     gameEntity,
			Products: productxList,
		}
		productList[i] = &productListEntity
	}

	detail := dao_site.Home{
		GameList:    gameList,
		ProductList: productList,
	}

	err = g.Redis().SetEX(ctx, "frontent_home", detail, 600)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_SAVE_ERROR, response.CodeMsg(response.CACHE_SAVE_ERROR))
	}
	return &detail, nil
}
