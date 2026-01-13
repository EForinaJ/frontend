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
        <view class=" bg-white 
        border-[#000] border-6rpx border-solid inner-shadow
         rounded-lg">
          <view class=" flex items-center border-b-[#000] border-b-6rpx border-b-solid">
            <view class="bg-black h-100rpx flex items-center px-10rpx">
              <view class=" tracking-[10rpx] text-white text-xl font-bold">
                KOOK
              </view>
            </view>
            <view class="gap-2 flex-1 h-100rpx px-20rpx
            flex justify-around items-center">
              <view 
              v-for="value in 8"
              class=" font-bold
              text-2xl text-[#000]">2</view>
            </view>
          </view>

          <view class="p-20rpx flex items-center gap-3">
            <sar-avatar size="96rpx" src="https://fastly.jsdelivr.net/npm/@sard/assets/pic1.jpg" />
            <view class="">
              <view class="font-bold text-[36rpx]">
                小可
              </view>
              <view class="text-text-secondary font-bold text-sm">
                在线时间: 12:00~24:00
              </view>
            </view>
          </view>
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
</style>
  
