package product

import (
	"context"

	v1 "server/api/product/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetCategoryList(ctx context.Context, req *v1.GetCategoryListReq) (res *v1.GetCategoryListRes, err error) {
	list, err := service.Category().GetList(ctx, req.GameId)
	if err != nil {
		return nil, err
	}
	res = &v1.GetCategoryListRes{
		List: list,
	}
	return
}
