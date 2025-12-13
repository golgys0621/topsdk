package domain


type AlibabaAlihealthDrugKytWesSearchstatusPageInfoDto struct {
    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSearchstatusPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugKytWesSearchstatusPageInfoDto {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusPageInfoDto) SetResultList(v []AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) *AlibabaAlihealthDrugKytWesSearchstatusPageInfoDto {
    s.ResultList = &v
    return s
}
