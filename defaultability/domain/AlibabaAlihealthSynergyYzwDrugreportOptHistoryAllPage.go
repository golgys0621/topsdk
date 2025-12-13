package domain


type AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllPage struct {
    /*
        数据总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        页码     */
    Page  *int64 `json:"page,omitempty" `

    /*
        页的大小     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        结果列表     */
    ResultList  *[]AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllPage) SetTotalNum(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllPage) SetPage(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllPage {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllPage) SetPageSize(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllPage {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllPage) SetResultList(v []AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllPage {
    s.ResultList = &v
    return s
}
