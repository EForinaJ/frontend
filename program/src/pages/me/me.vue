<template>
  <view 
  :style="{minHeight: getWindowHeight()-(getTitleBarHeight()+getStatusBarHeight())+ 'px'}"
  class="page-x w-full">
    <view class="fill" :style="{height:getStatusBarHeight()+getTitleBarHeight()+'px'}"></view>
    
    <view v-if="tokenStore.hasLogin" class="mt-20rpx mx-30rpx flex items-center gap-3">
      <view class="size-120rpx rounded-full border-4rpx border-white border-solid">
        <sar-avatar :src="account.avatar" size="120rpx" icon-size="48rpx" />
      </view>
      <view class="flex-1">
        <view class="font-bold text-xl text-text-primary leading-none">
          {{ account.name }}
        </view>
        <view class="text-text-secondary text-xs mt-10rpx line-clamp-1">{{ account.description == '' ? "Ta什么都没写。" :  account.description}}</view>
      </view>
      <view @click="handelPush('/pages-me/setting/index')">
        <i class="iconfont icon-a-016_shezhi text-text-secondary text-[48rpx]"></i>
      </view>
    </view>
    <view v-if="!tokenStore.hasLogin"
      @click="toLoginPage"
      class="mt-20rpx mx-30rpx flex items-center gap-3" 
    >
      <sar-avatar  size="108rpx" icon-size="48rpx" />
      <view class="flex items-center justify-between">
        <text class="text-text-primary font-700">登录/注册</text>
      </view>
    </view>

    <view v-if="tokenStore.hasLogin" class="mx-30rpx mt-100rpx">
      <view class="balance-bg p-20rpx 
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
            {{ account.balance }}<text class="text-xs">{{ siteStore.siteInfo.symbol }}</text>
          </view>
        </view>
      </view>
    </view>

    <view class="mt-40rpx mx-30rpx">
      <view class="p-30rpx bg-white rounded-lg">
        <view class="flex items-center justify-between">
          <text class="text-text-primary text-[36rpx] font-bold leading-none">我的订单</text>
        </view>
        <view class="mt-30rpx flex items-center justify-between">
          <view @click="handelPush(`/pages-me/order/list?tab=${0}`)" class="flex items-center justify-center flex-col gap-2">
            <sar-cool-icon  size="80rpx" shape="oval" background="#e7d5fa" color="#892fe8">
              <i class="iconfont icon-jisudaozhang  text-xl"></i>
            </sar-cool-icon>
            <view class="text-text-secondary text-xs font-700">
              全部
            </view>
          </view>
          
          <view 
          @click="handelPush(`/pages-me/order/list?tab=${OrderStatus.PendingPayment}`)"
          class="flex items-center justify-center flex-col gap-2">
            <sar-cool-icon  size="80rpx" shape="oval" background="#e7d5fa" color="#892fe8">
              <i class="iconfont icon-daifukuan  text-xl"></i>
            </sar-cool-icon>
            <view class="text-text-secondary text-xs font-700">
              待支付
            </view>
          </view>

          <view @click="handelPush(`/pages-me/order/list?tab=${OrderStatus.PendingService}`)"
          class="flex items-center justify-center flex-col gap-2">
            <sar-cool-icon  size="80rpx" shape="oval" background="#e7d5fa" color="#892fe8">
              <i class="iconfont icon-daifahuo  text-xl"></i>
            </sar-cool-icon>
            <view class="text-text-secondary text-xs font-700">
              待服务
            </view>
          </view>

          <view @click="handelPush(`/pages-me/order/list?tab=${OrderStatus.InProgress}`)"
          class="flex items-center justify-center flex-col gap-2">
            <sar-cool-icon  size="80rpx" shape="oval" background="#e7d5fa" color="#892fe8">
              <i class="iconfont icon-haoyouxiadan  text-xl"></i>
            </sar-cool-icon>
            <view class="text-text-secondary text-xs font-700">
              进行中
            </view>
          </view>

          <view @click="handelPush(`/pages-me/order/list?tab=${OrderStatus.Completed}`)"
          class="flex items-center justify-center flex-col gap-2">
            <sar-cool-icon  size="80rpx" shape="oval" background="#e7d5fa" color="#892fe8">
              <i class="iconfont icon-yiwancheng  text-xl"></i>
            </sar-cool-icon>
            <view class="text-text-secondary text-xs font-700">
              已完成
            </view>
          </view>
        </view>
      </view>
    </view>

    <view class="mt-40rpx mx-30rpx">
      <view class="bg-white rounded-lg p-10rpx">
        <sar-grid clickable >
          <sar-grid-item @click="handelPush(`/pages-protocol/about/index`)">
              <template #text>
                <text class="text-text-secondary">
                  关于我们
                </text>
              </template>
              <template #icon>
                <i class="iconfont text-text-secondary icon-a-016_yiwen text-[48rpx]"></i>
              </template>
          </sar-grid-item>
          <sar-grid-item @click="handelPush(`/pages-protocol/policy/index`)">
              <template #text>
                <text class="text-text-secondary">
                  隐私政策
                </text>
              </template>
              <template #icon>
                <i class="iconfont text-text-secondary icon-a-016_tianxiedizhi-37 text-[48rpx]"></i>
              </template>
          </sar-grid-item>
          <sar-grid-item >
              <template #text>
                <text class="text-text-secondary">
                  客服帮助
                </text>
              </template>
              <template #icon>
                <i class="iconfont text-text-secondary icon-a-016_lianxikefu text-[48rpx]"></i>
              </template>
          </sar-grid-item>
          <sar-grid-item >
              <template #text>
                <text class="text-text-secondary">
                  邀请分享
                </text>
              </template>
              <template #icon>
                <i class="iconfont text-text-secondary icon-a-016_fenxiang text-[48rpx]"></i>
              </template>
          </sar-grid-item>
        </sar-grid>
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
definePage({
  style: {
    navigationBarTitleText: '我的',
    navigationStyle:'custom'
  },
})
const siteStore = useSiteStore()
const accountStore = useAccountStore()
const tokenStore = useTokenStore()
const { accountInfo:account } = storeToRefs(accountStore)



const handelPush = (url:string) => {
  uni.navigateTo({url: url})
}


</script>


<style lang="scss" scoped>
.page-x {
  background:
    linear-gradient(to bottom, transparent, #f5f5f5 400rpx),
    linear-gradient(to right,  #bb313e25,
       #bb313e25,
       #d7222925,
       #dd4a1625,
       #e4761525,
       #f5c50025,
       #f0e92725,
       #b1ce2425,
       #48a93525,
       #03944525,
       #157c4f25,
       #176a5825,
       #1b556325,
       #1d386f25,
       #1d386f25,
       #20277825,
       #52266325,
       #8a244b25);
}
.balance-bg {
  background: linear-gradient(to right, #693FF8 ,#caaff3,#f5f5f5 );
}
</style>
