package site

import (
	"context"

	v1 "server/api/site/v1"
	"server/internal/consts"
	"server/internal/service"
)

func (c *ControllerV1) GetUserAgreement(ctx context.Context, req *v1.GetUserAgreementReq) (res *v1.GetUserAgreementRes, err error) {
	detail, err := service.Site().GetProtocol(ctx, consts.UserAgreement)
	if err != nil {
		return nil, err
	}
	res = &v1.GetUserAgreementRes{
		Protocol: detail,
	}
	return
}
