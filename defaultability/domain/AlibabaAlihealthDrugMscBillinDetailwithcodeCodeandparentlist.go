package domain


type AlibabaAlihealthDrugMscBillinDetailwithcodeCodeandparentlist struct {
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

func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeCodeandparentlist) SetCode(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeCodeandparentlist {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeCodeandparentlist) SetCodeLevel(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeCodeandparentlist {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeCodeandparentlist) SetParentCode(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeCodeandparentlist {
    s.ParentCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscBillinDetailwithcodeCodeandparentlist) SetHasParentCode(v string) *AlibabaAlihealthDrugMscBillinDetailwithcodeCodeandparentlist {
    s.HasParentCode = &v
    return s
}
