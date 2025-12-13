package domain


type AlibabaAlihealthDrugWesListauthsPage struct {
    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回集合     */
    ResultList  *[]AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthDrugWesListauthsPage) SetTotalNum(v int64) *AlibabaAlihealthDrugWesListauthsPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugWesListauthsPage) SetResultList(v []AlibabaAlihealthDrugWesListauthsPSynonymUserEntInfoDTO) *AlibabaAlihealthDrugWesListauthsPage {
    s.ResultList = &v
    return s
}
