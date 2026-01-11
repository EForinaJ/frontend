package auth

import (
	"context"

	v1 "server/api/auth/v1"
	"server/internal/service"
)

func (c *ControllerV1) ProgramPhoneLogin(ctx context.Context, req *v1.ProgramPhoneLoginReq) (res *v1.ProgramPhoneLoginRes, err error) {
	detail, err := service.Auth().ProgramPhoneLogin(ctx, req.ProgramPhoneLogin)
	if err != nil {
		return nil, err
	}
	res = &v1.ProgramPhoneLoginRes{
		LoginRes: detail,
	}
	return
}
