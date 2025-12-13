package domain


type AlibabaAlihealthSynergyYzwBillSignedQueryPage struct {
    /*
        当前页     */
    Page  *int64 `json:"page,omitempty" `

    /*
        分页大小(不能超过20)     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        集合     */
    Result  *[]AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO `json:"result,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwBillSignedQueryPage) SetPage(v int64) *AlibabaAlihealthSynergyYzwBillSignedQueryPage {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryPage) SetPageSize(v int64) *AlibabaAlihealthSynergyYzwBillSignedQueryPage {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryPage) SetResult(v []AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) *AlibabaAlihealthSynergyYzwBillSignedQueryPage {
    s.Result = &v
    return s
}
