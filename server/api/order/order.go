// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package order

import (
	"context"

	"server/api/order/v1"
)

type IOrderV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
}
