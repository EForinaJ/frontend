<template>
    <view 
       :style="{minHeight: getWindowHeight() + 'px'}"
       class="bg-#f5f5f5">
        <view class="bg-white w-full px-30rpx fixed box-border z-33">
            <view :style="{height:getStatusBarHeight()+'px'}" ></view>
            <view class="flex items-center"
                :style="{height:getTitleBarHeight()+'px'}" >
                <view @click="goBack" class="h-full flex items-center">
                    <sar-icon name="left" size="40rpx" />
                </view>
                <view class="flex-1 flex items-center justify-center mr-40rpx">
                    <view class="font-700
                        text-text-primary text-[36rpx] font-bold
                        leading-none">
                        服务列表
                    </view>
                </view>
            </view>
        </view>
        <view :style="{height:getStatusBarHeight()+'px'}"></view>
        <view :style="{height:getTitleBarHeight()+'px'}"></view>
    
        <view class="flex">
            <sar-sidebar
                v-model:current="current"
                round
                :root-style="{
                    position: 'sticky',
                    top: 0,
                    height: scrollViewHeight,
                    background: '#fff',
                }"
                @change="handeleCategory"
                >
                <sar-sidebar-item
                    v-for="item in categoryList"
                    :key="item.id"
                    :name="item.id"
                    :title="item.name"
                />

                <view style="height: env(safe-area-inset-bottom); flex: none"></view>
            </sar-sidebar>

            <view class="flex-1 mx-30rpx">
                <sar-pull-down-refresh
                    ref="pullDownRefresh"
                    :loading="fetchListLoading"
                    :root-style="`min-height: ${scrollViewHeight}`"
                    @refresh="onRefresh"
                >
                    <sar-scroll-spy 
                    v-model:current="current"
                    :root-style="{ height: scrollViewHeight }"
                    :offset="44"
                    >   
                    <view style="height: 20px; flex: none"></view>
                    <template>
                        <view v-if="fetchListLoading" class="mt-20rpx mx-30rpx">
                            <view v-for="item in 6" :key="item" class="p-20rpx rounded-lg bg-white mb-30rpx">
                                <sar-skeleton />
                            </view>
                        </view>
                        <view v-if="!fetchListLoading">
                            <view v-if="list.length != 0">
                                <view 
                                @click="handleDetail(item.id)"
                                v-for="item in list" :key="item.id" class="flex mb-30rpx" >
                                    <sar-image
                                        :src="item.pic"
                                        width="160rpx"
                                        height="160rpx"
                                        class="rounded-tl-lg rounded-bl-lg"
                                    />
                                    <view class="flex-1 p-14rpx flex flex-col justify-between
                                    bg-white rounded-tr-lg rounded-br-lg">
                                        <view class="text-sm font-semibold">
                                            {{ item.name }}
                                        </view>
                                        <view class="flex justify-between items-center">
                                            <view class="text-xs text-red">
                                                {{ item.price }}{{ site.symbol }}/{{ item.unit }}
                                            </view>
                                            <view class="text-xs">
                                                销量{{ item.salesCount }}
                                            </view>
                                        </view>
                                    </view>
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
                    </template>
                    <view style="height: env(safe-area-inset-bottom); flex: none"></view>
                    </sar-scroll-spy>
                </sar-pull-down-refresh>
            </view>
        </view>
        
    <sar-toast-agent />
    <sar-dialog-agent />
   </view>
</template>

<script setup lang="ts">
import { getProductCategoryList, getProductList } from '@/api/product';
import useBoolean from '@/hooks/boolean';
import { NOT_FOUND_PAGE } from '@/router/config';
import { useSiteStore } from '@/store/site';
import { getStatusBarHeight, getTitleBarHeight, getWindowHeight } from '@/utils/systemInfo';
import { LoadMoreStatus, toast } from 'sard-uniapp';

definePage({
   style: {
       navigationBarTitleText: '订单列表',
       navigationStyle: 'custom'
   },
})
const { siteInfo:site } = useSiteStore()
const statusBarHeight = getStatusBarHeight() + 'px'
const titleBarHeight = getTitleBarHeight() + 'px'
const navbarHeight = `calc(${statusBarHeight} + ${titleBarHeight})`
const scrollViewHeight = `calc(100vh - ${navbarHeight})`


const categoryList = ref<Product.Dao.CategoryList[]>([
    {
        id: 0,
        name: "全部"
    }
])
const current = ref(0)
const getCategory = async (game:number) => {
    try {
        const res = await getProductCategoryList({gameId:game})
        categoryList.value = [...categoryList.value,...res.list]
    } catch (error) {
        toast.fail(error)  
    }
}


const query = reactive<Product.Dto.Query>({
    page:1,
    limit:10,
    gameId:0,
    categoryId:0,
})
const list = ref<Product.Dao.List[]>([])
const { bool:fetchListLoading, setTrue:fetchListSetTrue, setFalse:fetchListSetFalse } = useBoolean()
const loadMoreStatus = ref<LoadMoreStatus>('incomplete')
const getList = async () =>{
    loadMoreStatus.value = 'loading'
    try {
        const res = await getProductList(query)
        list.value = [...list.value,...res.list]
        loadMoreStatus.value = list.value.length < res.total ? 'incomplete' : 'complete'
    } catch (error) {
      toast.fail(error)  
    }
}

const pullDownRefresh = ref()
onPageScroll(({ scrollTop }) => {
  pullDownRefresh.value?.enableToRefresh(scrollTop === 0)
})
const onRefresh = async () => {
    query.page  = 1
    list.value = []
    fetchListSetTrue()
    await getList()
    fetchListSetFalse()
}


const onLoadMore = async () => {
  if (!fetchListLoading.value) {
    query.page ++
    await getList()
  }
}

const onReload = async () => {
    if (!fetchListLoading.value) {
        query.page  = 1
        await getList()
    }
}
const handeleCategory = async (e:number) => {
    query.page  = 1
    query.categoryId  = e
    list.value = []
    fetchListSetTrue()
    await getList()
    fetchListSetFalse()
}

const handleDetail = (id:number) =>{
    uni.navigateTo({url: `/pages/product/detail?id=${id}`})
}
const goBack = () =>{
    uni.navigateBack()
}
onLoad(async (options)=>{
    if (options.gameId == undefined) {
        uni.navigateTo({ url: NOT_FOUND_PAGE })
    }
    query.gameId =  Number(options.gameId)
    await getCategory(Number(options.gameId))
    fetchListSetTrue()
    await getList()
    fetchListSetFalse()
})

onShareAppMessage((option)=>{
    return{
      title: site.title,
      path: `/pages/product/list?gameId=${query.gameId}`
    }
  }
)

// #ifdef MP-WEIXIN
onShareTimeline(()=>{
  return {
    title: site.title,
    path: `/pages/product/list?gameId=${query.gameId}`
  }
})
// #endif
</script>

<style lang="scss" scoped>
.sar-sidebar-item_current{
    background: #f5f5f5;
}

</style>