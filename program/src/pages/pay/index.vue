<template>
    <view 
        :style="{minHeight: getWindowHeight() + 'px'}"
        class="bg-[#f5f4f9]">
        <Bar arrow
            title="收银台"
        />
        <view v-if="!fetchLoading && detail" class="mx-30rpx mt-60rpx flex items-center justify-center flex-col gap-5">
            <text class="text-text-secondary font-700">实付款</text>
            <text class="text-text-primary text-2xl font-700">{{detail.actualAmount}}{{ siteStore.siteInfo.symbol }}</text>
        </view>

        <view class="mx-30rpx mt-80rpx">
            <sar-list card>
                <sar-list-item @click="handlePayType(PayType.Wechat)">
                    <template #title>
                        <view class="flex items-center gap-2">
                            <i class="iconfont icon-weichat text-[#3eb744] text-[40rpx]"></i>
                            <text>微信支付</text>
                        </view>
                    </template>
                    <template #description>
                        亿万用户的选择，更快更安全
                    </template>
                    <template #value>
                        <sar-icon
                        v-if="payType == PayType.Wechat"
                        color="var(--sar-primary)"
                        size="32rpx"
                        name="success"
                        />
                    </template>
                </sar-list-item>
                <sar-list-item @click="handlePayType(PayType.Balance)">
                    <template #title>
                        <view class="flex items-center gap-2">
                            <i class="iconfont icon-yingshouyue text-[#ff8848] text-[40rpx]"></i>
                            <text>余额支付</text>
                        </view>
                    </template>
                    <template #description>
                        <text class="text-red">我的余额:{{accountStore.accountInfo.balance}}{{ siteStore.siteInfo.symbol }}</text>
                    </template>
                    <template #value>
                        <sar-icon
                        v-if="payType == PayType.Balance"
                        color="var(--sar-primary)"
                        size="32rpx"
                        name="success"
                        />
                    </template>
                </sar-list-item>
            </sar-list>
        </view>

        <view class="mx-30rpx mt-40rpx">
            <sar-button 
            :loading="fetchPayLoading"
            @click="throttledPayment"  round  type="default">确认支付</sar-button>
        </view>

        <sar-toast-agent />
        <sar-dialog-agent />
    </view>
</template>

<script setup lang="ts">
import { getOrderDetail, getOrderStatus, postOrderBalancePay, postOrderWechatMiniProgramPay } from '@/api/order';
import { getRechargesStatus } from '@/api/recharge';
import Bar from '@/components/bar.vue';
import { OrderStatus } from '@/enum/status';
import { PayType } from '@/enum/type';
import useBoolean from '@/hooks/boolean';
import { NOT_FOUND_PAGE } from '@/router/config';
import { useAccountStore } from '@/store';
import { useSiteStore } from '@/store/site';
import { getWindowHeight } from '@/utils/systemInfo';
import ThrottleUtils from '@/utils/throttle';
import { toast } from 'sard-uniapp';

definePage({
   needLogin:true,
   style: {
       navigationBarTitleText: '收银台',
       navigationStyle: 'custom'
   },
})

const accountStore = useAccountStore()
const siteStore = useSiteStore()

const payType = ref<number>(PayType.Wechat)

const handlePayType = (e:number) => {
    payType.value = e
}


const handleBalance = async () => {
    if (accountStore.accountInfo.balance < detail.value.actualAmount) {
        toast.fail("余额不足，请充值") 
        return
    }
    fetchPaySetTrue()
    uni.showLoading({
        title:"支付中"
    })
    try {
        await postOrderBalancePay({id:id.value})
        //  跳转支付成功页面
        uni.redirectTo({
            url:`/pages/pay/success?id=${id.value}`
        })
    } catch (error) {
      toast.fail(error)  
    }
    uni.hideLoading()
    fetchPaySetFalse()
}


const handleWxMiniProgramPay = async () => {
    fetchPaySetTrue()
    uni.showLoading({
        title:"支付中"
    })
    try {
        const res = await postOrderWechatMiniProgramPay({id:id.value})
        startPolling(res.id)
        uni.requestPayment({
            provider:"wxpay",
            orderInfo:"",
            ...res,
            fail:()=>{
                stopPolling();
                toast.fail("取消支付")
            }
        })
        //  跳转支付成功页面
    } catch (error) {
      toast.fail(error)  
    }
    uni.hideLoading()
    fetchPaySetFalse()
}



const { bool:fetchPayLoading, setTrue:fetchPaySetTrue, setFalse:fetchPaySetFalse } = useBoolean()
const handlePay = () => {
   if (payType.value == PayType.Balance) {
        handleBalance()
   }

   if (payType.value == PayType.Wechat) {
        handleWxMiniProgramPay()
   }
}

// 创建节流支付函数 - 1秒内只能点击一次
const throttledPayment = ThrottleUtils.create(handlePay, {
  wait: 1000,
  leading: true,    // 立即执行第一次点击
  trailing: false   // 不执行最后一次点击
});

const pollingInterval = ref<number | NodeJS.Timeout>(null);
const maxPollingCount = 60; // 最大轮询次数
let pollingCount = 0;

// 停止轮询
const stopPolling = () => {
    if (pollingInterval.value) {
        clearInterval(pollingInterval.value);
        pollingInterval.value = null;
    }
};
const startPolling = (id:number) => {
    pollingInterval.value = setInterval(async () => {
      try {
        const res = await getOrderStatus({id})
        
        // 如果订单状态为成功或失败，停止轮询
        if (res.status === OrderStatus.PendingService) {
            stopPolling();
            //  跳转支付成功页面
            uni.redirectTo({
                url:`/pages/pay/success?id=${id}`
            })
            return
        }

        // 如果达到最大轮询次数，停止轮询
        pollingCount++;
        if (pollingCount >= maxPollingCount) {
          stopPolling();
          // 可以在这里处理超时逻辑
          toast.fail("订单校验超时")
        }
      } catch (error) {
        toast.fail("订单校验异常")
        // 出错时可以选择停止轮询，或者继续轮询
        // 这里我们选择停止轮询
        stopPolling();
      }
    }, 2000); // 每2秒轮询一次
};


// 组件卸载时清除定时器
onUnmounted(() => {
    stopPolling();
});


const id = ref<number>(0)
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
        uni.navigateTo({ url: NOT_FOUND_PAGE })
    }
    id.value = Number(options.id)

    getData()
})
</script>