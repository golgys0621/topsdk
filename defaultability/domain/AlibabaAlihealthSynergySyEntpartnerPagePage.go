package domain


type AlibabaAlihealthSynergySyEntpartnerPagePage struct {
    /*
        合作企业     */
    ResultList  *[]AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel `json:"result_list,omitempty" `

    /*
        合作企业数量     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthSynergySyEntpartnerPagePage) SetResultList(v []AlibabaAlihealthSynergySyEntpartnerPageSyEntPartnerListModel) *AlibabaAlihealthSynergySyEntpartnerPagePage {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPagePage) SetTotalNum(v int64) *AlibabaAlihealthSynergySyEntpartnerPagePage {
    s.TotalNum = &v
    return s
}
