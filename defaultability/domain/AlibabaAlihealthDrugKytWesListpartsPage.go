package domain


type AlibabaAlihealthDrugKytWesListpartsPage struct {
    /*
        仅代表本次查询返回的记录数，小于等于page_size时，可根据需要继续请求下一页     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugKytWesListpartsPEntParDto `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesListpartsPage) SetTotalNum(v int64) *AlibabaAlihealthDrugKytWesListpartsPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsPage) SetResultList(v []AlibabaAlihealthDrugKytWesListpartsPEntParDto) *AlibabaAlihealthDrugKytWesListpartsPage {
    s.ResultList = &v
    return s
}
