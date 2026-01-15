<template>
  <view class="bg-[#fdfdfd]"
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
      </view>
      <view :style="{height:20+getStatusBarHeight()+getTitleBarHeight()+'px'}"></view>

      <view class="mx-30rpx">
        <view class="bg-[#663df3] rounded-lg mt-20rpx">
          <view class="pos-relative" :style="{height:getTitleBarHeight()+'px'}" >
            <view class="w-full flex justify-center items-center gap-3 pos-absolute top-[-50rpx]">
              <view class=" flex items-center justify-center text-shadow
              h-90rpx w-90rpx bg-[#232323] rounded-lg border-[#fff] border-6rpx border-solid
              text-3xl text-[#ffff]">欢</view>

              <view class=" flex items-center justify-center text-shadow
              h-90rpx w-90rpx bg-[#232323] rounded-lg border-[#fff] border-6rpx border-solid
              text-3xl text-[#e7e6f9]">迎</view>

              <view class=" flex items-center justify-center text-shadow
              h-90rpx w-90rpx bg-[#232323] rounded-lg border-[#fff] border-6rpx border-solid
              text-3xl text-[#e7e6f9]">光</view>

              <view class=" flex items-center justify-center text-shadow
              h-90rpx w-90rpx bg-[#232323] rounded-lg border-[#fff] border-6rpx border-solid
              text-3xl text-[#e7e6f9]">临</view>
            </view>
          </view>

          <view class="flex gap-1 px-40rpx pb-40rpx">
            <view class="rounded-tl-xl rounded-bl-xl pos-relative
            flex-1 bg-[#1B0660]  clip-l p-20rpx">
              <view class="stroke-join 
              text-white tracking-[4rpx]
              text-[36rpx] font-700" data-content="加入我们">
                加入我们
              </view>
              <view class="w-fit text-shadow-c py-8rpx
              bg-black rounded-3xl px-30rpx mt-10rpx
               text-sm text-white font-semibold">
                GO
              </view>
              <view class=" pos-absolute bottom-[10rpx] right-[40rpx]">
                <image class="w-80rpx h-110rpx" src="~@/static/images/join.png"/>
              </view>
            </view>
            <view class="flex-col flex gap-3
            flex-1  clip-r">
              <view class="flex-1 flex items-center justify-center  gap-2
              bg-[#1B0660] rounded-tr-xl rounded-br-xl">
                <image class="size-36rpx" src="~@/static/images/kook.png"/>
                <view class="text-white text-sm">在线客服</view>
              </view>
              <view class="flex-1 text-white text-sm font-semibold
              flex items-center justify-center
              bg-[#1B0660] rounded-tr-xl rounded-br-xl">
                下单须知
              </view>
            </view>
          </view>
        </view>
      </view>


      <!-- <view class="mx-30rpx mt-40rpx">
        <view class="text-xl font-bold pos-relative
        w-fit 
        after:content-[''] after:w-full 
        after:bg-gradient-to-r after:from-[#f03740] after:to-transparent
        after:inline-block
        after:rounded-lg after:h-10rpx 
        after:absolute after:bottom-4rpx after:left-0
        ">
           热门商品
        </view>
        <view class="shadow-2xl 
        bg-red-3 rounded-3xl mt-20rpx">
          <view class="flex justify-between items-center px-30rpx py-16rpx">
            <view class="flex items-center gap-2">
              <view class="flex items-center justify-center">
                <image class="size-40rpx" src="~@/static/images/sanjz.png"/>
              </view>
              <view class="font-bold text-[36rpx] text-[#ffff]">
                三角洲行动
              </view>
            </view>
            <view class="flex gap-2 items-center bg-white py-8rpx px-12rpx rounded-2xl">
              <text class="text-xs">全部</text>
              <view class="size-28rpx bg-black rounded-full flex items-center justify-center">
                <view class="i-carbon:chevron-right text-white text-xs" />
              </view>
            </view>
          </view>
          <view class="bg-white rounded-3xl p-40rpx">
            <sar-scroll-list
             scrollbar-bg="rgba(var(--sar-danger-rgb), 0.2)"
            thumb-bg="var(--sar-danger)"
            >
                <view class="flex gap-3">
                  <view v-for="item in 8" class="
                  shadow-2xl
                  w-260rpx rounded-xl bg-black">
                    <sar-image class="rounded-xl"
                      src="https://fastly.jsdelivr.net/npm/@sard/assets/images/cat1.jpg"
                      width="100%"
                      height="180rpx"
                    />
                    <view class="px-16rpx py-16rpx">
                      <text class="text-sm line-clamp-1 leading-none overflow-hidden
                    text-white ">保底300w双子星猛攻单</text>
                    </view>
                  </view>
                </view>
            </sar-scroll-list>
          </view>
        </view>
      </view> -->

      <view class="mx-30rpx mt-40rpx">
        <view class="text-xl font-bold pos-relative
        w-fit 
        after:content-[''] after:w-full 
        after:bg-gradient-to-r after:from-[#663df3] after:to-transparent
        after:inline-block
        after:rounded-lg after:h-10rpx 
        after:absolute after:bottom-4rpx after:left-0
        ">
          人气服务
        </view>
        <view class="
        bg-[#663df3] rounded-lg p-20rpx mt-20rpx">
          <view class="rounded-lg">
            <view class="bg-white flex
             rounded-tl-3xl h-80rpx
            ">
              <view class="bg-[#663df3]">
                <view class="w-fit pl-30rpx pr-70rpx flex items-center justify-center  gap-2
                bg-white rounded-tl-3xl rounded-tr-3xl h-full">
                  <view class="rounded-full px-8rpx flex items-center justify-center">
                    <image class="size-40rpx" src="~@/static/images/sanjz.png"/>
                  </view>
                  <view class="font-bold text-[36rpx]">
                    三角洲行动
                  </view>
                </view>
              </view>

              <view class="flex-1
              bg-[#663df3] h-full flex justify-end items-center
              rounded-bl-3xl">
                <view class="flex gap-2 items-center bg-white py-6rpx px-16rpx
                 rounded-t-2xl rounded-br-2xl">
                    <text class="text-xs">全部</text>
                    <view class="size-28rpx bg-black rounded-full flex items-center justify-center">
                      <view class="i-carbon:chevron-right text-white text-xs" />
                    </view>
                </view>
              </view>
            </view>
            <view class="bg-white p-40rpx
            rounded-tr-3xl rounded-bl-3xl rounded-br-3xl">
              <sar-scroll-list
                scrollbar-bg="rgba(var(--sar-danger-rgb), 0.2)"
                thumb-bg="var(--sar-danger)"
              >
                  <view class="flex gap-3">
                    <view v-for="item in 8" class="w-260rpx rounded-xl bg-black">
                      <sar-image class="rounded-xl"
                        src="https://fastly.jsdelivr.net/npm/@sard/assets/images/cat1.jpg"
                        width="100%"
                        height="180rpx"
                      />
                      <view class="px-16rpx py-16rpx">
                        <text class="text-sm line-clamp-1 leading-none overflow-hidden
                      text-white ">保底300w双子星猛攻单</text>
                      </view>
                    </view>
                  </view>
              </sar-scroll-list>
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


const rows = ref(
  Array(16)
    .fill(0)
    .reduce(
      (rows, _, i) => {
        rows[i % 2].push({
          text: '文字' + (i + 1),
          icon: 'image',
        })
        return rows
      },
      [[], []],
    ),
)
const description = ref(
  'unibest 是一个集成了多种工具和技术的 uniapp 开发模板，由 uniapp + Vue3 + Ts + Vite5 + UnoCss + VSCode 构建，模板具有代码提示、自动格式化、统一配置、代码片段等功能，并内置了许多常用的基本组件和基本功能，让你编写 uniapp 拥有 best 体验。',
)
console.log('index/index 首页打印了')

onLoad(() => {
  console.log('测试 uni API 自动引入: onLoad')
})
</script>


<style lang="scss" scoped>
.text-shadow-c{
  text-shadow: #f03740 -1px -3px, #2addfd 3px 0px;
}
.stroke-join {
  position: relative;
  -webkit-text-stroke: 4rpx #663df3;
}
.stroke-join::after {
  content: attr(data-content);
  position: absolute;
  top: 0;
  left: 0;
  -webkit-text-stroke: 0;
}


.clip-r {
  clip-path: polygon(10% 0%, 100% 0%, 100% 100%, 0% 100%);
  
}
.clip-l {
  clip-path: polygon(0% 0%, 100% 0%, 90% 100%, 0% 100%);
}

.bg-g {
  // background:
  // linear-gradient(to bottom, transparent, #fdfdfd 400rpx),
  //   linear-gradient(to right,  #ffdee9c4 40%,#b5fffc8f);
  background: 
  linear-gradient(to right,  #663df3,#663df3);
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
  
