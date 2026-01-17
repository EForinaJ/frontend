package dao_site

type Contact struct {
	Platform string `json:"platform" dc:"联系平台"`
	Number   string `json:"number" dc:"频道号码"`
	Wechat   string `json:"wechat" dc:"客服微信"`
	Icon     string `json:"icon" dc:"平台图标"`
}
