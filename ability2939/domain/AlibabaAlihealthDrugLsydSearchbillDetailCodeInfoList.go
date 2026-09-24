package domain


type AlibabaAlihealthDrugLsydSearchbillDetailCodeInfoList struct {
    /*
        码     */
    Code  *string `json:"code,omitempty" `

    /*
        码层级     */
    CodeLevel  *string `json:"code_level,omitempty" `

}

func (s *AlibabaAlihealthDrugLsydSearchbillDetailCodeInfoList) SetCode(v string) *AlibabaAlihealthDrugLsydSearchbillDetailCodeInfoList {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailCodeInfoList) SetCodeLevel(v string) *AlibabaAlihealthDrugLsydSearchbillDetailCodeInfoList {
    s.CodeLevel = &v
    return s
}
