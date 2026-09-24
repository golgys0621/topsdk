package domain


type AlibabaAlihealthDrugMscRetailSearchbillPageInfoDto struct {
    /*
        返回结果     */
    ResultList  *[]AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo `json:"result_list,omitempty" `

    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugMscRetailSearchbillPageInfoDto) SetResultList(v []AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) *AlibabaAlihealthDrugMscRetailSearchbillPageInfoDto {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugMscRetailSearchbillPageInfoDto {
    s.TotalNum = &v
    return s
}
