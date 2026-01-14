<template>
  <view class="bg-[#f2f4f3]"
  :style="{minHeight: getWindowHeight()-(getTitleBarHeight()+getStatusBarHeight())+ 'px'}">
      <view  class="
      bg-g w-full px-30rpx fixed box-border z-33">
        <view :style="{height:getStatusBarHeight()+'px'}" ></view>
        <view :style="{height:getTitleBarHeight()+'px'}" >
          <!-- <image class="w-[180rpx] h-full" src="~@/static/images/xl.png"/> -->
          <view class="text-white text-3xl h-full
            tracking-wider 
            font-bold  flex items-center
          ">
            PRO玩家
          </view>
        </view>
        <view class="pos-relative" :style="{height:getTitleBarHeight()+'px'}" >
          <view class="flex left-0 items-center gap-3 pos-absolute bottom-[-50rpx]">
            <view class=" flex items-center justify-center
            h-90rpx w-90rpx bg-[#232323] rounded-lg border-[#fff] border-6rpx border-solid
            text-3xl text-[#e7e6f9]">欢</view>

            <view class=" flex items-center justify-center
            h-90rpx w-90rpx bg-[#232323] rounded-lg border-[#fff] border-6rpx border-solid
            text-3xl text-[#e7e6f9]">迎</view>

            <view class=" flex items-center justify-center
            h-90rpx w-90rpx bg-[#232323] rounded-lg border-[#fff] border-6rpx border-solid
            text-3xl text-[#e7e6f9]">光</view>

            <view class=" flex items-center justify-center
            h-90rpx w-90rpx bg-[#232323] rounded-lg border-[#fff] border-6rpx border-solid
            text-3xl text-[#e7e6f9]">临</view>
          </view>
        </view>
      </view>
      <view :style="{height:40+getStatusBarHeight()+getTitleBarHeight()+getTitleBarHeight()+'px'}"></view>

      <view class="mx-30rpx">
        <swiper
         class="h-200rpx"
          autoplay
          :interval="3000"
        >
          <swiper-item class="slide slide1">1</swiper-item>
          <swiper-item class="slide slide2">2</swiper-item>
          <swiper-item class="slide slide3">3</swiper-item>
        </swiper>
      </view>

      <!-- <view class="mt-40rpx mx-30rpx">
          <view class="bg-[#693ff8] rounded-3xl p-20rpx">
            sd
          </view>
      </view> -->

      <view class="mx-30rpx mt-40rpx"> 
        <view class=" bg-white  px-40rpx pb-60rpx pt-40rpx
        border-[#000] border-6rpx border-solid inner-shadow
         rounded-lg">
          <view class="flex flex-wrap gap-8">
            <view class=" pos-relative flex items-center justify-center flex-col">
              <sar-cool-icon shape="oval" background="#892fe8" color="#fff">
                <image class="size-60rpx mb-20rpx" src="~@/static/images/sanjz.png"/>
              </sar-cool-icon>
              <view class="w-full flex items-center justify-center
              pos-absolute bottom-[-30rpx] z-3
              bg-[#000] text-white text-xs px-10rpx py-8rpx font-bold rounded-full">
                三角洲
              </view>
            </view>
          </view>
        </view>
      </view>

      <view class="mx-30rpx mt-40rpx"> 
        <view class=" bg-white h-200rpx
        border-[#000] border-6rpx border-solid inner-shadow
         rounded-lg">
          
        </view>
      </view>

  </view>
</template>

<script lang="ts" setup>
import { getStatusBarHeight, getTitleBarHeight, getWindowHeight } from '@/utils/systemInfo'

defineOptions({
  name: 'Home',
})
definePage({
  // 使用 type: "home" 属性设置首页，其他页面不需要设置，默认为page
  type: 'home',
  style: {
    // 'custom' 表示开启自定义导航栏，默认 'default'
    navigationStyle: 'custom',
    navigationBarTitleText: '首页',
  },
})

const description = ref(
  'unibest 是一个集成了多种工具和技术的 uniapp 开发模板，由 uniapp + Vue3 + Ts + Vite5 + UnoCss + VSCode 构建，模板具有代码提示、自动格式化、统一配置、代码片段等功能，并内置了许多常用的基本组件和基本功能，让你编写 uniapp 拥有 best 体验。',
)
console.log('index/index 首页打印了')

onLoad(() => {
  console.log('测试 uni API 自动引入: onLoad')
})
</script>


<style lang="scss" scoped>
.inner-shadow {
  box-shadow: 10rpx 10rpx 0rpx rgba(0, 0, 0, 1); /* 内阴影 */
}

.bg-g {
  background:
    linear-gradient(to right,  #693ff8 40%,#693ff8);
}
.box-g {
  background:
    linear-gradient(to bottom, transparent, #fdfdfd 500rpx),
    linear-gradient(to right,  #f5f5f5,#f5f5f5);
}

.slide {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 120rpx;
  font-weight: bold;
  color: var(--sar-white);
  border-radius: 16rpx;
}
.slide1 {
  background-color: rgba(var(--sar-red-rgb), 0.3);
}
.slide2 {
  background-color: rgba(var(--sar-green-rgb), 0.3);
}
.slide3 {
  background-color: rgba(var(--sar-blue-rgb), 0.3);
}
</style>
  
