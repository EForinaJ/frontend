package v1

import (
	dao_category "server/internal/type/category/dao"
	dao_product "server/internal/type/product/dao"
	dto_product "server/internal/type/product/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type GetCategoryListReq struct {
	g.Meta `path:"/product/category/list" method:"get" tags:"商品" summary:"商品分类列表"`
	GameId int64 `p:"gameid" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetCategoryListRes struct {
	List []*dao_category.List `json:"list" dc:"分类列表"`
}

type GetListReq struct {
	g.Meta `path:"/product/list" method:"get" tags:"商品" summary:"商品列表"`
	*dto_product.Query
}
type GetListRes struct {
	Total int                 `json:"total" dc:"总数"`
	List  []*dao_product.List `json:"list" dc:"商品列表"`
}

type GetDetailReq struct {
	g.Meta `path:"/product/detail" method:"get" tags:"商品" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_product.Detail
}
