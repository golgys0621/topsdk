package domain


type AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessPageInfoDTO struct {
    /*
        返回列表     */
    Result  *[]AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessBillProcessPartSuccessExcelDTO `json:"result,omitempty" `

    /*
        pages     */
    Pages  *int64 `json:"pages,omitempty" `

    /*
        totalNum     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        pageSize     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        page     */
    Page  *int64 `json:"page,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessPageInfoDTO) SetResult(v []AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessPageInfoDTO {
    s.Result = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessPageInfoDTO) SetPages(v int64) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessPageInfoDTO {
    s.Pages = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessPageInfoDTO) SetTotalNum(v int64) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessPageInfoDTO {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessPageInfoDTO) SetPageSize(v int64) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessPageInfoDTO {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessPageInfoDTO) SetPage(v int64) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessPageInfoDTO {
    s.Page = &v
    return s
}
