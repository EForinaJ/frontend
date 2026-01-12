import { http } from '@/http/http'
import { IAccountInfoRes } from './types/login'
import { IAccountDto } from './types/account'
/**
 * 获取用户信息
 */
export function getAccountInfo() {
  return http.get<IAccountInfoRes>('/account')
}


/**
 * 修改用户信息
 * @param loginForm 用户信息表单
 */
export function postAccountEdit(data: IAccountDto) {
  return http.post('/account/edit', data)
}


// /**

//  */
// export function wxPhoneBind(data: { code: string }) {
//   return http.post<{phone:string}>('/account/wx/program/phone', data)
// }
