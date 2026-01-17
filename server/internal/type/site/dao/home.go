package dao_site

type Home struct {
	GameList    []*Game        `json:"gameList" dc:"游戏列表"`
	ProductList []*ProductList `json:"productList" dc:"商品列表"`
}

type Game struct {
	Id   int64  `json:"id" dc:"游戏ID"`
	Name string `json:"name" dc:"游戏名称"`
	Pic  string `json:"pic" dc:"游戏图片"`
}

type ProductList struct {
	Game     *Game      `json:"game" dc:"游戏领域"`
	Products []*Product `json:"products" dc:"游戏列表"`
}
type Product struct {
	Id   int64  `json:"id" dc:"商品id"`
	Name string `json:"name" dc:"商品名称"`
	Pic  string `json:"pic" dc:"商品图片"`
}
