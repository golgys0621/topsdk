package domain


type AlibabaAlihealthSynergySyOutboxListPage struct {
    /*
        列表     */
    ResultList  *[]AlibabaAlihealthSynergySyOutboxListSyOutboxModel `json:"result_list,omitempty" `

    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthSynergySyOutboxListPage) SetResultList(v []AlibabaAlihealthSynergySyOutboxListSyOutboxModel) *AlibabaAlihealthSynergySyOutboxListPage {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListPage) SetTotalNum(v int64) *AlibabaAlihealthSynergySyOutboxListPage {
    s.TotalNum = &v
    return s
}
