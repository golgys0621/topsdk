package domain


type AlibabaAlihealthDrugLsydQueryBillcountBillTypeCountDTO struct {
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

func (s *AlibabaAlihealthDrugLsydQueryBillcountBillTypeCountDTO) SetBillType(v string) *AlibabaAlihealthDrugLsydQueryBillcountBillTypeCountDTO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryBillcountBillTypeCountDTO) SetBillTypeName(v string) *AlibabaAlihealthDrugLsydQueryBillcountBillTypeCountDTO {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryBillcountBillTypeCountDTO) SetCount(v int64) *AlibabaAlihealthDrugLsydQueryBillcountBillTypeCountDTO {
    s.Count = &v
    return s
}
