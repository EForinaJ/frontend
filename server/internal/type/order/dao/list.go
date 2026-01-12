package dao_order

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id          int64       `json:"id" dc:"ID"`
	Code        string      `json:"code" dc:"订单编号"`
	Product     *Product    `json:"product" dc:"订单商品"`
	Quantity    int         `json:"quantity" dc:"数量"`
	Unit        string      `json:"unit" dc:"单位"`
	UnitPrice   float64     `json:"unitPrice" dc:"单价"`
	TotalAmount float64     `json:"totalAmount" dc:"订单总额"`
	Status      int         `json:"status" dc:"订单状态"`
	CreateTime  *gtime.Time `json:"createTime" dc:"下单时间"`
}
