<template>
<view  
:style="{minHeight: getWindowHeight()-(getTitleBarHeight()+getStatusBarHeight())+ 'px'}"
class="bg-[#f5f4f9]" >
    

    <view class="pt-40rpx mx-30rpx">
    <sar-list hide-border card>
        <sar-list-item arrow @click="handleAvatar">
        <template #title>
            <text class="text-text-primary">头像</text>
        </template>
        <template #value>
            <sar-avatar size="84rpx" icon-size="48rpx" v-if="avatar"  :src="avatar" />
            <sar-avatar size="84rpx" icon-size="48rpx" v-if="!avatar" :src="account.avatar" />
        </template>
        </sar-list-item>
    </sar-list>
    </view>

    <view class="mt-40rpx mx-30rpx">
    <sar-list hide-border card>
        <sar-list-item @click="handleName" arrow>
        <template #title>
            <text class="text-text-primary">昵称</text>
        </template>
        <template #value>
            <text class="text-text-secondary">{{ account.name }}</text>
        </template>
        </sar-list-item>
        <sar-list-item @click="handleSex" arrow>
        <template #title>
            <text class="text-text-primary">性别</text>
        </template>
        <template #value>
            <text class="text-text-secondary">{{ sex }}</text>
        </template>
        </sar-list-item>
        <sar-list-item @click="handleAge" arrow>
        <template #title>
            <text class="text-text-primary">出生日期</text>
        </template>
        <template #value>
            <text class="text-text-secondary">{{ formatDate(new Date(account.birthday), 'YYYY-MM-DD') }}</text>
        </template>
        </sar-list-item>
        <sar-list-item>
        <template #title>
            <text class="text-text-primary">手机号</text>
        </template>
        <template #value>
            <text class="text-text-secondary">{{ account.phone }}</text>
        </template>
        </sar-list-item>
        <sar-list-item @click="handleDescription" arrow>
        <template #title>
            <text class="text-text-primary">个人简介</text>
        </template>
        <template #value>
            <text class="text-text-secondary">{{ account.description == "" ? "请输入个人简介"  : account.description  }}</text>
        </template>
        </sar-list-item>
    </sar-list>
    </view>


    <!-- 年龄 -->
    <sar-datetime-picker-popout
    v-model="birthday"
    v-model:visible="ageVisible"
    title="请选择日期"
    type="yMd"
    :min="new Date(1970, 1, 1)"
    @change="handleSubmitAge"
    />


    <!-- 修改性别 -->
    <sar-action-sheet
        @select="handleSubmitSex"
        @close="closeSex"
    v-model:visible="sexVisible"
    cancel="取消"
    :item-list="[
        {name: '男'},
        {name: '女'},
        {name: '其他'},
    ]"
    />

    <!-- 修改昵称 -->
    <sar-popout
    :show-cancel="false"
    show-close
    title="修改名称"
    v-model:visible="nameVisible"
    @close="closeName"
    @confirm="handleSubmitName"
    >
    <view class="bg-[#f5f4f9] p-10rpx rounded-lg  mx-30rpx">
        <sar-input 
        root-style="
            font-weight: 700;
        "
        placeholder-style="
            color: rgba(0,0,0,.3);
        "
        v-model="model.name"
        borderless
        placeholder="请输入用户名称">
        </sar-input>
    </view>
    </sar-popout>



    <!-- 修改简介 -->
    <sar-popout
    :show-cancel="false"
    show-close
    title="个人简介"
    v-model:visible="descriptionVisible"
    @close="closeDescription"
    @confirm="handleSubmitDescription"
    >
    <view class="bg-[#f5f4f9] p-10rpx rounded-lg  mx-30rpx">
        <sar-input 
        root-style="
            font-weight: 700;
        "
        placeholder-style="
            color: rgba(0,0,0,.3);
        "
        type="textarea"
        autoHeight
        v-model="model.description"
        borderless
        placeholder="请输入个人简介">
        </sar-input>
    </view>
    </sar-popout>
