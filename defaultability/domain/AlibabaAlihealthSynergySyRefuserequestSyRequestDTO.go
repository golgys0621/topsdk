package domain


type AlibabaAlihealthSynergySyRefuserequestSyRequestDTO struct {
    /*
        企业（本企业id）     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        被索取ID     */
    Id  *int64 `json:"id,omitempty" `

    /*
        拒绝理由     */
    Reason  *string `json:"reason,omitempty" `

}

func (s *AlibabaAlihealthSynergySyRefuserequestSyRequestDTO) SetRefEntId(v string) *AlibabaAlihealthSynergySyRefuserequestSyRequestDTO {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyRefuserequestSyRequestDTO) SetId(v int64) *AlibabaAlihealthSynergySyRefuserequestSyRequestDTO {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthSynergySyRefuserequestSyRequestDTO) SetReason(v string) *AlibabaAlihealthSynergySyRefuserequestSyRequestDTO {
    s.Reason = &v
    return s
}
