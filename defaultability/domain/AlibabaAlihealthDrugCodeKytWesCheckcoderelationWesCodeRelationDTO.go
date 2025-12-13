package domain


type AlibabaAlihealthDrugCodeKytWesCheckcoderelationWesCodeRelationDTO struct {
    /*
        存在上下级关系时返回下级码     */
    Code  *string `json:"code,omitempty" `

    /*
        存在上下级关系时返回上级码     */
    ParentCode  *string `json:"parent_code,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesCheckcoderelationWesCodeRelationDTO) SetCode(v string) *AlibabaAlihealthDrugCodeKytWesCheckcoderelationWesCodeRelationDTO {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesCheckcoderelationWesCodeRelationDTO) SetParentCode(v string) *AlibabaAlihealthDrugCodeKytWesCheckcoderelationWesCodeRelationDTO {
    s.ParentCode = &v
    return s
}
