import { http } from '@/http/http'

/**
 * 获取用户信息
 */
export function getSite() {
  return http.get<Site.Dao.Info>('/site')
}


/**
 * 获取字典信息
 */
export function getSiteDictData(code:string) {
  return http.get<{
    list: Site.Dao.DictData[]
}>('/site/dict/data',{code})
}