<template>
  <view class="bg-#f5f5f5"
  :style="{minHeight: getWindowHeight()-(getTitleBarHeight()+getStatusBarHeight())+ 'px'}">
      <view  class="
      bg-primary w-full px-30rpx fixed box-border z-33">
        <view :style="{height:getStatusBarHeight()+'px'}" ></view>
        <view :style="{height:getTitleBarHeight()+'px'}" >
          <!-- <image class="w-[180rpx] h-full" src="~@/static/images/xl.png"/> -->
          <view class="text-white text-3xl h-full
            tracking-wider 
            font-bold  flex items-center
          ">
            {{site.title}}
          </view>
        </view>
      </view>
      <view :style="{height:40+getStatusBarHeight()+getTitleBarHeight()+'px'}"></view>

      <view class="mx-30rpx">
        <view class="pos-relative h-400rpx bg-primary rounded-lg mt-20rpx">
          <view class="pos-relative" :style="{height:getTitleBarHeight()+'px'}" >
            <view class="w-full flex justify-center items-center gap-3 pos-absolute top-[-50rpx]">
              <view class=" flex items-center justify-center text-shadow
              h-90rpx w-90rpx bg-primary rounded-lg border-[#fff] border-6rpx border-solid
              text-3xl text-[#ffff]">欢</view>

              <view class=" flex items-center justify-center text-shadow
              h-90rpx w-90rpx bg-primary rounded-lg border-[#fff] border-6rpx border-solid
              text-3xl text-[#e7e6f9]">迎</view>

              <view class=" flex items-center justify-center text-shadow
              h-90rpx w-90rpx bg-primary rounded-lg border-[#fff] border-6rpx border-solid
              text-3xl text-[#e7e6f9]">光</view>

              <view class=" flex items-center justify-center text-shadow
              h-90rpx w-90rpx bg-primary rounded-lg border-[#fff] border-6rpx border-solid
              text-3xl text-[#e7e6f9]">临</view>
            </view>
          </view>
  
          <view class="ribbon-2">
            <view class="flex-1 flex items-center justify-center  gap-2
              bg-secondary rounded-tr-xl rounded-br-xl">
                <image class="size-48rpx" :src="site.contact.icon"/>
                <view class="text-white text-[36rpx] font-bold">{{site.contact.platform}}频道: {{site.contact.number}}</view>
              </view>
          </view>

          <view class="flex gap-1 px-40rpx pb-40rpx">
            <view @click="handelPush(`/pages/join/index`)"
            class="rounded-tl-xl rounded-bl-xl pos-relative
            flex-1 bg-secondary  clip-l p-20rpx">
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
            <button open-type="contact"
            class="flex-1 w-full text-white text-sm font-semibold
              flex items-center justify-center
              bg-secondary rounded-tr-xl rounded-br-xl">
              在线客服
            </button>
              <view 
               @click="handelPush(`/pages-protocol/instructions/index`)"
              class="flex-1 text-white text-sm font-semibold
              flex items-center justify-center
              bg-secondary rounded-tr-xl rounded-br-xl">
                预约须知
              </view>
            </view>
          </view>
        </view>
      </view>

      <view v-if="!fetchLoading && detail" class="mx-30rpx mt-40rpx">
        <view class="shadow-2xl 
        bg-secondary rounded-3xl mt-20rpx">
          <view class="flex justify-between items-center px-30rpx py-16rpx">
            <view class="flex items-center gap-2">
              <view class="flex items-center justify-center">
                <image class="size-40rpx" src="~@/static/images/game.png"/>
              </view>
              <view class="font-bold text-[36rpx] text-[#ffff]">
                游戏领域
              </view>
            </view>
          </view>
          <view class="bg-primary rounded-3xl p-40rpx">
            <sar-scroll-list
              scrollbar-bg="rgba(var(--sar-danger-rgb), 0.2)"
              thumb-bg="var(--sar-danger)"
            >
                <view class="flex gap-3">
                  <view 
                  v-for="item in detail.gameList" :key="item.id"
                  @click="handelPush(`/pages/product/list?gameId=${item.id}`)"
                  class=" w-160rpx
                  px-16rpx bg-tertiary py-20rpx rounded-lg mb-30rpx
                  flex-col flex items-center justify-center">
                    <image class="size-60rpx" :src="item.pic"/>
                    <view class="text-xs mt-16rpx text-white">{{item.name}}</view>
                  </view>
                </view>
            </sar-scroll-list>
          </view>
        </view>
      </view>

      <view v-if="!fetchLoading && detail" class="mx-30rpx mt-40rpx">
        <view class="text-xl font-bold pos-relative
        w-fit 
        after:content-[''] after:w-full 
        after:bg-gradient-to-r after:from-[#693FF8] after:to-transparent
        after:inline-block
        after:rounded-lg after:h-10rpx 
        after:absolute after:bottom-4rpx after:left-0
        ">
          热门预约
        </view>
        <view 
        v-for="(item) in detail.productList" :key="item.game.id"
        class="bg-primary rounded-lg p-20rpx mt-20rpx">
          <view class="rounded-lg">
            <view class="bg-white flex
             rounded-tl-3xl h-80rpx
            ">
              <view class="bg-primary">
                <view class="w-fit pl-30rpx pr-70rpx flex items-center justify-center  gap-2
                bg-white rounded-tl-3xl rounded-tr-3xl h-full">
                  <view class="rounded-full px-8rpx flex items-center justify-center">
                    <image class="size-40rpx" :src="item.game.pic"/>
                  </view>
                  <view class="font-bold text-[36rpx]">
                    {{ item.game.name }}
                  </view>
                </view>
              </view>

              <view class="flex-1
              bg-primary h-full flex justify-end items-center
              rounded-bl-3xl">
                <view  @click="handelPush(`/pages/product/list?gameId=${item.game.id}`)"
                class="flex gap-2 items-center bg-white py-6rpx px-16rpx
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
                    <view @click="handelPush(`/pages/product/detail?id=${vitem.id}`)"
                    v-for="vitem in item.products" :key="vitem.id" class="w-260rpx rounded-xl bg-black">
                      <sar-image class="rounded-xl"
                        :src="vitem.pic"
                        width="100%"
                        height="180rpx"
                      />
                      <view class="px-16rpx py-16rpx">
                        <text class="text-sm line-clamp-1 leading-none overflow-hidden
                      text-white ">{{ vitem.name }}</text>
                      </view>
                    </view>
                  </view>
              </sar-scroll-list>
            </view>
          </view>
        </view>
      </view>
      
      <view :style="{height:getTitleBarHeight()+'px'}"></view>
  </view>
</template>

<script lang="ts" setup>
import { getSiteHome } from '@/api/site'
import useBoolean from '@/hooks/boolean'
import { useSiteStore } from '@/store/site'
import { getStatusBarHeight, getTitleBarHeight, getWindowHeight } from '@/utils/systemInfo'
import { toast } from 'sard-uniapp'

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

const { siteInfo:site } = useSiteStore()

const handelPush = (url:string) => {
  uni.navigateTo({url: url})
}


const { bool:fetchLoading, setTrue:fetchSetTrue, setFalse:fetchSetFalse } = useBoolean()
const detail = ref<Site.Dao.Home>()
const getData = async () => {
    fetchSetTrue()
    uni.showLoading({
        title:"加载中"
    })
    try {
        const res = await getSiteHome()
        detail.value = res
        uni.hideLoading()
    } catch (error) {
      toast.fail(error)  
    }
    fetchSetFalse()
}

onShareAppMessage((option)=>{
    return{
      title:site.title,
      path: '/pages/index/index'
    }
  }
)

// #ifdef MP-WEIXIN
onShareTimeline(()=>{
  return {
    title: '分享到朋友圈',
    path: '/pages/index/index'
  }
})
// #endif
onLoad(() => {
  getData()
})
</script>


<style lang="scss" scoped>
.ribbon-2 {
  position: absolute;
  left: -20rpx;
  bottom: 40rpx;
  width: calc(100% + 40rpx);
  height: 100rpx;
  background-color: #3D2F51;

  /* 内容居中 */
  display: flex;
  justify-content: center;
  align-items: center;

  &::before, &::after {
    content: '';
    position: absolute;
  }

  &::before {
    left: 0;
    height: 0;
    width: 0;
    border-top: 20rpx solid #3D2F51;
    border-left: 20rpx solid transparent;
    bottom: -20rpx;
  }

  &::after {
    right: 0;
    bottom: -20rpx;
    height: 0;
    width: 0;
    border-top: 20rpx solid #3D2F51;
    border-right: 20rpx solid transparent;
  }
}
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


</style>
  
