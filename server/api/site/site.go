// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package site

import (
	"context"

	"server/api/site/v1"
)

type ISiteV1 interface {
	GetSite(ctx context.Context, req *v1.GetSiteReq) (res *v1.GetSiteRes, err error)
}
