import { http } from "@/http/http"

export function getBillList(data:Bill.Dto.Query) {
  return http.get<{
    list: Bill.Dao.List[],
    total: number
  }>('/bill/list',data)
}
