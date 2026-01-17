package service

import (
	"context"
	dao_product "server/internal/type/product/dao"
	dto_product "server/internal/type/product/dto"
)

// 定义显示接口
type IProduct interface {
	GetList(ctx context.Context, req *dto_product.Query) (total int, res []*dao_product.List, err error)
	GetDetail(ctx context.Context, id int64) (res *dao_product.Detail, err error)
	GetWxMiniProgramCode(ctx context.Context, id int64) (res string, err error)
}

// 定义接口变量
var localProduct IProduct

// 定义一个获取接口的方法
func Product() IProduct {
	if localProduct == nil {
		panic("implement not found for interface IProduct, forgot register?")
	}
	return localProduct
}

// 定义一个接口实现的注册方法
func RegisterProduct(i IProduct) {
	localProduct = i
}
