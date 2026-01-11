import { http } from "@/http/http"

export function getOrderList(data:Order.Dto.Query) {
  return http.get<{
    list: Order.Dao.List[],
    total: number
  }>('/order/list',data)
}

export function getOrderDetail(data:{id:number}) {
  return http.get<Order.Dao.Detail>('/order/detail',data)
}



export function postOrderCancel(data: {id:number}) {
  return http.post('/order/cancel', data)
}


export function postOrderDelete(data: {id:number}) {
  return http.post('/order/delete', data)
}


// 订单支付-----------------------------

export function postOrderBalancePay(data: {id:number}) {
  return http.post('/order/balance/pay', data)
}


export function postOrderWechatMiniProgramPay(data: {id:number}) {
  return http.post<Order.Dao.WechatMiniProgram>('/order/wechat/mini/program/pay', data)
}

export function getOrderStatus(data:{id:number}) {
  return http.get<{
    status: number
  }>('/order/status',data)
}


