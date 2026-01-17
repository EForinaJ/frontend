// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package product

import (
	"context"

	"server/api/product/v1"
)

type IProductV1 interface {
	GetCategoryList(ctx context.Context, req *v1.GetCategoryListReq) (res *v1.GetCategoryListRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	GetWxMiniProgramCode(ctx context.Context, req *v1.GetWxMiniProgramCodeReq) (res *v1.GetWxMiniProgramCodeRes, err error)
}
