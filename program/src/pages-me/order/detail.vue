<template>
    <view 
       :style="{minHeight: getWindowHeight() + 'px'}"
       class="bg-[#f5f4f9]">
       <Bar arrow
        title="订单详情"
       />
       <view v-if="!fetchLoading && detail">
            <view class="mx-30rpx mt-40rpx rounded-lg">
                <sar-card>
                    <template #title>
                        <view class="flex items-center gap-2">
                            <text class="text-text-secondary text-xs font-700">{{detail.product.game ?? "游戏分类"}}</text>
                            <sar-tag plain theme="primary" size="small">{{detail.product.type ?? "游戏类型"}}</sar-tag>
                        </view>
                    </template>
                    <template #extra>
                        <text v-if="detail.status == OrderStatus.PendingPayment"  class="text-pending text-xs font-700">
                            待支付
                        </text>
                        <text v-if="detail.status == OrderStatus.PendingService" class="text-pending text-xs font-700">
                            待服务
                        </text>
                        <text v-if="detail.status == OrderStatus.InProgress" class="text-inProgress text-xs font-700">
                            进行中
                        </text>
                        <text v-if="detail.status == OrderStatus.Completed" class="text-completed text-xs font-700">
                            已完成
                        </text>
                        <text v-if="detail.status == OrderStatus.Cancel" class="text-text-secondary text-xs font-700">
                            交易关闭
                        </text>
                        <text v-if="detail.status == OrderStatus.Refund" class="text-refund text-xs font-700">
                            已退款
                        </text>
                    </template>
                    <template #footer>
                        <view class="flex justify-end gap-2">
                            <sar-button 
                            @click="handleCancel(detail.id)"
                            v-if="detail.status == OrderStatus.PendingPayment" size="mini" inline type="outline">取消订单</sar-button>
                            <sar-button 
                                @click="handleDelete(detail.id)"
                            v-if="detail.status == OrderStatus.Cancel" 
                            size="mini" inline type="outline" theme="danger">删除订单</sar-button>

                            <sar-button 
                            @click="handlePay(detail.id,detail.status)"
                            v-if="detail.status == OrderStatus.PendingPayment" 
                            size="mini" inline type="default">去支付</sar-button>
                            <sar-button 
                            @click="handleComment(detail.id)"
                            v-if="detail.status == OrderStatus.Completed" 
                            size="mini" inline type="default">去评价</sar-button>
                        </view>
                    </template>
                    <view class="flex gap-2">
                        <sar-image
                            :src="detail.product.pic"
                            width="108rpx"
                            height="108rpx"
                            radius="12rpx"
                        ></sar-image>
                        <view class="flex-1">
                            <view class="break-all break-words text-text-primary font-700">
                                {{detail.product.name}}
                            </view>
                            <view class="break-all break-words text-text-secondary text-xs mt-10rpx">
                                {{detail.specifications}}
                            </view>
                        </view>
                        <view class="flex flex-col items-center justify-center gap-1">
                            <text class="text-text-primary text-xs font-700">{{detail.unitPrice}}{{siteStore.siteInfo.symbol}}</text>
                            <text class="text-text-primary text-xs font-700">*{{detail.quantity}}</text>
                        </view>
                    </view>
                    <view class="flex items-center justify-end gap-2">
                        <view class="mt-20rpx flex justify-end text-xs">
                            <text class="text-text-primary">总额</text>
                            <text class="text-text-primary font-700">{{detail.totalAmount}}{{siteStore.siteInfo.symbol }}</text>
                        </view>
                        <text class="text-text-primary font-700">-</text>
                        <view class="mt-20rpx flex justify-end text-xs">
                            <text class="text-text-primary">优惠</text>
                            <text class="text-text-primary font-700">{{detail.discountAmount}}{{siteStore.siteInfo.symbol }}</text>
                        </view>
                    </view>
                    <view class="mt-20rpx flex justify-end text-xs">
                        <text class="text-text-primary">实付款</text>
                        <text class="text-text-primary font-700">{{detail.actualAmount}}{{siteStore.siteInfo.symbol }}</text>
                    </view>
                </sar-card>
            </view>

            <view class="mt-40rpx bg-white rounded-lg mx-30rpx">
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

            <view class="mx-30rpx mt-40rpx rounded-lg">
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
                                <text class="text-xs text-text-secondary">下单时间</text>
                            </template>
                            <template #value>
                                <text class="text-xs text-text-primary">{{detail.createTime ? detail.createTime : "暂无"}}</text>
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
                        <sar-list-item>
                            <template #title>
                                <text class="text-xs text-text-secondary">开始时间</text>
                            </template>
                            <template #value>
                                <text class="text-xs text-text-primary">{{detail.startTime ? detail.startTime : "暂无"}}</text>
                            </template>
                        </sar-list-item>
                        <sar-list-item>
                            <template #title>
                                <text class="text-xs text-text-secondary">完成时间</text>
                            </template>
                            <template #value>
                                <text class="text-xs text-text-primary">{{detail.finishTime ? detail.finishTime : "暂无"}}</text>
                            </template>
                        </sar-list-item>
                    </sar-list>
            </view>
            

            <view class="mx-30rpx pb-40rpx mt-40rpx rounded-lg">
                    <sar-card>
                        <template #title>
                            <text class="text-text-secondary text-xs">订单备注</text>
                        </template>

                        <view class="break-all break-words text-text-secondary font-700 text-xs">
                            {{detail.requirements ? detail.requirements : "无备注"}}
                        </view>
                    </sar-card>
            </view>
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
       <sar-dialog-agent />
    </view>
</template>

<script setup lang="ts">
import { getOrderDetail, postOrderCancel, postOrderDelete } from '@/api/order';
import Bar from '@/components/bar.vue';
import { OrderStatus } from '@/enum/status';
import useBoolean from '@/hooks/boolean';
import { NOT_FOUND_PAGE } from '@/router/config';
import { useSiteStore } from '@/store/site';
import {  getWindowHeight } from '@/utils/systemInfo';
import { dialog, toast } from 'sard-uniapp';

definePage({
   needLogin:true,
   style: {
       navigationBarTitleText: '订单详情',
       navigationStyle: 'custom'
   },
})
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

const handleComment = (id:number) =>{
    uni.navigateTo({
        url: `/pages-me/order/comment?id=${id}`
    });
}

const handleCancel = async (id:number) =>{

    dialog.confirm({
        message: '确定取消吗',
        onConfirm: async ()=>{
            try {
                await postOrderCancel({id})
                await getData()
            } catch (error) {
                toast.fail(error)  
            }
        }
    })
   
}


const handleDelete = async (id:number) =>{

    dialog.confirm({
        title: '提示',
        message: '确定删除？',
        onConfirm: async ()=>{
            try {
                await postOrderDelete({id})
                uni.navigateBack()
            } catch (error) {
                toast.fail(error)  
            }
        }
    })
    
}



const  handlePay = async (orderId:number,status:number) => {
    if (status != OrderStatus.PendingPayment) {
        toast.fail("订单已支付")  
        return
    }
    uni.navigateTo({
        url: `/pages/pay/index?id=${orderId}`
    });
}

onLoad((options)=>{
    if (options.id == undefined) {
        uni.navigateTo({ url: NOT_FOUND_PAGE })
    }
    id.value = Number(options.id)
    
    getData()
})
</script>