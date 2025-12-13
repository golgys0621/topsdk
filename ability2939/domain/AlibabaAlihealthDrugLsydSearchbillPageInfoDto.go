package domain


type AlibabaAlihealthDrugLsydSearchbillPageInfoDto struct {
    /*
        返回结果     */
    ResultList  *[]AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo `json:"result_list,omitempty" `

    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugLsydSearchbillPageInfoDto) SetResultList(v []AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) *AlibabaAlihealthDrugLsydSearchbillPageInfoDto {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugLsydSearchbillPageInfoDto {
    s.TotalNum = &v
    return s
}
