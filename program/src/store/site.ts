
import { getSite } from '@/api/site'
import { defineStore } from 'pinia'
import { ref } from 'vue'


// 初始化状态
const siteInfoState: Site.Dao.Info = {
  title: "",    // 网站标题
  domain: "",      // 网站域名
  logo: "",       // 网站Logo图片地址
  icon: "",        // 网站图标地址
  description: "", // 网站描述
  address: "",    // 网站地址
  icp: "",         // ICP备案号
  copyright: "",   // 版权信息
  fileSize : 0,
  imageSize: 0,
  fileType: [],
  imageType: [],
  symbol: "",   // 货币类型
}
export const useSiteStore = defineStore(
  'site',
  () => {
    // 定义信息
    const siteInfo = ref<Site.Dao.Info>({ ...siteInfoState })
    // 设置信息
    const setSiteInfo = (val: Site.Dao.Info) => {
      siteInfo.value = val
    }
    // 删除信息
    const clearSiteInfo = () => {
      siteInfo.value = { ...siteInfoState }
      uni.removeStorageSync('site')
    }
    /**
     * 获取信息
     */
    const fetchSiteInfo = async () => {
      const res = await getSite()
      setSiteInfo(res)
      return res
    }


    return {
      siteInfo,
      clearSiteInfo,
      fetchSiteInfo,
      setSiteInfo,
    }
  },
  {
    persist: true,
  },
)
