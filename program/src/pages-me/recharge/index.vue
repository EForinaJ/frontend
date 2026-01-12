<template>
     <view 
     
        :style="{minHeight: getWindowHeight() + 'px'}"
        class="page-container">
        <NavBar 
        arrow
        :search="false" title="充值"/>

        <view class="mt-60rpx mx-30rpx bg-white p-30rpx rounded-lg">
            <view class="text-text-primary text-sm font-bold">充值{{ siteStore.siteInfo.symbol }}</view>
            <view class="mt-30rpx">
                <view class="bg-[#f5f4f9]  rounded-lg ">
                    <sar-input 
                        root-style="
                        font-weight: 700;
                        "
                        placeholder-style="
                        color: rgba(0,0,0,.3);
                        "
                        type="digit" 
                        v-model="inputValue" 
                        @input="handleInput"
                        inputmode="numeric"
                        borderless
                        placeholder="请输入充值金额">
                    </sar-input>
                </view>
            </view>
        </view>

        <view class="mt-60rpx mx-30rpx bg-white pt-30rpx px-30rpx rounded-lg">
            <sar-row :gap="30">
                <sar-col v-for="item in rechargeMoney" :key="item" :span="4">
                    <view 
                        @click="handleMoney(item)"
                        :class="model.money == item ? 'border-primary' : 'border-[#e7eefa]'"
                    class="mb-[30rpx] border-[4rpx] border-solid rounded-lg py-20rpx
                    flex flex-col items-center justify-center gap-2
                    ">
                        <view class="text-text-primary text-xl font-700">{{ item }}</view>
                        <view class="text-text-secondary text-xs">{{ siteStore.siteInfo.symbol }}</view>
                    </view>
                </sar-col>
            </sar-row>
        </view>

        <view class="mt-60rpx mx-30rpx bg-white p-30rpx rounded-lg">
            <view class="text-text-primary text-sm font-bold">支付方式</view>
            <view class="mt-30rpx">
                <sar-radio-group  root-style="margin: -30rpx">
                    <template #custom>
                    <sar-list card>
                        <sar-list-item
                        >
                            <template #icon>
                                <view class="w-full h-full flex items-center justify-center">
                                    <image 
                                        class="w-[48rpx] h-[48rpx]"
                                        src="/static/images/wechatPay.png"
                                    />
                                </view>
                            </template>
                            <template #title>
                                <view class="text-text-primary font-700">
                                    微信支付
                                </view>
                            </template>
                            <template #value>
                                <sar-radio readonly />
                            </template>
                        </sar-list-item>
                    </sar-list>
                    </template>
                </sar-radio-group>
            </view>
        </view>

        <view class="mt-60rpx mx-30rpx">
            <sar-button 
            :loading="fetchRechargeBoolean"
            @click="handleRecharge" round  type="default">
                {{ fetchRechargeBoolean ? "正在创建充值订单" : "立即充值" }}
            </sar-button>
        </view>
        <sar-toast-agent />
    </view>
</template>

<script setup lang="ts">
import { toast } from 'sard-uniapp'

import {  getRechargesStatus, postWeChatMiniProgramRecharge } from "@/api/recharge";
// import { getSiteDictData } from "@/api/site";


import NavBar from "@/components/nav-bar.vue";
import { useSiteStore } from "@/store/site";
import {  getWindowHeight } from '@/utils/systemInfo';
import useBoolean from '@/hooks/boolean';
import ThrottleUtils from '@/utils/throttle';
import { PayStatus } from '@/enum/status';
import { useAccountStore } from '@/store';

definePage({
    needLogin:true,
    style: {
        navigationBarTitleText: '充值',
        navigationStyle: 'custom'
    },
})

const siteStore = useSiteStore()


type Model = Recharge.Dto.Create;
const model = ref(createDefaultModel());
function createDefaultModel(): Model {
  return {
    money: null,
  };
}


const handleMoney = (e:number) => {
    model.value.money = e
    inputValue.value = e.toString()
}


const rechargeMoney = ref<number[]>()
const getData = async ()=>{
    uni.showLoading({
        title:"加载中"
    })
    // const res = await getSiteDictData("RECHARGE_MONEY")
    // rechargeMoney.value = res?.list.map((item)=>{
    //     return Number(item.value)
    // })
    uni.hideLoading()
}

const inputValue = ref('');
const handleInput = (e: any) => {
  // 一定要加nextTick，否则特殊情况的更改不生效【如：000时，更改为0】[citation:1]
  nextTick(() => {
    let value = e;
    // 如果当前输入为空，直接允许清空
    if (!value) {
        inputValue.value = '';
      return;
    }
    
    // 1. 当第一位为0时，第二位只能输入小数点[citation:1]
    if (value?.charAt(0) === '0' && value.charAt(1) && value.charAt(1) !== '.') {
      value = '0';
    }
    
    // 2. 如果第一位输入的是小数点，则补0[citation:1]
    if (value.length === 1 && value.charAt(0) === '.') {
      value = '0.';
    }
    
    // 3. 清除"数字"和"."以外的字符[citation:1]
    value = value.replace(/[^\d.]/g, '');
    // 4. 只保留第一个. 清除多余的[citation:1]
    value = value.replace(/\.{2,}/g, '.');
    // 5. 保留2位小数[citation:1]
    value = value.match(/^\d*(\.?\d{0,2})/g)?.[0] || '';
    inputValue.value = value;
    model.value.money = value;
  });
};

const handleRecharge = async () =>{
    if (model.value.money == null || model.value.money == 0) {
        toast.fail('请输入充值数量')
        return
    }
    toast.loading("发起支付")
    throttledPayment()
}






const { bool:fetchRechargeBoolean,setTrue:setfetchRechargeTrue,setFalse:setfetchRechargeFalse } = useBoolean()
const payment = async () =>{
    setfetchRechargeTrue()
    try {
        const res =  await postWeChatMiniProgramRecharge(model.value)
        startPolling(res.id)
        setfetchRechargeFalse()
        uni.requestPayment({
            provider:"wxpay",
            orderInfo:"",
            ...res,
            success:()=>{
                toast.loading("校验订单")
            },
            fail:()=>{
                toast.fail("取消支付")
                stopPolling();
            }
        })
    } catch (error) {
        toast.fail(error)
        setfetchRechargeFalse()
    }
}
// 创建节流支付函数 - 1秒内只能点击一次
const throttledPayment = ThrottleUtils.create(payment, {
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
        const res = await getRechargesStatus({id})
        
        // 如果订单状态为成功或失败，停止轮询
        if (res.status === PayStatus.SUCCESS) {
            stopPolling();
            toast.success("充值成功")
            // 更新用户余额
            const accountStore =  useAccountStore()
            accountStore.setAccountBalance(model.value.money)
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

onLoad(()=>{
    getData()
})
</script>