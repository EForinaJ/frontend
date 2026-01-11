// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package account

import (
	"context"

	"server/api/account/v1"
)

type IAccountV1 interface {
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error)
}
