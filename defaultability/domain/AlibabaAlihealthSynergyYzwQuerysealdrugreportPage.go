package domain


type AlibabaAlihealthSynergyYzwQuerysealdrugreportPage struct {
    /*
        数据总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        结果列表     */
    ResultList  *[]AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO `json:"result_list,omitempty" `

    /*
        页码     */
    Page  *int64 `json:"page,omitempty" `

    /*
        页的大小     */
    PageSize  *int64 `json:"page_size,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportPage) SetTotalNum(v int64) *AlibabaAlihealthSynergyYzwQuerysealdrugreportPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportPage) SetResultList(v []AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO) *AlibabaAlihealthSynergyYzwQuerysealdrugreportPage {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportPage) SetPage(v int64) *AlibabaAlihealthSynergyYzwQuerysealdrugreportPage {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportPage) SetPageSize(v int64) *AlibabaAlihealthSynergyYzwQuerysealdrugreportPage {
    s.PageSize = &v
    return s
}
