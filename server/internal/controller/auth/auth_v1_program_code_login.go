package auth

import (
	"context"

	v1 "server/api/auth/v1"
	"server/internal/service"
)

func (c *ControllerV1) ProgramCodeLogin(ctx context.Context, req *v1.ProgramCodeLoginReq) (res *v1.ProgramCodeLoginRes, err error) {
	detail, err := service.Auth().ProgramCodeLogin(ctx, req.ProgramCodeLogin)
	if err != nil {
		return nil, err
	}
	res = &v1.ProgramCodeLoginRes{
		LoginRes: detail,
	}
	return
}
