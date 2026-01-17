package dao_product

type List struct {
	Id         int64   `json:"id" dc:"ID"`
	Name       string  `json:"name" dc:"商品名称"`
	Pic        string  `json:"pic" dc:"商品图片"`
	Category   string  `json:"category" dc:"所属分类"`
	Game       string  `json:"game" dc:"所属游戏"`
	Price      float64 `json:"price" dc:"价格"`
	SalesCount int     `json:"salesCount" dc:"销量"`
	Unit       string  `json:"unit" dc:"单位"`
	Rate       int     `json:"rate" dc:"评分"`
}
