package domain


type AlibabaAlihealthSynergyYzwQuerydrugreportassPage struct {
    /*
        当前页 defalutValue:1    */
    Page  *int64 `json:"page,omitempty" `

    /*
        分页大小(不能超过20) defalutValue:20    */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        结果     */
    Result  *[]AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO `json:"result,omitempty" `

    /*
        总记录数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassPage) SetPage(v int64) *AlibabaAlihealthSynergyYzwQuerydrugreportassPage {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassPage) SetPageSize(v int64) *AlibabaAlihealthSynergyYzwQuerydrugreportassPage {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassPage) SetResult(v []AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) *AlibabaAlihealthSynergyYzwQuerydrugreportassPage {
    s.Result = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassPage) SetTotalNum(v int64) *AlibabaAlihealthSynergyYzwQuerydrugreportassPage {
    s.TotalNum = &v
    return s
}
