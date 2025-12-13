package domain


type AlibabaAlihealthDrugMscListbillprocesspartsuccessPageInfoDTO struct {
    /*
        返回列表     */
    Result  *[]AlibabaAlihealthDrugMscListbillprocesspartsuccessBillProcessPartSuccessExcelDTO `json:"result,omitempty" `

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

func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessPageInfoDTO) SetResult(v []AlibabaAlihealthDrugMscListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) *AlibabaAlihealthDrugMscListbillprocesspartsuccessPageInfoDTO {
    s.Result = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessPageInfoDTO) SetPages(v int64) *AlibabaAlihealthDrugMscListbillprocesspartsuccessPageInfoDTO {
    s.Pages = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessPageInfoDTO) SetTotalNum(v int64) *AlibabaAlihealthDrugMscListbillprocesspartsuccessPageInfoDTO {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessPageInfoDTO) SetPageSize(v int64) *AlibabaAlihealthDrugMscListbillprocesspartsuccessPageInfoDTO {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessPageInfoDTO) SetPage(v int64) *AlibabaAlihealthDrugMscListbillprocesspartsuccessPageInfoDTO {
    s.Page = &v
    return s
}
