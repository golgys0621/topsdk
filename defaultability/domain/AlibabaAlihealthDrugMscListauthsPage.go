package domain


type AlibabaAlihealthDrugMscListauthsPage struct {
    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回集合     */
    ResultList  *[]AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthDrugMscListauthsPage) SetTotalNum(v int64) *AlibabaAlihealthDrugMscListauthsPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsPage) SetResultList(v []AlibabaAlihealthDrugMscListauthsPSynonymUserEntInfoDTO) *AlibabaAlihealthDrugMscListauthsPage {
    s.ResultList = &v
    return s
}
