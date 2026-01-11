<template>
    <sar-upload 
    :max-count="1"
    :after-read="afterRead"
    >
        <template #default="{ list, onSelect, onRemove, onImageClick }">
            <view @click="onSelect">
                <slot></slot>
            </view>
        </template>
    </sar-upload>
</template>

<script setup lang="ts">
import { useSiteStore } from '@/store/site';

const {siteInfo:site} = useSiteStore()
defineProps({
    disabled:{
        type:Boolean,
        default: false
    }
})

const emits = defineEmits(["success"])


const getType = (file:any) => {
  var filename=file;
  var index1=filename.lastIndexOf(".");
  var index2=filename.length;
  var type=filename.substring(index1,index2);
  return type;
}

const afterRead =async (prams:any)=>{
    console.log(prams)

    // if (!prams) return;

    // const file = prams.file.file
    // const type = getType(file.name)
    // if (site.mediaType.indexOf(type) == -1 ) {
    //     uni.showToast({
    //         title: `文件类型不正确`,
    //         icon: 'none',
    //     })
    //     prams.onError()
    //     prams.onFinish()
    //     return
    // }

    
    // if ((site.bigFileSize * 1024 * 1024) < file.size) {
    //     uni.showToast({
    //         title: `文件大小不能超过 ${site.bigFileSize / 1024 / 1024}MB`,
    //         icon: 'none',
    //     })
        
    //     prams.onError()
    //     prams.onFinish()
    //     return
    // }

    // uploadMinFile(prams)
    // // 如果文件尺寸小于后台设置并且小于5m则用普通上传
    // if (file.size < systemStore.base.minFileSize*1204*1024) {
    //     uploadMinFile(prams)
    //     return
    // }

    

    /* 
    // 获取分片个数
    const chunkCount = Math.ceil(file.size / chunkSize.value)
    // 获取md5
    const now = new Date();
    const nowStr = format(now, 'yyyy-MM-dd');
    const fileMd5str = file.size + file.name + userInfo?.name + userInfo?.id + nowStr
    const fileMd5 = hex_md5(fileMd5str)
    // 向后端查询验证切片上传
    
    const {data,error}= await fetchCheckChunk({
        identifier:fileMd5
    })
    if (error) {
        prams.onError()
        prams.onFinish()
        return
    }


    let chunkList = data.result
    if (chunkList.length == chunkCount) {
        // 秒传
        mergeChunk(file.name,fileMd5,file.size)
        return
    }

     // 已上传部分切片,继续上传遗留切片
    if (chunkList.length != chunkCount && chunkList.length > 0) {
        // "文件断点续传"

        let uploadedChunks = [];

        for (let i = 0; i < chunkCount; i++) {
            
            //分片开始位置
            let start = i * chunkSize.value
            let end = Math.min(file.size,start + chunkSize.value)

            let _chunkFile = file.slice(start,end)
        
            let formData = new FormData()
            formData.append("identifier",fileMd5)
            formData.append("chunkNumber",`${i}`)
            formData.append("totalChunks",`${chunkCount}`)
            formData.append("file",_chunkFile)

            uploadedChunks.push(formData)
        }
        // 需求 根据arr 里面的数据 移除 arrobj 里面对应下标数据返回一个新数组
        chunkList.forEach((item:any)=>{
            uploadedChunks[item] = undefined
        })

        uploadedChunks = uploadedChunks.filter((item)=>{
            return item != undefined
        })

        for (let index = 0; index < uploadedChunks.length; index++) {
            await fetchUploadChunk(uploadedChunks[index])
            
        }
        mergeChunk(file.name,fileMd5,file.size)
        prams.onFinish()
        return
    }

    for (let i = 0; i < chunkCount; i++) {
                
        //分片开始位置
        let start = i * chunkSize.value
        let end = Math.min(file.size,start + chunkSize.value)

        
        let _chunkFile = file.slice(start,end)
    


        let formData = new FormData()
        formData.append("identifier",fileMd5)
        formData.append("chunkNumber",`${i}`)
        formData.append("totalChunks",`${chunkCount}`)
        formData.append("file",_chunkFile)

        await fetchUploadChunk(formData)
        // await this.$axios.post(api.UploadChunk, formData)
        // this.progress =  Math.round((((i + 1) / chunkCount) * 100) - 1)
    }

    //合并分片 所有上传结束通知后端进行合并分片
    mergeChunk(file.name,fileMd5,file.size)
    prams.onFinish()
    return
    */
}

const mergeChunk = async (fileName:string, identifier:string, size:number) =>{

    // const {data} = await fetchMergeChunk({
    //     fileName: fileName,
    //     identifier: identifier,
    //     size: size,
    // })

    // emits("success",data.links[0])
}

const uploadMinFile = async(file:any)=>{
   
    // const formData = new FormData()
    // formData.append("file",file.file.file)
    
    // const {error,data} = await fetchUploadFile(formData)
    // if (error) {
    //     file.onError()
    //     return
    // }
    // file.onFinish()
    // emits("success",data.links[0])
}
</script>