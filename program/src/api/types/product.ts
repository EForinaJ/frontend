 /** The global namespace for the app */
 declare namespace Product {
    namespace Dao{
        type CategoryList = {
            id: number,
            name: string,
        }

        type List = {
            id: number,
            name: string,
            pic: string,
            category: string,
            game: string,
            price: number,
            salesCount: number,
            unit: string,
            rate: number,
        }

        type Detail = {
            id: number,
            code: string,
            name: string,
            pic: string,
            category: string,
            game: string,
            price: number,
            salesCount: number,
            unit: string,
            rate: number,
            content: string,
            description: string,
        }
    }
    namespace Dto{
        interface Query  {
            page: number;
            limit: number;
            gameId?: number
            categoryId?: number
        }
    }
}
 