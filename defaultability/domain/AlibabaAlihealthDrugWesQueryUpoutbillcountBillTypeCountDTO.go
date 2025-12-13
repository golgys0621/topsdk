package domain


type AlibabaAlihealthDrugWesQueryUpoutbillcountBillTypeCountDTO struct {
    /*
        单据类型枚举     */
    BillType  *string `json:"bill_type,omitempty" `

    /*
        单据类型     */
    BillTypeName  *string `json:"bill_type_name,omitempty" `

    /*
        单据数量     */
    Count  *int64 `json:"count,omitempty" `

}

func (s *AlibabaAlihealthDrugWesQueryUpoutbillcountBillTypeCountDTO) SetBillType(v string) *AlibabaAlihealthDrugWesQueryUpoutbillcountBillTypeCountDTO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryUpoutbillcountBillTypeCountDTO) SetBillTypeName(v string) *AlibabaAlihealthDrugWesQueryUpoutbillcountBillTypeCountDTO {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryUpoutbillcountBillTypeCountDTO) SetCount(v int64) *AlibabaAlihealthDrugWesQueryUpoutbillcountBillTypeCountDTO {
    s.Count = &v
    return s
}
