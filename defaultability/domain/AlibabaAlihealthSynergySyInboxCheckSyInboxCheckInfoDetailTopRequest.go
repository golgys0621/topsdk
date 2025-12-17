package domain


type AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoDetailTopRequest struct {
    /*
        资料id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        是否通过     */
    IsPass  *bool `json:"is_pass,omitempty" `

}

func (s *AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoDetailTopRequest) SetId(v int64) *AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoDetailTopRequest {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoDetailTopRequest) SetIsPass(v bool) *AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoDetailTopRequest {
    s.IsPass = &v
    return s
}
