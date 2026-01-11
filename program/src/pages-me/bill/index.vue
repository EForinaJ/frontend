<template>
    <view 
       :style="{minHeight: getWindowHeight() + 'px'}"
       class="page-container">
       <NavBar  arrow
       :search="false" title="我的账单"/>
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
                            <view class="text-text-secondary text-xs flex items-center justify-center mb-30rpx">
                                {{ item.createTime }}
                            </view>
                            <sar-card>
                                <template #title>
                                    <text class="font-700">{{ filterType(item.type) }}</text>
                                </template>
                                <template #footer>
                                    <text class="text-xs text-text-secondary font-700">{{filterContent(item.type,item.money)}}</text>
                                </template>
                                <view class="flex items-center justify-center flex-col">
                                    <text class="text-3xl font-700 text-text-primary">{{ item.money }}</text>
                                    <text class="text-xl font-700 text-text-primary">{{ siteStore.siteInfo.symbol }}</text>
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
   </view>
</template>

<script setup lang="ts">

import { getBillList } from "@/api/bill";
import NavBar from "@/components/nav-bar.vue";
import { BillType } from "@/enum/type";
import useBoolean from "@/hooks/boolean";
import { useSiteStore } from "@/store/site";
import {  getWindowHeight } from '@/utils/systemInfo';
import { LoadMoreStatus, toast } from "sard-uniapp";
definePage({
   needLogin:true,
   style: {
       navigationBarTitleText: '我的账单',
       navigationStyle: 'custom'
   },
})

const siteStore = useSiteStore()


const query = reactive<Bill.Dto.Query>({
    page:1,
    limit:8,
    type:0,
})
const list = ref<Bill.Dao.List[]>([])
const { bool:fetchListLoading, setTrue:fetchListSetTrue, setFalse:fetchListSetFalse } = useBoolean()
const loadMoreStatus = ref<LoadMoreStatus>('incomplete')
const getData = async () =>{
    loadMoreStatus.value = 'loading'
    try {
        const res = await getBillList(query)
        list.value = [...list.value,...res.list]
        loadMoreStatus.value = query.limit < res.list.length ? 'incomplete' : 'complete'
    } catch (error) {
      toast.fail(error)  
    }
}

const filterType = (e:number) =>{
    const label: Record<number, string> = {
        [BillType.Recharge]: '余额充值',
        [BillType.Refund]: '订单退款',
        [BillType.Order]: '订单支付',
    };

    return label[e]
}
const filterContent = (e:number,money:number) =>{
    const label: Record<number, string> = {
        [BillType.Recharge]: "充值订单，" + "用户充值" + money +  siteStore.siteInfo.symbol,
        [BillType.Refund]: "订单退款，" + "平台返回订单" + money +  siteStore.siteInfo.symbol,
        [BillType.Order]: "订单支付，" + "购买服务支付了" +  money +  siteStore.siteInfo.symbol,
    };

    return label[e]
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

onLoad(async ()=>{
    fetchListSetTrue()
    await  getData()
    fetchListSetFalse()
})
</script>