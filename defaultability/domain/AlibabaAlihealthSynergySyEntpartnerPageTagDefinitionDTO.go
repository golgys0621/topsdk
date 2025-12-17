package domain


type AlibabaAlihealthSynergySyEntpartnerPageTagDefinitionDTO struct {
    /*
        id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        name     */
    Name  *string `json:"name,omitempty" `

}

func (s *AlibabaAlihealthSynergySyEntpartnerPageTagDefinitionDTO) SetId(v int64) *AlibabaAlihealthSynergySyEntpartnerPageTagDefinitionDTO {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageTagDefinitionDTO) SetName(v string) *AlibabaAlihealthSynergySyEntpartnerPageTagDefinitionDTO {
    s.Name = &v
    return s
}
