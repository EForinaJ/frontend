<template>
    <view 
       :style="{minHeight: getWindowHeight() + 'px'}"
       class="bg-[#f5f4f9]">
       <Bar arrow
        title="商品分享"
       />
       <l-painter :board="poster"/>
    </view>
</template>

<script setup lang="ts">
import lPainter from '@/components/painter/components/l-painter/l-painter.vue';
import { getProductDetail, getProductWxMiniProgramCode } from '@/api/product';
import Bar from '@/components/bar.vue';
import useBoolean from '@/hooks/boolean';
import { NOT_FOUND_PAGE } from '@/router/config';
import { useSiteStore } from '@/store/site';
import {  getStatusBarHeight, getTitleBarHeight, getWindowHeight } from '@/utils/systemInfo';
import {  toast } from 'sard-uniapp';
import { getImageUrl } from '@/utils/imageHelper';

definePage({
   style: {
       navigationBarTitleText: '商品分享',
       navigationStyle: 'custom'
   },
})
const id = ref<number>(0)
const {siteInfo:site} = useSiteStore()

const poster = ref()

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
        await  getCode()


        poster.value = {
            css: {
                width: "750rpx",
                paddingBottom: "40rpx",
                background: "linear-gradient(,#000 0%, #ff5000 100%)"
            },
            views: [
                // {
                //     src: detail.value.pic,
                //     type: "image",
                //     css: {
                //         background: "#fff",
                //         objectFit: "cover",
                //         marginLeft: "40rpx",
                //         marginTop: "40rpx",
                //         width: "84rpx",
                //         border: "2rpx solid #fff",
                //         boxSizing: "border-box",
                //         height: "84rpx",
                //         borderRadius: "50%"
                //     }
                // },
                // {
                //     type: "view",
                //     css: {
                //         marginTop: "40rpx",
                //         paddingLeft: "20rpx",
                //         display: "inline-block"
                //     },
                //     views: [
                //         {
                //             text: "隔壁老王",
                //             type: "text",
                //             css: {
                //                 display: "block",
                //                 paddingBottom: "10rpx",
                //                 color: "#fff",
                //                 fontSize: "32rpx",
                //                 fontWeight: "bold"
                //             }
                //         },
                //         {
                //             text: "为您挑选了一个好物",
                //             type: "text",
                //             css: {
                //                 color: "rgba(255,255,255,.7)",
                //                 fontSize: "24rpx"
                //             },
                //         }
                //     ],
                // },
                {
                    css: {
                        marginLeft: "40rpx",
                        marginTop: "30rpx",
                        padding: "32rpx",
                        boxSizing: "border-box",
                        background: "#fff",
                        borderRadius: "16rpx",
                        width: "670rpx",
                        boxShadow: "0 20rpx 58rpx rgba(0,0,0,.15)"
                    },
                    views: [
                        {
                            src: detail.value.pic,
                            type: "image",
                            css: {
                                borderRadius: "16rpx",
                                objectFit: "cover",
                                objectPosition: "50% 50%",
                                width: "606rpx",
                                height: "606rpx"
                            },
                        },
                        {
                            css: {
                                marginTop: "32rpx",
                                color: "#FF0000",
                                fontWeight: "bold",
                                fontSize: "28rpx",
                                lineHeight: "1em"
                            },
                            views: [{
                                text: detail.value.price,
                                type: "text",
                                css: {
                                    verticalAlign: "bottom",
                                    fontSize: "58rpx"
                                },
                            }, 
                            {
                                text: site.symbol,
                                type: "text",
                                css: {
                                    verticalAlign: "bottom"
                                },
                            },
                            {
                                text: `销量${detail.value.salesCount}`,
                                type: "text",
                                css: {
                                    verticalAlign: "bottom",
                                    paddingLeft: "10rpx",
                                    fontWeight: "normal",
                                    color: "#999999"
                                }
                            }],

                            type: "view"
                        }, 
                        {
                            css: {
                                marginTop: "32rpx",
                                fontSize: "26rpx",
                                color: "#8c5400"
                            },
                            views: [{
                                    text: detail.value.game,
                                    type: "text",
                                    css: {
                                        padding: "4rpx 16rpx",
                                        color: "#fff",
                                        background: "#ffb400",
                                        borderRadius: "16rpx",
                                    },
                                },
                                {
                                    text: detail.value.category,
                                    type: "text",
                                    css: {
                                        padding: "4rpx 16rpx",
                                        marginLeft: "16rpx",
                                        color: "#fff",
                                        background: "#00BFFF",
                                        borderRadius: "16rpx",
                                    },
                                }, 
                                // {
                                //     text: "30天最低价",
                                //     type: "text",
                                //     css: {
                                //         marginLeft: "16rpx",
                                //         background: "#fff4d9",
                                //         textDecoration: "line-through"
                                //     },
                                // }, 
                                // {
                                //     text: "满减优惠",
                                //     type: "text",
                                //     css: {
                                //         marginLeft: "16rpx",
                                //         background: "#fff4d9"
                                //     },
                                // },
                                // {
                                //     text: "超高好评",
                                //     type: "text",
                                //     css: {
                                //         marginLeft: "16rpx",
                                //         background: "#fff4d9"
                                //     },

                                // }
                            ],
                            type: "view"
                        }, 
                        {
                            css: {
                                marginTop: "20rpx"
                            },
                            views: [
                                {
                                    text: detail.value.name,
                                    type: "text",
                                    css: {
                                        paddingRight: "32rpx",
                                        boxSizing: "border-box",
                                        lineClamp: 2,
                                        color: "#333333",
                                        lineHeight: "1.8em",
                                        fontSize: "36rpx",
                                        width: "478rpx"
                                    },
                                }, 
                                {
                                    src: qrCode.value,
                                    type: "image",
                                    css: {
                                        borderRadius: "16rpx",
                                        objectFit: "cover",
                                        objectPosition: "50% 50%",
                                        width: "128rpx",
                                        height: "128rpx",
                                    },
                                },
                            ],
                            type: "view"
                        }],
                    type: "view"
                }
            ]
        }
    } catch (error) {
      toast.fail(error)  
    }
    fetchSetFalse()
}
const qrCode = ref<string>("")
const getCode = async () =>{
    try {
        const res = await getProductWxMiniProgramCode({id:id.value})
        qrCode.value = getImageUrl(res.qrCode);
        uni.hideLoading()
    } catch (error) {
      toast.fail(error)  
    }
    
}

onLoad((options)=>{
    if (options.id == undefined) {
        uni.navigateTo({ url: NOT_FOUND_PAGE })
    }
    id.value = Number(options.id)
    getData()
})
</script>

<style lang="scss" scoped>
</style>