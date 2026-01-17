<template>
    <view 
       :style="{minHeight: getWindowHeight() + 'px'}"
       class="bg-[#f5f4f9]">
       <Bar arrow
        title="商品详情"
       />
       <view v-if="!fetchLoading && detail">
        <sar-image
            :src="detail.pic"
            width="100%"
            height="400rpx"
        />
            <view class="mx-30rpx mt-40rpx">
                <view class="bg-white p-20rpx rounded-lg">
                    <view class="flex items-center gap-1">
                        <view class="text-xs font-bold text-coolGray">
                            服务编码:{{ detail.code }}
                        </view>
                        <view @click="handleCopy(detail.code)" class="i-carbon-copy text-xs text-coolGray" />
                    </view>
                    <view class="text-[36rpx] font-bold mt-16rpx">
                        {{ detail.name }}
                    </view>
                    <view class="bg-tertiary/10 p-20rpx rounded-lg mt-16rpx">
                        <view class="text-xs text-primary ">
                            {{ detail.description }}
                        </view>
                    </view>
                    <view class="flex justify-between items-center mt-16rpx">
                        <view class="flex gap-2">
                            <sar-tag size="small" theme="primary">{{ detail.game }}</sar-tag>
                            <sar-tag size="small" theme="success">{{ detail.category }}</sar-tag>
                            <view class="text-xs text-secondary/60">
                                销量{{ detail.salesCount }}
                            </view>
                        </view>
                        <view class="text-red font-bold">
                            <text class="text-base">{{ detail.price }}</text>
                            <text class="text-xs">{{ site.symbol }}</text>
                            <text class="text-base">/</text>
                            <text class="text-base">{{ detail.unit }}</text>
                        </view>
                    </view>
                </view>
            </view>

            <view class="mx-30rpx mt-40rpx">
                <view class="bg-white p-20rpx rounded-lg">
                    <view class="flex items-center justify-center">服务说明</view>
                    <view class="mt-20rpx">
                        <mp-html :content="detail.content" />
                    </view>
                </view>
            </view>
        </view>

        <view class="fixed bottom-0 w-full bg-white">
            <view class="mx-30rpx h-100rpx flex gap-3">
                <view class="w-100rpx flex items-center justify-center">
                    <sar-button @click="handleShare"  type="text" round size="medium" >
                        <view class="i-solar-box-minimalistic-bold text-xl"/>
                    </sar-button>
                </view>
                <view class="w-100rpx flex items-center justify-center">
                    <sar-button open-type="share" type="text" round size="medium" >
                        <view class="i-solar-square-share-line-bold-duotone text-xl"/>
                    </sar-button>
                </view>
                <view class="flex-1 flex items-center justify-center">
                    <sar-button theme="danger">咨询服务</sar-button>
                </view>
            </view>
            <view style="height: env(safe-area-inset-bottom); flex: none"></view>
        </view>
        
       <sar-toast-agent />
    </view>
</template>

<script setup lang="ts">
import { getProductDetail } from '@/api/product';
import Bar from '@/components/bar.vue';
import useBoolean from '@/hooks/boolean';
import { NOT_FOUND_PAGE } from '@/router/config';
import { useSiteStore } from '@/store/site';
import { getImageUrl } from '@/utils/imageHelper';
import {  getWindowHeight } from '@/utils/systemInfo';
import {  PopupProps, toast } from 'sard-uniapp';

definePage({
   style: {
       navigationBarTitleText: '商品详情',
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

const handleCopy = (e:string) =>{
    uni.setClipboardData({
    data: e,
    success: () => {
        uni.showToast({
            title: '复制成功',
            icon: 'success'
        });
    },
    fail: (err) => {
        console.log('复制失败', err);
    }
    });
}


const handleShare = async (e: PopupProps['effect']) => {
    uni.navigateTo({ url: `/pages/share/index?id=${id.value}` })
}



onLoad((options)=>{
    if (options.id == undefined) {
        uni.navigateTo({ url: NOT_FOUND_PAGE })
    }
    id.value = Number(options.id)
    getData()
})

onShareAppMessage((option)=>{
    return{
      title: detail.value.name,
      path: `/pages/product/detail?id=${id.value}`
    }
  }
)

// #ifdef MP-WEIXIN
onShareTimeline(()=>{
  return {
    title: detail.value.name,
    path: `/pages/product/detail?id=${id.value}`
  }
})
// #endif
</script>