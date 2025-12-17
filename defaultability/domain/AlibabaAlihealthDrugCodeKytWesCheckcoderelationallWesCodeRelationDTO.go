package domain


type AlibabaAlihealthDrugCodeKytWesCheckcoderelationallWesCodeRelationDTO struct {
    /*
        码     */
    Code  *string `json:"code,omitempty" `

    /*
        父码     */
    ParentCode  *string `json:"parent_code,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallWesCodeRelationDTO) SetCode(v string) *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallWesCodeRelationDTO {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallWesCodeRelationDTO) SetParentCode(v string) *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallWesCodeRelationDTO {
    s.ParentCode = &v
    return s
}
