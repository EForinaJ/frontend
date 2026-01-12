package site

import (
	"context"

	v1 "server/api/site/v1"
	"server/internal/consts"
	"server/internal/service"
)

func (c *ControllerV1) GetPrivacyPolicy(ctx context.Context, req *v1.GetPrivacyPolicyReq) (res *v1.GetPrivacyPolicyRes, err error) {
	detail, err := service.Site().GetProtocol(ctx, consts.PrivacyPolicy)
	if err != nil {
		return nil, err
	}
	res = &v1.GetPrivacyPolicyRes{
		Protocol: detail,
	}
	return
}
