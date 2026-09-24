package domain


type AlibabaAlihealthDrugYljgSearchbillDetailCodeInfoList struct {
    /*
        码     */
    Code  *string `json:"code,omitempty" `

    /*
        码层级     */
    CodeLevel  *string `json:"code_level,omitempty" `

}

func (s *AlibabaAlihealthDrugYljgSearchbillDetailCodeInfoList) SetCode(v string) *AlibabaAlihealthDrugYljgSearchbillDetailCodeInfoList {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailCodeInfoList) SetCodeLevel(v string) *AlibabaAlihealthDrugYljgSearchbillDetailCodeInfoList {
    s.CodeLevel = &v
    return s
}
