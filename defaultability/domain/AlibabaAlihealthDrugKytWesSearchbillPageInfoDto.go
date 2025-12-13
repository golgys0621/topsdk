package domain


type AlibabaAlihealthDrugKytWesSearchbillPageInfoDto struct {
    /*
        返回结果     */
    ResultList  *[]AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo `json:"result_list,omitempty" `

    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSearchbillPageInfoDto) SetResultList(v []AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) *AlibabaAlihealthDrugKytWesSearchbillPageInfoDto {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugKytWesSearchbillPageInfoDto {
    s.TotalNum = &v
    return s
}
