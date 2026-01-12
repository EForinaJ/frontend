import { http } from '@/http/http'

export function getSite() {
  return http.get<Site.Dao.Info>('/site')
}

export function getSiteAboutUs() {
  return http.get<Site.Dao.Protocol>('/site/about/us')
}

export function getSitePrivacyPolicy() {
  return http.get<Site.Dao.Protocol>('/site/privacy/policy')
}

export function getSiteUserAgreement() {
  return http.get<Site.Dao.Protocol>('/site/user/agreement')
}

