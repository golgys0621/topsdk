package domain


type AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist struct {
    /*
        追溯码     */
    Code  *string `json:"code,omitempty" `

    /*
        码级别     */
    CodeLevel  *string `json:"code_level,omitempty" `

    /*
        父码     */
    ParentCode  *string `json:"parent_code,omitempty" `

    /*
        是否有父码     */
    HasParentCode  *string `json:"has_parent_code,omitempty" `

}

func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist) SetCode(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist) SetCodeLevel(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist) SetParentCode(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist {
    s.ParentCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist) SetHasParentCode(v string) *AlibabaAlihealthDrugMscBilloutDetailwithcodesCodeandparentlist {
    s.HasParentCode = &v
    return s
}
