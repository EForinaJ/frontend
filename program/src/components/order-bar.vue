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
            class="pl-30rpx flex items-center
            ">
                <view @click="goBack" class="h-full flex items-center">
                    <sar-icon name="left" size="40rpx" />
                </view>
                <view class="flex-1 flex items-center justify-center">
                    <view class=" font-700
                        text-text-primary
                        leading-none
                    ">订单列表</view>
                </view>
            </view>
            <view class="mx-30rpx h-80rpx flex items-center">
                <view class="w-full">
                    <sar-tabs 
                    @change="handleTab"
                    v-model:current="current" :list="option" 
                    scrollable
                    type="pill"
                    ></sar-tabs>
                </view>
            </view>
        </view>
       
        <view :style="{height:40+getStatusBarHeight()+getTitleBarHeight()+'px'}"></view>
    </view>
</template>

<script setup lang="ts">
import { OrderStatus } from '@/enum/status';
import { getStatusBarHeight, getTitleBarHeight } from '@/utils/systemInfo';
import { TabOption } from 'sard-uniapp';
interface Props {
    /** the edit row data */
    arrow?:boolean | null;
  /** the edit row data */
  title?:string | null;
  /** the edit row data */
  search?:boolean | null;
}

const emit = defineEmits(['change'])

defineOptions({
  name: 'OrderBar',
})

const current = ref(0)
const option = ref<TabOption[]>([
  { title: '全部', name: 0 },
  { title: '待付款', name: OrderStatus.PendingPayment },
  { title: '待服务', name: OrderStatus.PendingService  },
  { title: '进行中', name: OrderStatus.InProgress },
  { title: '已完成', name: OrderStatus.Completed },
])



const handleTab = (name: string) => {
    emit('change', name)
}

const goBack = () =>{
    uni.navigateBack()
}

</script>

<style lang="scss" scoped>
.bar-box {
    background: 
  linear-gradient(to bottom,transparent,#e7eefa 400rpx),
  linear-gradient(to right,#FFF1EB,#ACE0F9);
}

</style>