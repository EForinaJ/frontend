import { http } from '@/http/http'
import { IAuthLoginRes } from './types/login'


/**
 * 登录表单
 */
export interface ILoginForm {
    phone: string
    password: string
}

/**
 * 用户登录
 * @param loginForm 登录表单
 */
export function login(loginForm: ILoginForm) {
  return http.post<IAuthLoginRes>('/login', loginForm)
}


/**
 * 微信登录
 * @param params 微信登录参数，包含code
 * @returns Promise 包含登录结果
 */
export function wxLogin(data: { code: string }) {
  return http.post<IAuthLoginRes>('/program/code/login', {code:data.code, provider:'weixin'})
}

/**
 * 微信手机号登录
 * @param params 微信登录参数，包含code
 * @returns Promise 包含登录结果
 */
export function wxPhoneLogin(data: { code: string, phoneCode:string }) {
  return http.post<IAuthLoginRes>('/program/phone/login', {code:data.code, phoneCode:data.phoneCode, provider:'weixin'})
}



