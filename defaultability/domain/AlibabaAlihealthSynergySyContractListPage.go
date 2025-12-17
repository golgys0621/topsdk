package domain


type AlibabaAlihealthSynergySyContractListPage struct {
    /*
        合同列表     */
    ResultList  *[]AlibabaAlihealthSynergySyContractListSyContractModel `json:"result_list,omitempty" `

    /*
        合同数量     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractListPage) SetResultList(v []AlibabaAlihealthSynergySyContractListSyContractModel) *AlibabaAlihealthSynergySyContractListPage {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListPage) SetTotalNum(v int64) *AlibabaAlihealthSynergySyContractListPage {
    s.TotalNum = &v
    return s
}
