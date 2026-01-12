package v1

import (
	dao_order "server/internal/type/order/dao"
	dto_order "server/internal/type/order/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/order/list" method:"get" tags:"订单" summary:"订单列表"`
	*dto_order.Query
}
type GetListRes struct {
	Total int               `json:"total" dc:"总数"`
	List  []*dao_order.List `json:"list" dc:"订单列表"`
}

type GetDetailReq struct {
	g.Meta `path:"/order/detail" method:"get" tags:"订单" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_order.Detail
}
