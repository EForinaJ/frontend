<template>
    <view 
       :style="{minHeight: getWindowHeight() + 'px'}"
       class="bg-#f5f5f5">
        <view class="bg-white w-full px-30rpx fixed box-border z-33">
            <view :style="{height:getStatusBarHeight()+'px'}" ></view>
            <view 
                class="flex items-center"
            :style="{height:getTitleBarHeight()+'px'}" >
                <view @click="goBack" class="h-full flex items-center">
                    <sar-icon name="left" size="40rpx" />
                </view>
                <view class="flex-1 flex items-center justify-center mr-40rpx">
                    <view class="font-700
                        text-text-primary text-[36rpx] font-bold
                        leading-none">
                        订单列表
                    </view>
                </view>
            </view>
            <view class="h-80rpx flex items-center">
                <view class="w-full">
                    <sar-tabs 
                        @change="handleTab"
                        v-model:current="query.status" 
                        :list="option" 
                        style="--sar-tabs-line-bg: transparent">
                        <template #line>
                            <sar-icon
                                family="demo-icons"
                                name="smile-line"
                                color="var(--sar-danger)"
                                size="40rpx"
                            />
                        </template>
                    </sar-tabs>
                </view>
            </view>
        </view>
        <view :style="{height:40+getStatusBarHeight()+getTitleBarHeight()+'px'}"></view>
      
        <sar-pull-down-refresh
            ref="pullDownRefresh"
            :loading="fetchListLoading"
            :root-style="`min-height: ${getWindowHeight()}px`"
            @refresh="onRefresh"
        >
            <view v-if="fetchListLoading" class="mt-20rpx mx-30rpx">
                <view v-for="item in 6" :key="item" class="p-20rpx rounded-lg bg-white mb-30rpx">
                    <sar-skeleton />
                </view>
            </view>
            <view v-if="!fetchListLoading" class="mt-20rpx mx-30rpx">
                <view v-if="list.length > 0" >
                        <view v-for="item in list" :key="item.id" class="mb-30rpx">
                            <sar-card>
                                <template #title>
                                    <view class="flex items-center gap-2">
                                        <text class="text-text-secondary text-xs font-700">{{ item.product.game }}</text>
                                        <sar-tag plain theme="primary" size="small">{{item.product.category}}</sar-tag>
                                    </view>
                                </template>
                                <template #extra>
                                    <text v-if="item.status == OrderStatus.PendingPayment" class="text-pending text-xs font-700">
                                        待支付
                                    </text>
                                    <text v-if="item.status == OrderStatus.PendingService" class="text-pending text-xs font-700">
                                        待服务
                                    </text>
                                    <text v-if="item.status == OrderStatus.InProgress" class="text-inProgress text-xs font-700">
                                        进行中
                                    </text>
                                    <text v-if="item.status == OrderStatus.Completed" class="text-completed text-xs font-700">
                                        已完成
                                    </text>
                                    <text v-if="item.status == OrderStatus.Cancel" class="text-text-secondary text-xs font-700">
                                        交易关闭
                                    </text>
                                    <text v-if="item.status == OrderStatus.Refund" class="text-refund text-xs font-700">
                                        已退款
                                    </text>
                                </template>
                                <!-- <template #footer>
                                    <view class="flex justify-end gap-2">
                                        <sar-button 
                                        @click="handleCancel(item.id)"
                                        v-if="item.status == OrderStatus.PendingPayment" size="mini" inline type="outline">取消订单</sar-button>
                                        <sar-button 
                                        @click="handleDelete(item.id)"
                                        v-if="item.status == OrderStatus.Cancel" 
                                        size="mini" inline type="outline" theme="danger">删除订单</sar-button>

                                        <sar-button 
                                        @click="handlePay(item.id,item.status)"
                                        v-if="item.status == OrderStatus.PendingPayment" 
                                        size="mini" inline type="default">去支付</sar-button>
                                        <sar-button v-if="item.status == OrderStatus.Completed" size="mini" inline type="default">去评价</sar-button>
                                        <sar-button v-if="item.status == OrderStatus.Completed" size="mini" inline type="default" theme="danger">申请售后</sar-button>
                                    </view>
                                </template> -->
                                <view @click="handleDetail(item.id)" class="flex gap-2">
                                    <sar-image
                                        :src="item.product.pic"
                                        width="108rpx"
                                        height="108rpx"
                                        radius="12rpx"
                                    ></sar-image>
                                    <view class="flex-1">
                                        <view class="break-all break-words text-[36rpx] text-text-primary font-700">
                                            {{ item.product.name }}
                                        </view>
                                    </view>
                                    <view class="flex flex-col items-center justify-center gap-1">
                                        <text class="text-text-primary text-xs font-700">{{item.unitPrice}}{{siteStore.siteInfo.symbol}}/{{ item.unit }}</text>
                                        <text class="text-text-primary text-xs font-700">*{{item.quantity}}</text>
                                    </view>
                                </view>
                                <view class="mt-20rpx flex justify-end text-xs">
                                    <text class="text-text-primary">订单总额</text>
                                    <text class="text-text-primary font-700">{{item.totalAmount}}{{ siteStore.siteInfo.symbol }}</text>
                                </view>
                            </sar-card>
                        </view>
                        <view class="flex items-center justify-center">
                            <sar-load-more
                                :status="loadMoreStatus"
                                @load-more="onLoadMore"
                                @reload="onReload"
                                />
                        </view>
                </view>
                <view v-if="list.length == 0"  class="p-20rpx rounded-lg bg-white mb-30rpx">
                    <sar-empty  icon-size="96rpx"/>
                </view>
            </view>
        </sar-pull-down-refresh>
    <sar-toast-agent />
    <sar-dialog-agent />
   </view>
