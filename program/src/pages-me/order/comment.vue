<template>
     <view 
        :style="{minHeight: getWindowHeight() + 'px'}"
        class="bg-[#f5f4f9]">
        <Bar arrow
            title="订单评价"
        />
        <view class="mx-30rpx mt-40rpx ">
            <view class="font-700">
            商品评价
            </view>
            <view class="mt-20rpx ">
                <sar-card title="Trumpet to the Son of Heaven">
                    <template #title>
                        <text class="font-700">您对本次服务是否满意？</text>
                    </template>
                    <view class="flex items-center gap-2">
                        <sar-image
                            src="https://fastly.jsdelivr.net/npm/@sard/assets/images/cat1.jpg"
                            width="80rpx"
                            height="80rpx"
                            radius="12rpx"
                        ></sar-image>
                        <view>
                            商品收到贵司的代肝 
                        </view>
                    </view>
                    <view class="mt-20rpx">
                        <sar-rate v-model="value" @change="onChange" />
                    </view>
                    <view class="mt-20rpx">
                        <sar-input placeholder="请输入" type="textarea" />
                    </view>
                </sar-card>
            </view>
        </view>
        <view 
        class="fixed bottom-0 w-full">
            <view class="bg-white p-20rpx">
                <sar-button type="default">提交评价</sar-button>
            </view>
            <view :style="`height:${getSafeAreaHeight()}rpx`" class="bg-white"></view>
        </view>
    </view>
</template>

<script setup lang="ts">
import Bar from '@/components/bar.vue';
import { getSafeAreaHeight, getWindowHeight, safeAreaInsets } from '@/utils/systemInfo';


definePage({
   needLogin:true,
   style: {
       navigationBarTitleText: '订单评价',
       navigationStyle: 'custom'
   },
})

const value = ref(3)

const onChange = (value: any) => {
  console.log('change', value)
}
</script>