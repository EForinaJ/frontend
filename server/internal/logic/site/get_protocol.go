package site

import (
	"context"
	"server/internal/dao"
	dao_site "server/internal/type/site/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetProtocol implements service.ISite.
func (s *sSite) GetProtocol(ctx context.Context, key string) (res *dao_site.Protocol, err error) {
	options, err := g.Redis().Get(ctx, key)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
	}
	if !options.IsEmpty() {
		var detail *dao_site.Protocol
		err = options.Scan(&detail)
		if err != nil {
			return nil, utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
		}
		return detail, nil
	}

	obj, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, key).
		Fields(dao.SysConfig.Columns().Name, dao.SysConfig.Columns().Value).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	detail := dao_site.Protocol{
		Title:   gconv.String(obj.GMap().Get(dao.SysConfig.Columns().Name)),
		Content: gconv.String(obj.GMap().Get(dao.SysConfig.Columns().Value)),
	}

	err = g.Redis().SetEX(ctx, key, detail, 600)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_SAVE_ERROR, response.CodeMsg(response.CACHE_SAVE_ERROR))
	}
	return &detail, nil
}
