package domain


type AlibabaAlihealthDrugLsydRetailSearchbillPageInfoDto struct {
    /*
        返回结果     */
    ResultList  *[]AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo `json:"result_list,omitempty" `

    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugLsydRetailSearchbillPageInfoDto) SetResultList(v []AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) *AlibabaAlihealthDrugLsydRetailSearchbillPageInfoDto {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugLsydRetailSearchbillPageInfoDto {
    s.TotalNum = &v
    return s
}
