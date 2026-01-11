export enum SexType {
    Male = 1,      // 男
    Female = 2,    // 女
    Other = 3,     // 其他
}

export enum RechargeType {
    AlyPay = 1,// 支付宝
    Wechat = 2,// 微信
    Manual = 3,// 人工
}
export enum PayType {
    AlyPay = 1,// 支付宝
    Wechat = 2,// 微信
    Balance = 3,// 余额
}



export enum BillType {
    Recharge = 1, // 账户充值
    Refund  = 2, // 商品退款
    Order  = 3, // 订单支付
    SettlementCommission  = 4, // 结算佣金
}