package domain


type AlibabaAlihealthDrugKytWesQueryupoutbilllogPageInfoDTO struct {
    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO `json:"result_list,omitempty" `

    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogPageInfoDTO) SetResultList(v []AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO) *AlibabaAlihealthDrugKytWesQueryupoutbilllogPageInfoDTO {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogPageInfoDTO) SetTotalNum(v int64) *AlibabaAlihealthDrugKytWesQueryupoutbilllogPageInfoDTO {
    s.TotalNum = &v
    return s
}
