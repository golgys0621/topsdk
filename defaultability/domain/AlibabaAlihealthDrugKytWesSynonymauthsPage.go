package domain


type AlibabaAlihealthDrugKytWesSynonymauthsPage struct {
    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSynonymauthsPage) SetTotalNum(v int64) *AlibabaAlihealthDrugKytWesSynonymauthsPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSynonymauthsPage) SetResultList(v []AlibabaAlihealthDrugKytWesSynonymauthsResPSynonymDTO) *AlibabaAlihealthDrugKytWesSynonymauthsPage {
    s.ResultList = &v
    return s
}
