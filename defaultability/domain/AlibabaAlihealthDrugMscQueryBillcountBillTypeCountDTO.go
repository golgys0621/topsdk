package domain


type AlibabaAlihealthDrugMscQueryBillcountBillTypeCountDTO struct {
    /*
        单据类型枚举     */
    BillType  *int64 `json:"bill_type,omitempty" `

    /*
        单据类型     */
    BillTypeName  *string `json:"bill_type_name,omitempty" `

    /*
        单据数量     */
    Count  *int64 `json:"count,omitempty" `

}

func (s *AlibabaAlihealthDrugMscQueryBillcountBillTypeCountDTO) SetBillType(v int64) *AlibabaAlihealthDrugMscQueryBillcountBillTypeCountDTO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscQueryBillcountBillTypeCountDTO) SetBillTypeName(v string) *AlibabaAlihealthDrugMscQueryBillcountBillTypeCountDTO {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscQueryBillcountBillTypeCountDTO) SetCount(v int64) *AlibabaAlihealthDrugMscQueryBillcountBillTypeCountDTO {
    s.Count = &v
    return s
}
