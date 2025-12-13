package domain


type AlibabaAlihealthSynergyYzwQuerydrugbatchreportPage struct {
    /*
        数据总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        结果列表     */
    ResultList  *[]AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO `json:"result_list,omitempty" `

    /*
        页码     */
    Page  *int64 `json:"page,omitempty" `

    /*
        页的大小     */
    PageSize  *int64 `json:"page_size,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportPage) SetTotalNum(v int64) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportPage) SetResultList(v []AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportPage {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportPage) SetPage(v int64) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportPage {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportPage) SetPageSize(v int64) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportPage {
    s.PageSize = &v
    return s
}
