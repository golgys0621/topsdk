package domain


type AlibabaAlihealthDrugLsydQueryUpoutbillcountBillTypeCountDTO struct {
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

func (s *AlibabaAlihealthDrugLsydQueryUpoutbillcountBillTypeCountDTO) SetBillType(v string) *AlibabaAlihealthDrugLsydQueryUpoutbillcountBillTypeCountDTO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryUpoutbillcountBillTypeCountDTO) SetBillTypeName(v string) *AlibabaAlihealthDrugLsydQueryUpoutbillcountBillTypeCountDTO {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryUpoutbillcountBillTypeCountDTO) SetCount(v int64) *AlibabaAlihealthDrugLsydQueryUpoutbillcountBillTypeCountDTO {
    s.Count = &v
    return s
}
