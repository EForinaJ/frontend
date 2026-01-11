import type { IAccountInfoRes } from '@/api/types/login'
import { getAccountInfo, wxPhoneBind as _wxPhoneBind } from '@/api/user'
import { defineStore } from 'pinia'
import { ref } from 'vue'


// 初始化状态
const accountInfoState: IAccountInfoRes = {
  id: 0,
  name: "",
  avatar: "",
  address: "",
  phone: "",
  sex: 0,
  balance: 0,
  description: "",
  birthday: 0,
  experience:0,
  isSign:false,
  level:{
      id:0,
      icon:"",
      name:""
  }
}

export const useAccountStore = defineStore(
  'account',
  () => {
    // 定义用户信息
    const accountInfo = ref<IAccountInfoRes>({ ...accountInfoState })
    // 设置用户信息
    const setAccountInfo = (val: IAccountInfoRes) => {
      // 若头像为空 则使用默认头像
      if (!val.avatar) {
        val.avatar = accountInfoState.avatar
      }
      accountInfo.value = val
    }
    const setAccountAvatar = (avatar: string) => {
      accountInfo.value.avatar = avatar
    }
    const setAccountName = (name: string) => {
      accountInfo.value.name = name
    }
    const setAccountDescription = (description: string) => {
      accountInfo.value.description = description
    }
    const setAccountSex = (sex: number) => {
      accountInfo.value.sex = sex
    }
    const setAccountBirthday = (birthday: number) => {
      accountInfo.value.birthday = birthday
    }
    const setAccountBalance = (balance: number) => {
      accountInfo.value.balance =  accountInfo.value.balance + balance
    }
    // 删除用户信息
    const clearAccountInfo = () => {
      accountInfo.value = { ...accountInfoState }
      uni.removeStorageSync('account')
    }

    /**
     * 获取用户信息
     */
    const fetchAccountInfo = async () => {
      const res = await getAccountInfo()
      setAccountInfo(res)
      return res
    }


    return {
      accountInfo,
      clearAccountInfo,
      fetchAccountInfo,
      setAccountInfo,
      setAccountAvatar,
      setAccountName,
      setAccountSex,
      setAccountBirthday,
      setAccountDescription,
      setAccountBalance,
    }
  },
  {
    persist: true,
  },
)
