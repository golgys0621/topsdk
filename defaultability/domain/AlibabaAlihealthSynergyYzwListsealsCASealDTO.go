package domain


type AlibabaAlihealthSynergyYzwListsealsCASealDTO struct {
    /*
        印章名称     */
    SealName  *string `json:"seal_name,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwListsealsCASealDTO) SetSealName(v string) *AlibabaAlihealthSynergyYzwListsealsCASealDTO {
    s.SealName = &v
    return s
}
