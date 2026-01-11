export enum PayStatus {
    Peddling  = 1,// 待支付
    SUCCESS  = 2,// 支付成功
}

export enum OrderStatus {
    PendingPayment = 1, // 待付款 
    PendingService = 2, // 待服务
    InProgress = 3, // 进行中
    Completed = 4, // 已完成
    Cancel = 5, // 已取消
    Refund = 6, // 已退款
}