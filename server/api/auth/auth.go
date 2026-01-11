// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"server/api/auth/v1"
)

type IAuthV1 interface {
	ProgramCodeLogin(ctx context.Context, req *v1.ProgramCodeLoginReq) (res *v1.ProgramCodeLoginRes, err error)
	ProgramPhoneLogin(ctx context.Context, req *v1.ProgramPhoneLoginReq) (res *v1.ProgramPhoneLoginRes, err error)
}
