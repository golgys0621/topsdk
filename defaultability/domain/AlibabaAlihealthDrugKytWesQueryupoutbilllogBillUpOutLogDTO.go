package domain


type AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO struct {
    /*
        货主企业     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        单据ID     */
    BillId  *int64 `json:"bill_id,omitempty" `

    /*
        单据类型     */
    BillType  *int64 `json:"bill_type,omitempty" `

    /*
        单据类型描述     */
    BillTypeDesc  *string `json:"bill_type_desc,omitempty" `

    /*
        单据号     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        发货企业     */
    FromRefEntId  *string `json:"from_ref_ent_id,omitempty" `

    /*
        收货企业     */
    ToRefEntId  *string `json:"to_ref_ent_id,omitempty" `

    /*
        草坪做类型     */
    OperDate  *string `json:"oper_date,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO) SetBillId(v int64) *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO {
    s.BillId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO) SetBillType(v int64) *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO) SetBillTypeDesc(v string) *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO {
    s.BillTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO) SetBillCode(v string) *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO) SetFromRefEntId(v string) *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO {
    s.FromRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO) SetToRefEntId(v string) *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO {
    s.ToRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO) SetOperDate(v string) *AlibabaAlihealthDrugKytWesQueryupoutbilllogBillUpOutLogDTO {
    s.OperDate = &v
    return s
}
