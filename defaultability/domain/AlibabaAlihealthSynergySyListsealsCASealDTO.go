package domain


type AlibabaAlihealthSynergySyListsealsCASealDTO struct {
    /*
        印章名称     */
    SealName  *string `json:"seal_name,omitempty" `

    /*
        印章类型     */
    SealType  *string `json:"seal_type,omitempty" `

    /*
        印章大小单位厘米     */
    SealDesc  *string `json:"seal_desc,omitempty" `

}

func (s *AlibabaAlihealthSynergySyListsealsCASealDTO) SetSealName(v string) *AlibabaAlihealthSynergySyListsealsCASealDTO {
    s.SealName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListsealsCASealDTO) SetSealType(v string) *AlibabaAlihealthSynergySyListsealsCASealDTO {
    s.SealType = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListsealsCASealDTO) SetSealDesc(v string) *AlibabaAlihealthSynergySyListsealsCASealDTO {
    s.SealDesc = &v
    return s
}
