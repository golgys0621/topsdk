package domain


type AlibabaAlihealthDrugYljgQueryUpoutbillcountBillTypeCountDTO struct {
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

func (s *AlibabaAlihealthDrugYljgQueryUpoutbillcountBillTypeCountDTO) SetBillType(v string) *AlibabaAlihealthDrugYljgQueryUpoutbillcountBillTypeCountDTO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryUpoutbillcountBillTypeCountDTO) SetBillTypeName(v string) *AlibabaAlihealthDrugYljgQueryUpoutbillcountBillTypeCountDTO {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryUpoutbillcountBillTypeCountDTO) SetCount(v int64) *AlibabaAlihealthDrugYljgQueryUpoutbillcountBillTypeCountDTO {
    s.Count = &v
    return s
}
