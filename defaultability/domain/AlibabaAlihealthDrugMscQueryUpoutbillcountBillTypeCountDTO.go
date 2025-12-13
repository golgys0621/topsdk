package domain


type AlibabaAlihealthDrugMscQueryUpoutbillcountBillTypeCountDTO struct {
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

func (s *AlibabaAlihealthDrugMscQueryUpoutbillcountBillTypeCountDTO) SetBillType(v int64) *AlibabaAlihealthDrugMscQueryUpoutbillcountBillTypeCountDTO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscQueryUpoutbillcountBillTypeCountDTO) SetBillTypeName(v string) *AlibabaAlihealthDrugMscQueryUpoutbillcountBillTypeCountDTO {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscQueryUpoutbillcountBillTypeCountDTO) SetCount(v int64) *AlibabaAlihealthDrugMscQueryUpoutbillcountBillTypeCountDTO {
    s.Count = &v
    return s
}
