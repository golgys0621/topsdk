package domain


type AlibabaAlihealthDrugKytWesListupoutPageInfoDto struct {
    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo `json:"result_list,omitempty" `

    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesListupoutPageInfoDto) SetResultList(v []AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) *AlibabaAlihealthDrugKytWesListupoutPageInfoDto {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugKytWesListupoutPageInfoDto {
    s.TotalNum = &v
    return s
}
