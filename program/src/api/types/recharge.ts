

 /** The global namespace for the app */
declare namespace Recharge {
  namespace Dao{
      type WechatMiniProgram = {
          id: number;
          appId: string;
          timeStamp: string;
          nonceStr: string;
          package: string;
          signType: string;
          paySign: string;
      }
  }
  namespace Dto{
      interface Create  {
        money: number | null
      }
  }
}