</view>
</template>
    
    
<script lang="ts" setup>
import { storeToRefs } from 'pinia'
import { useAccountStore } from '@/store'
import { useTokenStore } from '@/store/token'
import { getStatusBarHeight, getTitleBarHeight, getWindowHeight } from '@/utils/systemInfo';
import { IAccountDto } from '@/api/types/account';
import { postAccountEdit } from '@/api/user';
import useBoolean from '@/hooks/boolean';
import { SexType } from '@/enum/type';
import { ActionSheetItem, formatDate } from 'sard-uniapp';


definePage({
    needLogin: true,
    style: {
    navigationBarTitleText: '个人信息',
    },
})

const accountStore = useAccountStore()
const tokenStore = useTokenStore()
// 使用storeToRefs解构accountInfo
const { accountInfo:account } = storeToRefs(accountStore)
const { data:avatar, run, } = useUpload()
const model = ref<IAccountDto>({
    name: account.value.name,
    sex: account.value.sex,
    address: account.value.address,
    birthday:  account.value.birthday,
    avatar:  account.value.avatar,
    description: account.value.description,
})



// 年龄
const birthday = ref<Date | undefined>(new Date(account.value.birthday))
const {bool:ageVisible,setTrue:setAgeTrue,setFalse:setAgeFalse} = useBoolean()
const handleAge = () =>{
    setAgeTrue()
}
const handleSubmitAge = async (value: Date) =>{
    model.value.birthday = value.getTime()
    await postAccountEdit(model.value)
    accountStore.setAccountBirthday(model.value.birthday)
    uni.showToast({
    title: '修改成功',
    icon: 'success',
    })
}

// 修改性别
const sex = computed(()=>{
    if (model.value.sex == SexType.Male) {
    return "男"
    }
    if (model.value.sex == SexType.Female) {
    return "女"
    }
    if (model.value.sex == SexType.Other) {
    return "其他"
    }
})
const {bool:sexVisible,setTrue:setSexTrue,setFalse:setSexFalse} = useBoolean()
const handleSex = () =>{
    setSexTrue()
}
const closeSex = () =>{
    model.value.sex = account.value.sex
    setSexFalse()
}
const handleSubmitSex =  async (item: ActionSheetItem, index: number) =>{
    if (item.name == "男") {
    model.value.sex = SexType.Male
    }
    if (item.name == "女") {
    model.value.sex = SexType.Female
    }
    if (item.name == "其他") {
    model.value.sex = SexType.Other
    }
    await postAccountEdit(model.value)
    accountStore.setAccountSex(model.value.sex)
    uni.showToast({
    title: '修改成功',
    icon: 'success',
    })
}

// 修改昵称
const {bool:nameVisible,setTrue:setNameTrue,setFalse:setNameFalse} = useBoolean()
const handleName = () =>{
    setNameTrue()
}
const closeName = () =>{
    model.value.name = account.value.name
    setNameFalse()
}
const handleSubmitName =  async () =>{
    await postAccountEdit(model.value)
    accountStore.setAccountName(model.value.name)
    uni.showToast({
    title: '修改成功',
    icon: 'success',
    })
}


// 修改简介
const {bool:descriptionVisible,setTrue:setDescriptionTrue,setFalse:setDescriptionFalse} = useBoolean()
const handleDescription = () =>{
    setDescriptionTrue()
}
const closeDescription = () =>{
    model.value.description = account.value.description
    setDescriptionFalse()
}
const handleSubmitDescription =  async () =>{
    await postAccountEdit(model.value)
    accountStore.setAccountDescription(model.value.description)
    uni.showToast({
    title: '修改成功',
    icon: 'success',
    })
}

// 修改头像
const handleAvatar = () => {
    run()
}
watch(
    () => avatar.value,
    async (newInfo) => {
        try {
            model.value.avatar = newInfo
            await postAccountEdit(model.value)
            accountStore.setAccountAvatar(newInfo)
            uni.showToast({
                title: '修改成功',
                icon: 'success',
            })
        } catch (error) {
            throw error
        }
    },
    { deep: true } // 深度监听嵌套对象
);

</script>
    
    
<style lang="scss">

</style>