</template>

<script setup lang="ts">
import { getOrderList, postOrderCancel, postOrderDelete } from "@/api/order";
import { OrderStatus } from "@/enum/status";

import useBoolean from "@/hooks/boolean";
import { useSiteStore } from "@/store/site";
import {  getStatusBarHeight, getTitleBarHeight, getWindowHeight, getWindowWidth, safeAreaInsets } from '@/utils/systemInfo';
import { dialog, LoadMoreStatus, TabOption, toast } from "sard-uniapp";
definePage({
   needLogin:true,
   style: {
       navigationBarTitleText: '订单列表',
       navigationStyle: 'custom'
   },
})

const current = ref(0)
const option = ref<TabOption[]>([
  { title: '全部', name: 0 },
  { title: '待付款', name: OrderStatus.PendingPayment },
  { title: '待服务', name: OrderStatus.PendingService  },
  { title: '进行中', name: OrderStatus.InProgress },
  { title: '已完成', name: OrderStatus.Completed },
])


const siteStore = useSiteStore()


const query = reactive<Order.Dto.Query>({
    page:1,
    limit:8,
    status:0,
})
const list = ref<Order.Dao.List[]>([])
const { bool:fetchListLoading, setTrue:fetchListSetTrue, setFalse:fetchListSetFalse } = useBoolean()
const loadMoreStatus = ref<LoadMoreStatus>('incomplete')
const getData = async () =>{
    loadMoreStatus.value = 'loading'
    try {
        const res = await getOrderList(query)
        list.value = [...list.value,...res.list]
        loadMoreStatus.value = query.limit < res.list.length ? 'incomplete' : 'complete'
    } catch (error) {
      toast.fail(error)  
    }
}

const onLoadMore = async () => {
  if (!fetchListLoading.value) {
    query.page ++
    await getData()
  }
}

const onReload = async () => {
    if (!fetchListLoading.value) {
        query.page  = 1
        await getData()
    }
}

const pullDownRefresh = ref()
onPageScroll(({ scrollTop }) => {
  pullDownRefresh.value?.enableToRefresh(scrollTop === 0)
})
const onRefresh = async () => {
    list.value = []
    fetchListSetTrue()
    await  getData()
    fetchListSetFalse()
}

const handleTab = async (e:number) => {
    list.value = []
    query.status = e
    fetchListSetTrue()
    await  getData()
    fetchListSetFalse()
}

const handleCancel = async (id:number) =>{

    dialog.confirm({
        message: '确定取消吗',
        onConfirm: async ()=>{
            try {
                await postOrderCancel({id})
                list.value = []
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
                list.value = []
                await getData()
            } catch (error) {
                toast.fail(error)  
            }
        }
    })
    
}

const  handlePay = async (id:number,status:number) => {
    if (status != OrderStatus.PendingPayment) {
        toast.fail("订单已支付")  
        return
    }
    uni.navigateTo({
        url: `/pages/pay/index?id=${id}`
    });
}


const handleDetail = (id:number) => {
    uni.navigateTo({
        url: `/pages-me/order/detail?id=${id}`
    });
}

const goBack = () =>{
    uni.navigateBack()
}
onLoad(async (options)=>{
    fetchListSetTrue()
    if (options.tab != undefined) {
        query.status = Number(options.tab)
    }
    await  getData()
    fetchListSetFalse()
})
</script>

<style lang="scss" scoped>
.bar-box {
    background:
    linear-gradient(to bottom, transparent, #f5f5f5 400rpx),
    linear-gradient(to right,  #CE9FFC,#CE9FFC);
}

</style>