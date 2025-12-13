package domain


type AlibabaAlihealthDrugWesQueryBillcountBillTypeCountDTO struct {
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

func (s *AlibabaAlihealthDrugWesQueryBillcountBillTypeCountDTO) SetBillType(v string) *AlibabaAlihealthDrugWesQueryBillcountBillTypeCountDTO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryBillcountBillTypeCountDTO) SetBillTypeName(v string) *AlibabaAlihealthDrugWesQueryBillcountBillTypeCountDTO {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryBillcountBillTypeCountDTO) SetCount(v int64) *AlibabaAlihealthDrugWesQueryBillcountBillTypeCountDTO {
    s.Count = &v
    return s
}
