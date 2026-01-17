import { http } from "@/http/http"

export function getProductCategoryList(data:{gameId:number}) {
  return http.get<{
    list: Product.Dao.CategoryList[]
  }>('/product/category/list',data)
}
export function getProductList(data:Product.Dto.Query) {
  return http.get<{
    list: Product.Dao.List[],
    total: number
  }>('/product/list',data)
}
export function getProductDetail(data:{id:number}) {
  return http.get<Product.Dao.Detail>('/product/detail',data)
}
