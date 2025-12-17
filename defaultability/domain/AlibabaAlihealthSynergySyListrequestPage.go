package domain


type AlibabaAlihealthSynergySyListrequestPage struct {
    /*
        列表     */
    ResultList  *[]AlibabaAlihealthSynergySyListrequestSyRequestDTO `json:"result_list,omitempty" `

    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthSynergySyListrequestPage) SetResultList(v []AlibabaAlihealthSynergySyListrequestSyRequestDTO) *AlibabaAlihealthSynergySyListrequestPage {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListrequestPage) SetTotalNum(v int64) *AlibabaAlihealthSynergySyListrequestPage {
    s.TotalNum = &v
    return s
}
