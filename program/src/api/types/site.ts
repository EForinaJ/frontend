 /** The global namespace for the app */
 declare namespace Site {
    namespace Dao{
        type Info = {
            title: string;       // 网站标题
            domain: string;      // 网站域名
            logo: string;        // 网站Logo图片地址
            icon: string;        // 网站图标地址
            description: string; // 网站描述
            address: string;     // 网站地址
            icp: string;         // ICP备案号
            copyright: string;   // 版权信息
            fileSize : number,
            imageSize: number,
            fileType: Array<string>,
            imageType: Array<string>,
            symbol: string;   // 货币类型
            contact: {
                platform:string
                number:string
                icon:string
                wechat:string
            }
        }
        type Protocol = {
            title: string;
            content: string;
        }

        type Home = {
            gameList: {
                id:number;
                name:string;
                pic:string;
            }[];
            productList: {
                game:{
                    id:number;
                    name:string;
                    pic:string;
                };
                products:{
                    id:number;
                    name:string;
                    pic:string;
                }[]
            }[];
        }
    }
}