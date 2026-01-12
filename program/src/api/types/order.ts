 /** The global namespace for the app */
 declare namespace Order {
    namespace Dao{
        type List = {
            id: number,
            product: product,
            quantity: number,
            unit: string,
            unitPrice: string,
            totalAmount: number,
            status: number,
            createTime: string,
        }

        type Detail = {
            id: number,
            code: string,
            product: product,
            unit: string,
            quantity: number,
            unitPrice: string,
            totalAmount: number,
            discountAmount: number,
            actualAmount: number,
            payAmount: number,
            requirements: string,
            status: number,
            createTime: string | null,
            startTime: string | null,
            finishTime: string | null,
            payTime: string | null,
        }

        type product = {
            id: number,
            name: string,
            pic: string,
            game: string,
            category: string,
        }

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
        interface Query  {
            page: number;
            limit: number;
            status?: number
        }
    }
}
 