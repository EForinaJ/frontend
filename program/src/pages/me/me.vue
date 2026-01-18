<template>
  <view 
  :style="{minHeight: getWindowHeight()-(getTitleBarHeight()+getStatusBarHeight())+ 'px'}"
  class="bg-#f5f5f5">
    <view class="bg-primary" :style="{height:40+getStatusBarHeight()+getTitleBarHeight()+'px'}"></view>
    
    <view v-if="tokenStore.hasLogin" 
    class="
    pos-relative
    mt-[-10rpx] mx-30rpx flex items-center gap-5">
      <view class="
      pos-absolute top-[-40rpx] 
      size-120rpx rounded-full border-6rpx border-white border-solid">
        <sar-avatar :src="getImageUrl(account.avatar)" size="120rpx" icon-size="48rpx" />
      </view>
      <view class="size-120rpx"></view>
      <view class="flex-1">
        <view class="font-bold text-xl text-text-primary leading-none">
          {{ account.name }}
        </view>
        <view class="text-secondary text-xs mt-10rpx line-clamp-1">{{ account.description == '' ? "Ta什么都没写。" :  account.description}}</view>
      </view>
      <view class="bg-secondary px-28rpx text-xs
      text-white py-8rpx rounded-tl-xl rounded-bl-xl"
      @click="handelPush('/pages-me/setting/index')">
        编辑
      </view>
    </view>
    <view v-if="!tokenStore.hasLogin"
      @click="toLoginPage"
      class="mt-[-10rpx] mx-30rpx flex items-center gap-5 pos-relative" 
    >
      <view class="
      pos-absolute top-[-40rpx] 
      size-120rpx rounded-full border-6rpx border-white border-solid">
        <sar-avatar  size="120rpx" icon-size="48rpx" />
      </view>
      <view class="size-120rpx"></view>
      <view class="flex items-center justify-between">
        <text class="text-text-primary font-700">登录/注册</text>
      </view>
    </view>

    <view v-if="tokenStore.hasLogin" class="mx-30rpx mt-100rpx">
      <view class="bg-primary p-20rpx 
      rounded-tr-[130rpx]
      rounded-lg">
        <view class="w-full pos-relative">
          <image class="pos-absolute right-[10rpx] top-[-140rpx]  w-[180rpx] h-[285rpx]" src="~@/static/images/xl.png"/>
        </view>
        <view class="text-[36rpx] font-bold text-white">
          我的余额
        </view>
        <view class="flex items-center mt-2">
          <view class="text-3xl text-white font-bold">
            {{ account.balance }}<text class="text-xs">{{ site.symbol }}</text>
          </view>
        </view>
      </view>
    </view>

    <view class="mt-40rpx mx-30rpx">
      <view class="p-30rpx bg-secondary rounded-lg">
        <view class="flex items-center justify-between">
          <text class="text-white text-[36rpx] font-bold leading-none">我的订单</text>
        </view>
        <view class="mt-30rpx flex items-center justify-between">
          <view @click="handelPush(`/pages-me/order/list?tab=${0}`)" class="flex items-center justify-center flex-col gap-2">
            <sar-cool-icon  size="80rpx" shape="oval" background="#892fe8" color="#fff">
              <i class="iconfont icon-jisudaozhang  text-xl"></i>
            </sar-cool-icon>
            <view class="text-white text-xs font-700">
              全部
            </view>
          </view>
          
          <view 
          @click="handelPush(`/pages-me/order/list?tab=${OrderStatus.PendingPayment}`)"
          class="flex items-center justify-center flex-col gap-2">
            <sar-cool-icon  size="80rpx" shape="oval" background="#892fe8" color="#fff">
              <i class="iconfont icon-daifukuan  text-xl"></i>
            </sar-cool-icon>
            <view class="text-white text-xs font-700">
              待支付
            </view>
          </view>

          <view @click="handelPush(`/pages-me/order/list?tab=${OrderStatus.PendingService}`)"
          class="flex items-center justify-center flex-col gap-2">
            <sar-cool-icon  size="80rpx" shape="oval" background="#892fe8" color="#fff">
              <i class="iconfont icon-daifahuo  text-xl"></i>
            </sar-cool-icon>
            <view class="text-white text-xs font-700">
              待服务
            </view>
          </view>

          <view @click="handelPush(`/pages-me/order/list?tab=${OrderStatus.InProgress}`)"
          class="flex items-center justify-center flex-col gap-2">
            <sar-cool-icon  size="80rpx" shape="oval" background="#892fe8" color="#fff">
              <i class="iconfont icon-haoyouxiadan  text-xl"></i>
            </sar-cool-icon>
            <view class="text-white text-xs font-700">
              进行中
            </view>
          </view>

          <view @click="handelPush(`/pages-me/order/list?tab=${OrderStatus.Completed}`)"
          class="flex items-center justify-center flex-col gap-2">
            <sar-cool-icon  size="80rpx" shape="oval" background="#892fe8" color="#fff">
              <i class="iconfont icon-yiwancheng  text-xl"></i>
            </sar-cool-icon>
            <view class="text-white text-xs font-700">
              已完成
            </view>
          </view>
        </view>
      </view>
    </view>

    <view class="mt-40rpx mx-30rpx">
      <view class="p-30rpx
      bg-secondary flex items-center justify-between rounded-lg">
        <view @click="handelPush(`/pages-protocol/about/index`)"
         class="flex-col flex items-center justify-center gap-2">
          <i class="iconfont text-white icon-a-016_yiwen text-[48rpx]"></i>
          <text class="text-white text-xs">
            关于我们
          </text>
        </view>
        <view @click="handelPush(`/pages-protocol/policy/index`)"
         class="flex-col flex items-center justify-center gap-2">
          <i class="iconfont text-white icon-a-016_tianxiedizhi-37 text-[48rpx]"></i>
          <text class="text-white text-xs">
            隐私政策
          </text>
        </view>
        <view @click="handelPush(`/pages-protocol/policy/index`)"
         class="flex-col flex items-center justify-center gap-2">
          <i class="iconfont text-white icon-a-016_lianxikefu text-[48rpx]"></i>
          <text class="text-white text-xs">
            客服帮助
          </text>
        </view>
        <view @click="handelPush(`/pages-protocol/policy/index`)"
         class="flex-col flex items-center justify-center gap-2">
          <i class="iconfont text-white icon-a-016_fenxiang text-[48rpx]"></i>
          <text class="text-white text-xs">
            邀请分享
          </text>
        </view>
      </view>
    </view>

  </view>
</template>


<script lang="ts" setup>
import { storeToRefs } from 'pinia'
import { useAccountStore } from '@/store'
import { useTokenStore } from '@/store/token'
import { getStatusBarHeight, getTitleBarHeight, getWindowHeight } from '@/utils/systemInfo';
import { toLoginPage } from "@/utils/toLoginPage";
import { useSiteStore } from "@/store/site";
import { OrderStatus } from '@/enum/status';
import { getImageUrl } from '@/utils/imageHelper';
definePage({
  style: {
    navigationBarTitleText: '我的',
    navigationStyle:'custom'
  },
})
const {siteInfo:site} = useSiteStore()
const accountStore = useAccountStore()
const tokenStore = useTokenStore()
const { accountInfo:account } = storeToRefs(accountStore)



const handelPush = (url:string) => {
  uni.navigateTo({url: url})
}

onShareAppMessage((option)=>{
    return{
      title: site.title,
      path: '/pages/index/index'
    }
  }
)

// #ifdef MP-WEIXIN
onShareTimeline(()=>{
  return {
    title: site.title,
    path: '/pages/index/index'
  }
})
// #endif
</script>


<style lang="scss" scoped>
.sad{
  background-color: red;
}
</style>
