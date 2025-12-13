package domain


type AlibabaAlihealthDrugYljgQueryBillcountBillTypeCountDTO struct {
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

func (s *AlibabaAlihealthDrugYljgQueryBillcountBillTypeCountDTO) SetBillType(v string) *AlibabaAlihealthDrugYljgQueryBillcountBillTypeCountDTO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryBillcountBillTypeCountDTO) SetBillTypeName(v string) *AlibabaAlihealthDrugYljgQueryBillcountBillTypeCountDTO {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryBillcountBillTypeCountDTO) SetCount(v int64) *AlibabaAlihealthDrugYljgQueryBillcountBillTypeCountDTO {
    s.Count = &v
    return s
}
