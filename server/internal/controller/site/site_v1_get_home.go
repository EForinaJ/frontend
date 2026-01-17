package site

import (
	"context"

	v1 "server/api/site/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetHome(ctx context.Context, req *v1.GetHomeReq) (res *v1.GetHomeRes, err error) {
	detail, err := service.Site().GetHome(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetHomeRes{
		Home: detail,
	}
	return
}
