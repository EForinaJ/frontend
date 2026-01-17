package auth

import (
	"context"
	"fmt"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/lib/jwt"
	dao_auth "server/internal/type/auth/dao"
	dto_auth "server/internal/type/auth/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// ProgramPhoneLogin implements service.IAuth.
func (s *sAuth) ProgramPhoneLogin(ctx context.Context, req *dto_auth.ProgramPhoneLogin) (res *dao_auth.LoginRes, err error) {
	if req.Provider == consts.WEIXIN {
		obj, err := s.weiXinPhone(ctx, req)
		if err != nil {
			return nil, err
		}

		return obj, nil
	}
	return
}

// weiXin implements service.IAuth.
func (s *sAuth) weiXinPhone(ctx context.Context, req *dto_auth.ProgramPhoneLogin) (res *dao_auth.LoginRes, err error) {

	options, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.WechatMiniProgramSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	json, err := gjson.DecodeToJson(options)
	if err != nil {
		return nil, utils_error.Err(response.FAILD, response.CodeMsg(response.FAILD))
	}

	appid := json.Get("appId")
	secret := json.Get("secret")

	wxAccounttoken := ""
	wxAccountTokenRdObj, err := g.Redis().Get(ctx, "wx_program_access_token")
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
	}
	if !wxAccountTokenRdObj.IsEmpty() {
		wxAccounttoken = wxAccountTokenRdObj.String()
	} else {
		// 请求微信access_token
		tokenUrl := fmt.Sprintf(
			"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
			appid, secret,
		)
		// 创建一个HTTP客户端请求对象
		resp, err := g.Client().Get(ctx, tokenUrl)
		if err != nil {
			return nil, utils_error.Err(response.EXCEPTION, response.CodeMsg(response.EXCEPTION))
		}
		defer resp.Close()
		result, err := gjson.DecodeToJson(resp.ReadAllString())
		if err != nil {
			return nil, utils_error.Err(response.EXCEPTION, response.CodeMsg(response.EXCEPTION))
		}
		wxAccounttoken = result.Get("access_token").String()
	}
	err = g.Redis().SetEX(ctx, "wx_program_access_token", wxAccounttoken, 7000)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_SAVE_ERROR, response.CodeMsg(response.CACHE_SAVE_ERROR))
	}

	// 请求微信phone
	phoneUrl := fmt.Sprintf(
		"https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s",
		wxAccounttoken,
	)
	phoneCode := fmt.Sprintf(
		`{"code":"%s"}`,
		req.PhoneCode,
	)
	phoneResp, err := g.Client().Post(ctx, phoneUrl, phoneCode)
	if err != nil {
		return nil, utils_error.Err(response.EXCEPTION, response.CodeMsg(response.EXCEPTION))
	}
	defer phoneResp.Close()
	phoneResult, err := gjson.DecodeToJson(phoneResp.ReadAllString())
	if err != nil {
		return nil, utils_error.Err(response.EXCEPTION, response.CodeMsg(response.EXCEPTION))
	}

	// 请求微信接口
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		appid, secret, req.Code,
	)
	// 创建一个HTTP客户端请求对象
	codeResp, err := g.Client().Get(ctx, url)
	if err != nil {
		return nil, utils_error.Err(response.EXCEPTION, response.CodeMsg(response.EXCEPTION))
	}
	defer codeResp.Close()
	codeResult, err := gjson.DecodeToJson(codeResp.ReadAllString())
	if err != nil {
		return nil, utils_error.Err(response.EXCEPTION, response.CodeMsg(response.EXCEPTION))
	}

	// 判断用户是否注册过
	userDetail, err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Phone, phoneResult.GetJson("phone_info").Get("phoneNumber")).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if userDetail.IsEmpty() {
		entity := g.Map{
			dao.SysUser.Columns().Name:       "新用户" + grand.S(6),
			dao.SysUser.Columns().Birthday:   gtime.NewFromTimeStamp(946656000000 / 1000).UTC(),
			dao.SysUser.Columns().CreateTime: gtime.Now(),
			dao.SysUser.Columns().UpdateTime: gtime.Now(),
			dao.SysUser.Columns().WxOpenId:   codeResult.Get("openid"),
			dao.SysUser.Columns().Phone:      phoneResult.GetJson("phone_info").Get("phoneNumber"),
		}
		newSalt := grand.S(6)
		newToken := consts.SYSTEMNAME + grand.S(6) + newSalt
		newToken = gmd5.MustEncryptString(newToken)
		entity[dao.SysUser.Columns().Salt] = newSalt
		entity[dao.SysUser.Columns().Password] = newToken
		entity[dao.SysUser.Columns().Sex] = 3

		//  获取配置
		userInfo, err := dao.SysConfig.Ctx(ctx).
			Where(dao.SysConfig.Columns().Key, consts.UserSetting).
			Value(dao.SysConfig.Columns().Value)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		userJosn := gjson.New(userInfo)
		//  获取默认头像和封面
		entity[dao.SysUser.Columns().Avatar] = userJosn.Get("avatar")
		entity[dao.SysUser.Columns().Cover] = userJosn.Get("cover")
		entity[dao.SysUser.Columns().Status] = consts.Enable
		entity[dao.SysUser.Columns().LevelId] = userJosn.Get("levelId").Int64()
		exp, err := dao.SysLevel.Ctx(ctx).
			Where(dao.SysLevel.Columns().Id, userJosn.Get("levelId").Int64()).
			Value(dao.SysLevel.Columns().Experience)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity[dao.SysUser.Columns().Experience] = exp.Int()

		rs, err := dao.SysUser.Ctx(ctx).Data(&entity).Insert()
		if err != nil {
			return nil, utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
		}
		rid, err := rs.LastInsertId()
		if err != nil {
			return nil, utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
		}

		token, err := jwt.GenToken(rid, gconv.String(entity[dao.SysUser.Columns().Name]))
		if err != nil {
			return nil, utils_error.Err(response.LOGIN_ERROR, response.CodeMsg(response.LOGIN_ERROR))
		}

		expire, err := g.Cfg().Get(ctx, "jwt.expire")
		if err != nil {
			return nil, utils_error.Err(response.LOGIN_ERROR, response.CodeMsg(response.LOGIN_ERROR))
		}

		loginRes := dao_auth.LoginRes{
			Token:  token,
			Expire: expire.Int(),
		}
		return &loginRes, nil
	} else {
		entity := g.Map{
			dao.SysUser.Columns().LoginIp:   g.RequestFromCtx(ctx).GetClientIp(),
			dao.SysUser.Columns().LoginTime: gtime.Now(),
		}
		if gconv.String(userDetail.GMap().Get(dao.SysUser.Columns().WxOpenId)) == "" {
			entity[dao.SysUser.Columns().WxOpenId] = codeResult.Get("openid")
		}
		_, err = dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Phone, userDetail.GMap().Get(dao.SysUser.Columns().Id)).
			Data(entity).Update()
		if err != nil {
			return nil, utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
		}

		token, err := jwt.GenToken(gconv.Int64(userDetail.GMap().Get(dao.SysUser.Columns().Id)), gconv.String(userDetail.GMap().Get(dao.SysUser.Columns().Name)))
		if err != nil {
			return nil, utils_error.Err(response.LOGIN_ERROR, response.CodeMsg(response.LOGIN_ERROR))
		}

		expire, err := g.Cfg().Get(ctx, "jwt.expire")
		if err != nil {
			return nil, utils_error.Err(response.LOGIN_ERROR, response.CodeMsg(response.LOGIN_ERROR))
		}
		loginRes := dao_auth.LoginRes{
			Token:  token,
			Expire: gconv.Int(expire),
		}

		return &loginRes, nil
	}
}
