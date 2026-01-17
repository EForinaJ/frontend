package site

import (
	"context"

	v1 "server/api/site/v1"
	"server/internal/consts"
	"server/internal/service"
)

func (c *ControllerV1) GetOrderInstructions(ctx context.Context, req *v1.GetOrderInstructionsReq) (res *v1.GetOrderInstructionsRes, err error) {
	detail, err := service.Site().GetProtocol(ctx, consts.OrderInstructions)
	if err != nil {
		return nil, err
	}
	res = &v1.GetOrderInstructionsRes{
		Protocol: detail,
	}
	return
}
