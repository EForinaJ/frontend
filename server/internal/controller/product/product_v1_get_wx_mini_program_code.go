package product

import (
	"context"

	v1 "server/api/product/v1"
	"server/internal/service"
)

func (c *ControllerV1) GetWxMiniProgramCode(ctx context.Context, req *v1.GetWxMiniProgramCodeReq) (res *v1.GetWxMiniProgramCodeRes, err error) {
	code, err := service.Product().GetWxMiniProgramCode(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetWxMiniProgramCodeRes{
		QrCode: code,
	}
	return
}
