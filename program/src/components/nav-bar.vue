<template>
    <view class="nav-bar-container">
        <view class="bar-box 
            fixed left-0 top-0 w-full z-36
        ">
            <view 
            :style="{height:getStatusBarHeight()+'px'}"
            ></view>
            <view 
            :style="{height:getTitleBarHeight()+'px'}"
            class="px-30rpx  flex items-center
            ">
                <view v-if="arrow" @click="goBack" class="h-full flex items-center">
                    <sar-icon name="left" size="40rpx" />
                </view>
                <view class="flex-1 flex items-center justify-center">
                    <view class=" font-700
                        text-text-primary
                        leading-none
                    ">{{ title }}</view>
                </view>
                <view @click="pushSearch" v-if="search"  class="search 
                    w-220rpx h-50rpx rounded-[60rpx] bg-white/50
                    border-2rpx border-solid border-white
                    ml-30rpx text-[#999] text-28rpx
                    flex items-center px-10rpx
                ">
                    <text class="ml-10rpx">搜索服务</text>
                </view>
            </view>
        </view>
        <view class="fill" :style="{height:getStatusBarHeight()+getTitleBarHeight()+'px'}"></view>
    </view>
</template>

<script setup lang="ts">
import { getStatusBarHeight, getTitleBarHeight } from '@/utils/systemInfo';
interface Props {
    /** the edit row data */
    arrow?:boolean | null;
    /** the edit row data */
    title?:string | null;
    /** the edit row data */
    search?:boolean | null;
}

const props = defineProps<Props>();

defineOptions({
  name: 'NavBar',
})

const goBack = () =>{
    uni.navigateBack()
}

const pushSearch = ()=>{
    uni.navigateTo({
        url: "/pages/search/index"
    })
}
</script>

<style lang="scss" scoped>
.bar-box {
    background: 
    linear-gradient(to bottom,transparent,#f5f5f5 400rpx),
    linear-gradient(to right, var(--primary-color), var(--second-color));
}

</style>