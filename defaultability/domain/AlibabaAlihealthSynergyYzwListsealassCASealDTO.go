package domain


type AlibabaAlihealthSynergyYzwListsealassCASealDTO struct {
    /*
        印章名称     */
    SealName  *string `json:"seal_name,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwListsealassCASealDTO) SetSealName(v string) *AlibabaAlihealthSynergyYzwListsealassCASealDTO {
    s.SealName = &v
    return s
}
