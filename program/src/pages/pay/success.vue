<template>
    <view 
        :style="{minHeight: getWindowHeight() + 'px'}"
        class="bg-[#f5f4f9]">
        <Bar arrow
            title="支付结果"
        />
        <view class="mt-40rpx bg-white rounded-lg mx-30rpx px-30rpx py-80rpx">
            <sar-result status="success" title="支付成功" description="请尽快联系客服进行服务派单">
               
            </sar-result>
        </view>

        <view v-if="!fetchLoading && detail" class="mt-40rpx bg-white rounded-lg mx-30rpx">
            <sar-list card>
                <sar-list-item arrow @click="weChatVisibleSetTrue">
                    <template #title>
                        <view class="flex items-center gap-2">
                            <i class="iconfont icon-weichat text-[#3eb744] text-[40rpx]"></i>
                            <text>微信客服号</text>
                        </view>
                    </template>
                    <template #description>
                        关注{{ siteStore.siteInfo.title }},随时掌上联系
                    </template>
                </sar-list-item>
                <sar-list-item arrow @click="copy">
                    <template #title>
                        <view class="flex items-center gap-2">
                            <sar-image
                                :src="siteStore.siteInfo.contact.icon"
                                width="40rpx"
                                height="40rpx"
                                radius="12rpx"
                            ></sar-image>
                            <text>{{siteStore.siteInfo.contact.platform}}值班客服</text>
                        </view>
                    </template>
                    <template #description>
                        在线时间:{{siteStore.siteInfo.contact.onlineTime}}
                    </template>
                </sar-list-item>
            </sar-list>
        </view>

        <view v-if="!fetchLoading && detail" class="mt-40rpx bg-white rounded-lg mx-30rpx">
            <sar-list card>
                <sar-list-item>
                    <template #title>
                        <text class="text-xs text-text-secondary">订单号</text>
                    </template>
                    <template #value>
                        <text class="text-xs text-text-primary">{{detail.code}}</text>
                    </template>
                </sar-list-item>
                <sar-list-item>
                    <template #title>
                        <text class="text-xs text-text-secondary">支付金额</text>
                    </template>
                    <template #value>
                        <text class="text-xs text-text-primary">{{detail.payAmount}}{{ siteStore.siteInfo.symbol }}</text>
                    </template>
                </sar-list-item>
                <sar-list-item>
                    <template #title>
                        <text class="text-xs text-text-secondary">支付时间</text>
                    </template>
                    <template #value>
                        <text class="text-xs text-text-primary">{{detail.payTime ? detail.payTime : "暂无"}}</text>
                    </template>
                </sar-list-item>
            </sar-list>
        </view>


        <sar-overlay :visible="weChatVisible" @click="weChatVisiblealse">
            <view class="flex items-center justify-center h-full">
                <view class="w-400rpx h-400rpx bg-white rounded-lg p-20rpx">
                    <sar-image
                        :src="siteStore.siteInfo.contact.wechat"
                        width="400rpx"
                        height="400rpx"
                    ></sar-image>
                </view>
            </view>
        </sar-overlay>
        <sar-toast-agent />
    </view>
</template>

<script setup lang="ts">
import { getOrderDetail } from '@/api/order';
import Bar from '@/components/bar.vue';
import useBoolean from '@/hooks/boolean';
import { NOT_FOUND_PAGE } from '@/router/config';
import { useSiteStore } from '@/store/site';
import { getWindowHeight } from '@/utils/systemInfo';
import { toast } from 'sard-uniapp';


definePage({
   needLogin:true,
   style: {
       navigationBarTitleText: '支付结果',
       navigationStyle: 'custom'
   },
})

const { bool:weChatVisible, setTrue:weChatVisibleSetTrue, setFalse:weChatVisiblealse } = useBoolean()


const copy = async () => {
    uni.setClipboardData({
    data: siteStore.siteInfo.contact.number, // 替换为你需要复制的实际文本
    success: function () {
        uni.showToast({
        title: `已复制平台号码`,
        icon: 'success'
        });
    },
    fail: function (err) {
        console.log('复制失败', err);
        // 处理复制失败的情况
    }
    });
}

const id = ref<number>(0)
const siteStore = useSiteStore()
const { bool:fetchLoading, setTrue:fetchSetTrue, setFalse:fetchSetFalse } = useBoolean()
const detail = ref<Order.Dao.Detail>()
const getData = async () => {
    fetchSetTrue()
    uni.showLoading({
        title:"加载中"
    })
    try {
        const res = await getOrderDetail({id:id.value})
        detail.value = res
    
        uni.hideLoading()
    } catch (error) {
      toast.fail(error)  
    }
    fetchSetFalse()
}

onLoad((options)=>{
    if (options.id == undefined) {
        uni.redirectTo({ url: NOT_FOUND_PAGE })
    }
    id.value = Number(options.id)
    getData()
})
</script>