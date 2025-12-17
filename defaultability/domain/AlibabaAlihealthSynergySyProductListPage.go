package domain


type AlibabaAlihealthSynergySyProductListPage struct {
    /*
        列表     */
    ResultList  *[]AlibabaAlihealthSynergySyProductListSyProductDTO `json:"result_list,omitempty" `

    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthSynergySyProductListPage) SetResultList(v []AlibabaAlihealthSynergySyProductListSyProductDTO) *AlibabaAlihealthSynergySyProductListPage {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductListPage) SetTotalNum(v int64) *AlibabaAlihealthSynergySyProductListPage {
    s.TotalNum = &v
    return s
}
