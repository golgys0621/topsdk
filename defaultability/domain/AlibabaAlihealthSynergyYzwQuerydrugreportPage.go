package domain


type AlibabaAlihealthSynergyYzwQuerydrugreportPage struct {
    /*
        当前页     */
    Page  *int64 `json:"page,omitempty" `

    /*
        分页大小(不能超过20)     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        结果     */
    Result  *[]AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO `json:"result,omitempty" `

    /*
        总记录数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQuerydrugreportPage) SetPage(v int64) *AlibabaAlihealthSynergyYzwQuerydrugreportPage {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportPage) SetPageSize(v int64) *AlibabaAlihealthSynergyYzwQuerydrugreportPage {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportPage) SetResult(v []AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) *AlibabaAlihealthSynergyYzwQuerydrugreportPage {
    s.Result = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportPage) SetTotalNum(v int64) *AlibabaAlihealthSynergyYzwQuerydrugreportPage {
    s.TotalNum = &v
    return s
}
