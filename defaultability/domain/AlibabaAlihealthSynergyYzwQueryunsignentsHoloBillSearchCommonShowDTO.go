package domain


type AlibabaAlihealthSynergyYzwQueryunsignentsHoloBillSearchCommonShowDTO struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQueryunsignentsHoloBillSearchCommonShowDTO) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwQueryunsignentsHoloBillSearchCommonShowDTO {
    s.RefEntId = &v
    return s
}
