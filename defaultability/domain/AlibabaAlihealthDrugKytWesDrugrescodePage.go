package domain


type AlibabaAlihealthDrugKytWesDrugrescodePage struct {
    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto `json:"result_list,omitempty" `

    /*
        当前页     */
    Page  *int64 `json:"page,omitempty" `

    /*
        分页大小     */
    PageSize  *int64 `json:"page_size,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesDrugrescodePage) SetTotalNum(v int64) *AlibabaAlihealthDrugKytWesDrugrescodePage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodePage) SetResultList(v []AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) *AlibabaAlihealthDrugKytWesDrugrescodePage {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodePage) SetPage(v int64) *AlibabaAlihealthDrugKytWesDrugrescodePage {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodePage) SetPageSize(v int64) *AlibabaAlihealthDrugKytWesDrugrescodePage {
    s.PageSize = &v
    return s
}
