package domain


type AlibabaAlihealthSynergySyInboxListPage struct {
    /*
        列表     */
    ResultList  *[]AlibabaAlihealthSynergySyInboxListSyInboxModel `json:"result_list,omitempty" `

    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthSynergySyInboxListPage) SetResultList(v []AlibabaAlihealthSynergySyInboxListSyInboxModel) *AlibabaAlihealthSynergySyInboxListPage {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxListPage) SetTotalNum(v int64) *AlibabaAlihealthSynergySyInboxListPage {
    s.TotalNum = &v
    return s
}
