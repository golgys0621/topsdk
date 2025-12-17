package domain


type AlibabaAlihealthSynergySyPartneragentpersonQueryPage struct {
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
    ResultList  *[]AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryPage) SetTotalNum(v int64) *AlibabaAlihealthSynergySyPartneragentpersonQueryPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryPage) SetPage(v int64) *AlibabaAlihealthSynergySyPartneragentpersonQueryPage {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryPage) SetPageSize(v int64) *AlibabaAlihealthSynergySyPartneragentpersonQueryPage {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryPage) SetResultList(v []AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO) *AlibabaAlihealthSynergySyPartneragentpersonQueryPage {
    s.ResultList = &v
    return s
}
