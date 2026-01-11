<script setup lang="ts">
import { onHide, onLaunch, onShow } from '@dcloudio/uni-app'
import { navigateToInterceptor } from '@/router/interceptor'
import { useAccountStore, useTokenStore } from './store';
import { useSiteStore } from './store/site';

onLaunch(async (options) => {
  uni.loadFontFace({
  family: 'HarmonyOS_Sans_SC_Regular',
    source: 'url("/static/fonts/HarmonyOS_Sans_SC_Regular.ttf")',
    success() {
        console.log('success')
    }
  })

  const siteStore = useSiteStore()
  await siteStore.fetchSiteInfo()

  const tokenStore = useTokenStore()
  if (tokenStore.hasLogin) {
    const accountStore = useAccountStore()
    await accountStore.fetchAccountInfo()

    if (accountStore.accountInfo.phone == "") {
      uni.navigateTo({
        url:"/pages-fg/login/login"
      })
      return
    }
  }
  
  
})
onShow((options) => {
  
  // 处理直接进入页面路由的情况：如h5直接输入路由、微信小程序分享后进入等
  // https://github.com/unibest-tech/unibest/issues/192
  if (options?.path) {
    navigateToInterceptor.invoke({ url: `/${options.path}`, query: options.query })
  }
  else {
    navigateToInterceptor.invoke({ url: '/' })
  }
})

</script>

<style lang="scss">
@import 'sard-uniapp/index.scss';

</style>
