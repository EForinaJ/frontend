package site

import (
	"context"

	v1 "server/api/site/v1"
	"server/internal/consts"
	"server/internal/service"
)

func (c *ControllerV1) GetAboutUs(ctx context.Context, req *v1.GetAboutUsReq) (res *v1.GetAboutUsRes, err error) {
	detail, err := service.Site().GetProtocol(ctx, consts.AboutUs)
	if err != nil {
		return nil, err
	}
	res = &v1.GetAboutUsRes{
		Protocol: detail,
	}
	return
}
