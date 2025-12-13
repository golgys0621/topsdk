package domain


type AlibabaAlihealthSynergyYzwSignednotstampedbillQueryPage struct {
    /*
        当前页     */
    Page  *int64 `json:"page,omitempty" `

    /*
        分页大小(不能超过20)     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        结果     */
    Result  *[]AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO `json:"result,omitempty" `

    /*
        总记录数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryPage) SetPage(v int64) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryPage {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryPage) SetPageSize(v int64) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryPage {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryPage) SetResult(v []AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryPage {
    s.Result = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryPage) SetTotalNum(v int64) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryPage {
    s.TotalNum = &v
    return s
}
