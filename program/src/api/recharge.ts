import { http } from '@/http/http'

export function postWeChatMiniProgramRecharge(data:Recharge.Dto.Create) {
  return http.post<Recharge.Dao.WechatMiniProgram>('/recharge/wechat/mini/program', data)
}


export function getRechargesStatus(data:{id:number}) {
  return http.get<{
    status: number
  }>('/recharge/status',data)
}
