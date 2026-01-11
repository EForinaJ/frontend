 /** The global namespace for the app */
declare namespace Bill {
    namespace Dao{
        type List = {
            id: number,
            createTime: string,
            type: number,
            money: number,
            content: string,
        }
    }
    namespace Dto{
        interface Query  {
            page: number;
            limit: number;
            type?: number
        }
    }
}
 