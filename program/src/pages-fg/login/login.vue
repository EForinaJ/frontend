<template>
  <view class="page-container w-full" :style="{minHeight: getWindowHeight()+ 'px'}">
    <NavBar :search="false"/>
    <!-- <view class="px-[30rpx] flex justify-end">
      <image class="w-[240rpx] h-[380rpx]" src="~@/static/images/xl.png"/>
    </view>
    <view class="px-[30rpx]">
      <view class="mt-50rpx font-700 text-2xl">
        欢迎来到<text>'PRO玩家'</text>
      </view>
    </view> -->
    <view 
    :style="{minHeight: (getWindowHeight()-getStatusBarHeight()-getTitleBarHeight()) + 'px'}"
    class="flex-col flex items-center justify-center">
      <view class="w-full">
        <view class="px-100rpx flex flex-col gap-40rpx">
          <sar-button 
          v-if="!hasBound"
          round
          :disabled="!checked"
          open-type="getPhoneNumber" @getphonenumber="handlePgoneLogin"
          size="large" background="#101010" type="default" >
            <i class="iconfont icon-shouji text-white text-xl mr-20rpx"></i>
            授权绑定手机
          </sar-button>
          <sar-button 
          v-if="hasBound"
          round
          @click="handleCodeLogin"
          size="large" type="default" >
            <i class="iconfont icon-weichat text-white text-xl mr-20rpx"></i>
            一键快捷登录
          </sar-button>
        </view>
        <view 
        class="mt-[60rpx] flex justify-center items-center gap-2">
          <sar-checkbox size="32rpx" v-model:checked="checked"></sar-checkbox>
          <view class="text-xs text-text-secondary ">
            登录即同意<text class="text-text-primary font-700 mx-8rpx">用户协议</text>和<text class=" ml-8rpx text-text-primary font-700">隐私政策</text>
          </view>
        </view>
      </view>
    </view>
    
  
    <view class="fixed bottom-[300rpx] w-full flex items-center justify-center">
      <sar-button @click="handleBack" color="#67676c" type="pale-text">暂不登录</sar-button>
    </view> 
  </view>
</template>
  
<script lang="ts" setup>
import NavBar from "@/components/nav-bar.vue";
import { useTokenStore } from '@/store/token'
import { getStatusBarHeight, getTitleBarHeight, getWindowHeight } from '@/utils/systemInfo';
definePage({
  style: {
    navigationBarTitleText: '登录',
    navigationStyle: 'custom'
  },
})
const tokenStore = useTokenStore()
const handlePgoneLogin = async (e:any) =>{
    if (tokenStore.hasLogin) {
      return
    }
    if (!checked.value) {
      uni.showToast({
        title: '请勾选授权服务',
      })
      return
    }
    await tokenStore.wxPhoneLogin(e.detail.code)
    // uni.navigateBack()
}
const handleBack = async () =>{
  if (tokenStore.hasLogin) {
    tokenStore.logout()
  }
  uni.navigateBack()
}

const checked = ref<boolean>(false)
const handleCodeLogin = async () =>{
  if (tokenStore.hasLogin) {
    return
  }
  await tokenStore.wxCodeLogin()
  // uni.navigateBack()
}

const hasBound = computed(()=>{
  return uni.getStorageSync('has_bound_phone')
});

onLaunch(()=>{
  if (tokenStore.hasLogin) {
    uni.navigateBack()
    return
  }
})
</script>
  
  
<style lang="scss" scoped>

</style>
