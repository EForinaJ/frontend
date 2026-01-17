package product

import (
	"context"
	"fmt"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/do"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type qRCodeData struct {
	Scene      int64  `json:"scene"`
	Page       string `json:"page"`
	EnvVersion string `json:"env_version"`
	IsHyaline  bool   `json:"is_hyaline"`
	CheckPath  bool   `json:"check_path"`
}

// GetWxMiniProgramCode implements service.IProduct.
func (s *sProduct) GetWxMiniProgramCode(ctx context.Context, id int64) (res string, err error) {
	obj, err := dao.SysProduct.Ctx(ctx).
		WherePri(id).
		Fields(dao.SysProduct.Columns().Name, dao.SysProduct.Columns().WxMiniProgramQrCode).
		One()
	if err != nil {
		return "", utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if obj.IsEmpty() {
		return "", utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}
	if gconv.String(obj.GMap().Get(dao.SysProduct.Columns().WxMiniProgramQrCode)) != "" {
		return gconv.String(obj.GMap().Get(dao.SysProduct.Columns().WxMiniProgramQrCode)), nil
	}

	options, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.WechatMiniProgramSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return "", utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	json, err := gjson.DecodeToJson(options)
	if err != nil {
		return "", utils_error.Err(response.FAILD, response.CodeMsg(response.FAILD))
	}

	appid := json.Get("appId")
	secret := json.Get("secret")

	wxAccounttoken := ""
	wxAccountTokenRdObj, err := g.Redis().Get(ctx, "wx_program_access_token")
	if err != nil {
		return "", utils_error.Err(response.CACHE_READ_ERROR, response.CodeMsg(response.CACHE_READ_ERROR))
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
			return "", utils_error.Err(response.EXCEPTION, response.CodeMsg(response.EXCEPTION))
		}
		defer resp.Close()
		result, err := gjson.DecodeToJson(resp.ReadAllString())
		if err != nil {
			return "", utils_error.Err(response.EXCEPTION, response.CodeMsg(response.EXCEPTION))
		}
		wxAccounttoken = result.Get("access_token").String()
	}
	err = g.Redis().SetEX(ctx, "wx_program_access_token", wxAccounttoken, 7000)
	if err != nil {
		return "", utils_error.Err(response.CACHE_SAVE_ERROR, response.CodeMsg(response.CACHE_SAVE_ERROR))
	}

	// 请求微信phone
	qRCodeUrl := fmt.Sprintf(
		"https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s",
		wxAccounttoken,
	)

	qRCodeDatas := qRCodeData{
		Scene:      id,
		Page:       "pages/product/detail",
		EnvVersion: "develop",
		IsHyaline:  true,
		CheckPath:  false,
	}
	qrjson := gjson.MustEncode(qRCodeDatas)
	qrCodeResp, err := g.Client().Post(ctx, qRCodeUrl, qrjson)
	if err != nil {
		return "", utils_error.Err(response.EXCEPTION, response.CodeMsg(response.EXCEPTION))
	}
	defer qrCodeResp.Close()
	qrResult := qrCodeResp.ReadAll()

	fileConfig, err := dao.SysConfig.Ctx(ctx).
		Where(dao.SysConfig.Columns().Key, consts.FileSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return "", utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	fjson := gjson.New(fileConfig)
	path := fjson.Get("path").String()
	savePath := "../../public/" + path + "/" + gtime.Date()
	orName := gconv.String(obj.GMap().Get(dao.SysProduct.Columns().Name)) + "微信小程序为毛"
	name := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36)+grand.S(6)) + ".png"
	filePath := gfile.Join(savePath, name)

	// 确保目录存在
	if !gfile.Exists(savePath) {
		gfile.Mkdir(savePath)
	}

	// 将二进制数据写入文件
	if err = gfile.PutBytes(filePath, qrResult); err != nil {
		return "", utils_error.Err(response.FILE_SAVE_ERROR, response.CodeMsg(response.FILE_SAVE_ERROR))
	}

	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return "", utils_error.Err(response.DB_TX_ERROR, response.CodeMsg(response.DB_TX_ERROR))
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	entites := &do.SysMedia{
		Path:       savePath[5:] + "/" + name,
		Ext:        ".png",
		MediaType:  "png",
		Name:       name,
		CreateTime: gtime.Now(),
		OrName:     orName,
		Size:       len(qrResult),
	}
	_, err = tx.Model(dao.SysMedia.Table()).Data(&entites).Insert()
	if err != nil {
		return "", utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	_, err = tx.Model(dao.SysProduct.Table()).
		Where(dao.SysProduct.Columns().Id, id).
		Data(g.Map{
			dao.SysProduct.Columns().WxMiniProgramQrCode: savePath[5:] + "/" + name,
		}).Update()
	if err != nil {
		return "", utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	return savePath[5:] + "/" + name, nil
}
