package domain


type AlibabaAlihealthSynergySyAgentpersonQueryPage struct {
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
    ResultList  *[]AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthSynergySyAgentpersonQueryPage) SetTotalNum(v int64) *AlibabaAlihealthSynergySyAgentpersonQueryPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonQueryPage) SetPage(v int64) *AlibabaAlihealthSynergySyAgentpersonQueryPage {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonQueryPage) SetPageSize(v int64) *AlibabaAlihealthSynergySyAgentpersonQueryPage {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonQueryPage) SetResultList(v []AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO) *AlibabaAlihealthSynergySyAgentpersonQueryPage {
    s.ResultList = &v
    return s
}
