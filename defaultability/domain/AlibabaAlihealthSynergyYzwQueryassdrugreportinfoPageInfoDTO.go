package domain


type AlibabaAlihealthSynergyYzwQueryassdrugreportinfoPageInfoDTO struct {
    /*
        内层小对象     */
    ResultList  *[]AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO `json:"result_list,omitempty" `

    /*
        数据总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        页码     */
    Page  *int64 `json:"page,omitempty" `

    /*
        页的大小     */
    PageSize  *int64 `json:"page_size,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoPageInfoDTO) SetResultList(v []AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoPageInfoDTO {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoPageInfoDTO) SetTotalNum(v int64) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoPageInfoDTO {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoPageInfoDTO) SetPage(v int64) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoPageInfoDTO {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoPageInfoDTO) SetPageSize(v int64) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoPageInfoDTO {
    s.PageSize = &v
    return s
}
