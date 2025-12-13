package domain


type AlibabaAlihealthDrugKytWesListbillprocesspartsuccessPageInfoDTO struct {
    /*
        返回列表     */
    Result  *[]AlibabaAlihealthDrugKytWesListbillprocesspartsuccessBillProcessPartSuccessExcelDTO `json:"result,omitempty" `

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

func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessPageInfoDTO) SetResult(v []AlibabaAlihealthDrugKytWesListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessPageInfoDTO {
    s.Result = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessPageInfoDTO) SetPages(v int64) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessPageInfoDTO {
    s.Pages = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessPageInfoDTO) SetTotalNum(v int64) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessPageInfoDTO {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessPageInfoDTO) SetPageSize(v int64) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessPageInfoDTO {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessPageInfoDTO) SetPage(v int64) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessPageInfoDTO {
    s.Page = &v
    return s
}
