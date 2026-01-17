<template>
    <view 
       :style="{minHeight: getWindowHeight() + 'px'}"
       class="bg-[#f5f4f9]">
       <Bar arrow
        title="加入我们"
       />
       <view>
            <view class="mx-30rpx mt-40rpx">
                <view class="background-x
                rounded-lg p-20rpx">
                    <view class="bg-white p-20rpx rounded-lg">
                        <view class="text-[36rpx] font-bold text-secondary">
                           入驻要求
                        </view>
                        <view class="flex gap-2
                        mt-20rpx bg-tertiary p-20rpx rounded-lg">
                            <view class="size-100rpx">
                                <image class="size-100rpx" src="~@/static/images/vel.png"/>
                            </view>
                            <view class="flex-1">
                                <view class="text-white text-sm font-bold">
                                    考核认证
                                </view>
                                <view class="text-white text-xs font-bold">
                                    完成并且通过{{ site.title }}考核内容以及认证流程，然后入驻
                                </view>
                            </view>
                        </view>
                        <view class="flex gap-2
                        mt-20rpx bg-tertiary p-20rpx rounded-lg">
                            <view class="flex-1">
                                <view class="text-white text-sm font-bold">
                                    资深游玩
                                </view>
                                <view class="text-white text-xs font-bold">
                                    有长时间的游戏经验，对游戏有热情，拥有不错的游戏理解
                                </view>
                            </view>
                            <view class="size-100rpx">
                                <image class="size-100rpx" src="~@/static/images/x-game.png"/>
                            </view>
                        </view>
                        <view class="flex gap-2
                        mt-20rpx bg-tertiary p-20rpx rounded-lg">
                            <view class="size-100rpx">
                                <image class="size-100rpx" src="~@/static/images/taidu.png"/>
                            </view>
                            <view class="flex-1">
                                <view class="text-white text-sm font-bold">
                                    服务态度
                                </view>
                                <view class="text-white text-xs font-bold">
                                    个人责任心，服务精神，履约精神达标，有良好的表达能力，遵守平台约定
                                </view>
                            </view>
                        </view>
                    </view>

                    <view class=" flex-col
                    flex items-center justify-center
                    my-40rpx">
                        <view class="p-10rpx rounded-lg bg-white">
                            <sar-image
                                :src="site.contact.wechat"
                                width="400rpx"
                                height="400rpx"
                                class="rounded-lg"
                            />
                        </view>
                        <view class="mt-30rpx text-xl font-bold">
                            添加客服备注来意
                        </view>
                    </view>
                </view>
            </view>
       </view>
    
    </view>
</template>

<script setup lang="ts">
import { getProductDetail } from '@/api/product';
import Bar from '@/components/bar.vue';
import useBoolean from '@/hooks/boolean';
import { useSiteStore } from '@/store/site';
import {  getWindowHeight } from '@/utils/systemInfo';
import {  toast } from 'sard-uniapp';

definePage({
   style: {
       navigationBarTitleText: '加入我们',
       navigationStyle: 'custom'
   },
})
const id = ref<number>(0)
const {siteInfo:site} = useSiteStore()
const { bool:fetchLoading, setTrue:fetchSetTrue, setFalse:fetchSetFalse } = useBoolean()
const detail = ref<Product.Dao.Detail>()
const getData = async () => {
    fetchSetTrue()
    uni.showLoading({
        title:"加载中"
    })
    try {
        const res = await getProductDetail({id:id.value})
        detail.value = res
        uni.hideLoading()
    } catch (error) {
      toast.fail(error)  
    }
    fetchSetFalse()
}


onLoad(()=>{
   
})
</script>

<style lang="scss" scoped>
.background-x {
    background: linear-gradient(
       to right,
       #bb313e25,
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
       #8a244b25
  );
}
</style>