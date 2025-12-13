package domain


type AlibabaAlihealthDrugYljgSearchbillPageInfoDTO struct {
    /*
        返回结果     */
    ResultList  *[]AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO `json:"result_list,omitempty" `

    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugYljgSearchbillPageInfoDTO) SetResultList(v []AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) *AlibabaAlihealthDrugYljgSearchbillPageInfoDTO {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillPageInfoDTO) SetTotalNum(v int64) *AlibabaAlihealthDrugYljgSearchbillPageInfoDTO {
    s.TotalNum = &v
    return s
}
