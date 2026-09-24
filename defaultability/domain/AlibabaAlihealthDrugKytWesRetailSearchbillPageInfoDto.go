package domain


type AlibabaAlihealthDrugKytWesRetailSearchbillPageInfoDto struct {
    /*
        返回结果     */
    ResultList  *[]AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo `json:"result_list,omitempty" `

    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesRetailSearchbillPageInfoDto) SetResultList(v []AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) *AlibabaAlihealthDrugKytWesRetailSearchbillPageInfoDto {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugKytWesRetailSearchbillPageInfoDto {
    s.TotalNum = &v
    return s
}
