package domain


type AlibabaAlihealthDrugYljgRetailSearchbillPageInfoDTO struct {
    /*
        返回结果     */
    ResultList  *[]AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO `json:"result_list,omitempty" `

    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugYljgRetailSearchbillPageInfoDTO) SetResultList(v []AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) *AlibabaAlihealthDrugYljgRetailSearchbillPageInfoDTO {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillPageInfoDTO) SetTotalNum(v int64) *AlibabaAlihealthDrugYljgRetailSearchbillPageInfoDTO {
    s.TotalNum = &v
    return s
}
