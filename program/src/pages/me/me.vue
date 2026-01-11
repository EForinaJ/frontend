<template>
  <view 
  :style="{minHeight: getWindowHeight()-(getTitleBarHeight()+getStatusBarHeight())+ 'px'}"
  class="page-container w-full">
    <NavBar :search="false"/>

    <view v-if="tokenStore.hasLogin" class="mx-30rpx flex items-center gap-3">
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
      class="mx-30rpx flex items-center gap-3"
    >
      <sar-avatar  size="108rpx" icon-size="48rpx" />
      <view class="flex items-center justify-between w-full">
        <text class="text-text-primary font-700">登录/注册</text>
      </view>
    </view>

    <view v-if="tokenStore.hasLogin" class="mx-30rpx mt-40rpx">
      <view class=" p-20rpx
      text-[#ece0e0] rounded-tr-lg rounded-tl-lg flex justify-between items-center
      bg-[linear-gradient(300deg,#4c4d51,#2a2a31_15%,#85858a_40%,#393a3c_60%,#393838_80%,#5e5f62_100%)]">
        <view class="flex items-center gap-1">
          <text class="font-700">{{ account.level.name }}</text>
          <sar-image
            :src="account.level.icon"
            width="100rpx"
            height="60rpx"
            mode="aspectFit"
          ></sar-image>
        </view>
        <view @click="handelPush('/pages-me/bill/index')"  class="flex items-center gap-1">
          <text class="text-2xs ">账单明细</text>
          <sar-icon name="caret-right" size="20rpx"/>
        </view>
      </view>
      <view class="bg-white px-40rpx py-20rpx
      rounded-br-lg rounded-bl-lg flex items-center justify-between
      ">
        <view class="flex items-start justify-center flex-col gap-1">
          <view class="text-text-primary text-xl font-700">
            {{ account.balance }} <text class="text-xs">{{ siteStore.siteInfo.symbol }}</text>
          </view>
          <view class="text-text-secondary text-2xs">
            余额
          </view>
        </view>
        <view class="flex items-center">
          <sar-button inline @click="handelPush('/pages-me/recharge/index')" type="pale-text">充值</sar-button>
        </view>
      </view>
    </view>

    <view class="mt-40rpx mx-30rpx">
      <view class="p-30rpx bg-white rounded-lg">
        <view class="flex items-center justify-between">
          <text class="text-text-primary font-bold leading-none">我的订单</text>
          <view @click="handelPush('/pages-me/order/list')" class="flex items-center gap-1">
            <text class="text-text-secondary text-xs leading-none">全部订单</text>
            <sar-icon size="28rpx" class="text-text-secondary" name="right" />
          </view>
        </view>
        <view class="mt-30rpx flex items-center justify-between">
          <view class="flex items-center justify-center flex-col gap-2">
            <sar-cool-icon  size="80rpx" shape="triangle" background="#fce22e" color="#fff">
              <i class="iconfont icon-daifukuan  text-xl"></i>
            </sar-cool-icon>
            <view class="text-text-secondary text-xs font-700">
              待支付
            </view>
          </view>

          <view class="flex items-center justify-center flex-col gap-2">
            <sar-cool-icon  size="80rpx" shape="triangle" background="#fce22e" color="#fff">
              <i class="iconfont icon-daifahuo  text-xl"></i>
            </sar-cool-icon>
            <view class="text-text-secondary text-xs font-700">
              待服务
            </view>
          </view>

          <view class="flex items-center justify-center flex-col gap-2">
            <sar-cool-icon  size="80rpx" shape="triangle" background="#fce22e" color="#fff">
              <i class="iconfont icon-haoyouxiadan  text-xl"></i>
            </sar-cool-icon>
            <view class="text-text-secondary text-xs font-700">
              进行中
            </view>
          </view>

          <view class="flex items-center justify-center flex-col gap-2">
            <sar-cool-icon  size="80rpx" shape="triangle" background="#fce22e" color="#fff">
              <i class="iconfont icon-yiwancheng  text-xl"></i>
            </sar-cool-icon>
            <view class="text-text-secondary text-xs font-700">
              已完成
            </view>
          </view>

          <!-- <view class="flex items-center justify-center flex-col gap-2">
            <sar-cool-icon size="80rpx" shape="triangle" background="#fce22e" color="#fff">
              <i class="iconfont icon-jisudaozhang  text-xl"></i>
            </sar-cool-icon>
            <view class="text-text-secondary text-xs font-700">
              退款/售后
            </view>
          </view> -->
        </view>
      </view>
    </view>

    <view class="mt-40rpx mx-30rpx">
      <view class="bg-white rounded-lg p-10rpx">
        <sar-grid clickable >
          <sar-grid-item>
              <template #text>
                <text class="text-text-secondary">
                  关于我们
                </text>
              </template>
              <template #icon>
                <i class="iconfont text-text-secondary icon-a-016_yiwen text-[48rpx]"></i>
              </template>
          </sar-grid-item>
          <sar-grid-item >
              <template #text>
                <text class="text-text-secondary">
                  平台协议
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
import NavBar from "@/components/nav-bar.vue";
import { storeToRefs } from 'pinia'
import { useAccountStore } from '@/store'
import { useTokenStore } from '@/store/token'
import { getStatusBarHeight, getTitleBarHeight, getWindowHeight } from '@/utils/systemInfo';
import { toLoginPage } from "@/utils/toLoginPage";
import { useSiteStore } from "@/store/site";
definePage({
  style: {
    navigationBarTitleText: '我的',
    navigationStyle:'custom'
  },
})
const siteStore = useSiteStore()
const accountStore = useAccountStore()
const tokenStore = useTokenStore()
// 使用storeToRefs解构accountInfo
const { accountInfo:account } = storeToRefs(accountStore)



const handelPush = (url:string) => {
  uni.navigateTo({url: url})
}


const rightsList  = ref([
        {
          icon: '💰',
          text: '4倍积分',
          bgColor: '#FFF9F0',
          iconBgColor: '#FFE8C5',
          badge: '热'
        },
        {
          icon: '📅',
          text: '周二会员日',
          bgColor: '#F0F9FF',
          iconBgColor: '#D1ECFF',
          badge: '新'
        },
        {
          icon: '🖥️',
          text: '跨屏特权',
          bgColor: '#F0FFF4',
          iconBgColor: '#D1F7C4'
        },
        {
          icon: '🎁',
          text: '生日礼遇',
          bgColor: '#FFF0F5',
          iconBgColor: '#FFD1E0'
        },
        {
          icon: '🚚',
          text: '免运费',
          bgColor: '#F5F0FF',
          iconBgColor: '#E6D1FF'
        },
        {
          icon: '⏰',
          text: '提前购',
          bgColor: '#F0FFF5',
          iconBgColor: '#D1F7E4'
        },
        {
          icon: '💎',
          text: '尊享折扣',
          bgColor: '#FFF5F0',
          iconBgColor: '#FFE6D1'
        },
        {
          icon: '🛡️',
          text: '专属客服',
          bgColor: '#F0F5FF',
          iconBgColor: '#D1E0FF'
        }
      ])

</script>


<style lang="scss" scoped>
  
.balance-bg {
  background: linear-gradient(to right, #ffd1ff,#f5f5f5,#fff 180%);
}
</style>
