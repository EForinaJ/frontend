package dao_category

type List struct {
	Id   int64  `json:"id" dc:"ID"`
	Name string `json:"name" dc:"分类名称"`
}
