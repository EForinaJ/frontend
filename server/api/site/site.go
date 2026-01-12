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
	GetAboutUs(ctx context.Context, req *v1.GetAboutUsReq) (res *v1.GetAboutUsRes, err error)
	GetPrivacyPolicy(ctx context.Context, req *v1.GetPrivacyPolicyReq) (res *v1.GetPrivacyPolicyRes, err error)
	GetUserAgreement(ctx context.Context, req *v1.GetUserAgreementReq) (res *v1.GetUserAgreementRes, err error)
}
