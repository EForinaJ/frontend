package service

import (
	"context"
	dao_order "server/internal/type/order/dao"
	dto_order "server/internal/type/order/dto"
)

// 定义显示接口
type IOrder interface {
	GetList(ctx context.Context, req *dto_order.Query) (total int, res []*dao_order.List, err error)
	GetDetail(ctx context.Context, id int64) (res *dao_order.Detail, err error)

	// Start(ctx context.Context, id int64) (err error)
	// Complete(ctx context.Context, id int64) (err error)
	// Settlement(ctx context.Context, req *dto_order.Settlement) (err error)

	// CheckSettlement(ctx context.Context, req *dto_order.Settlement) (err error)
	// CheckComplete(ctx context.Context, id int64) (err error)
	// CheckStart(ctx context.Context, id int64) (err error)
}

// 定义接口变量
var localOrder IOrder

// 定义一个获取接口的方法
func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

// 定义一个接口实现的注册方法
func RegisterOrder(i IOrder) {
	localOrder = i
}
