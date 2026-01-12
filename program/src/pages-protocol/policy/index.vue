<template>
  <view 
      :style="{minHeight: getWindowHeight() + 'px'}"
      class="bg-[#f5f4f9]">
      <Bar arrow
        :title="(!fetchLoading && detail) ? detail.title!: ''"
      />
      <view v-if="!fetchLoading && detail" class="mx-30rpx mt-40rpx bg-white p-20rpx rounded-lg">
        <mp-html :content="detail.content" />
      </view>
      <sar-toast-agent />
  </view>
</template>
  
<script setup lang="ts">
import { getSitePrivacyPolicy } from '@/api/site';
import Bar from '@/components/bar.vue';
import useBoolean from '@/hooks/boolean';
import {  getWindowHeight } from '@/utils/systemInfo';
import { onLoad } from '@dcloudio/uni-app';
import { toast } from 'sard-uniapp';
import { ref } from 'vue';


definePage({
  style: {
      navigationBarTitleText: '隐私政策',
      navigationStyle: 'custom'
  },
})
const { bool:fetchLoading, setTrue:fetchSetTrue, setFalse:fetchSetFalse } = useBoolean()
const detail = ref<Site.Dao.Protocol>()
const getData = async () => {
  fetchSetTrue()
  uni.showLoading({
      title:"加载中"
  })
  try {
      const res = await getSitePrivacyPolicy()
      detail.value = res
      uni.hideLoading()
  } catch (error) {
    toast.fail(error)  
  }
    fetchSetFalse()
}
  
  
  
onLoad((options)=>{
  getData()
})
</script>