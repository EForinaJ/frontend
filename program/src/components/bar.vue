<template>
    <view class="nav-bar-container">
        <view class="bg-white
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
  name: 'Bar',
})

const goBack = () =>{
    uni.navigateBack()
}

</script>

<style lang="scss" scoped>

</style>