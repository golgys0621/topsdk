package domain


type AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessPageInfoDTO struct {
    /*
        返回列表     */
    Result  *[]AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessBillProcessPartSuccessExcelDTO `json:"result,omitempty" `

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

func (s *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessPageInfoDTO) SetResult(v []AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessPageInfoDTO {
    s.Result = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessPageInfoDTO) SetPages(v int64) *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessPageInfoDTO {
    s.Pages = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessPageInfoDTO) SetTotalNum(v int64) *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessPageInfoDTO {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessPageInfoDTO) SetPageSize(v int64) *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessPageInfoDTO {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessPageInfoDTO) SetPage(v int64) *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessPageInfoDTO {
    s.Page = &v
    return s
}
