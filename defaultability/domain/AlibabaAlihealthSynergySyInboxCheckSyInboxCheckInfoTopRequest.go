package domain


type AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoTopRequest struct {
    /*
        委托人资料查收结果     */
    CheckInfoDetails  *[]AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoDetailTopRequest `json:"check_info_details,omitempty" `

    /*
        委托人id     */
    Id  *int64 `json:"id,omitempty" `

}

func (s *AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoTopRequest) SetCheckInfoDetails(v []AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoDetailTopRequest) *AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoTopRequest {
    s.CheckInfoDetails = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoTopRequest) SetId(v int64) *AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoTopRequest {
    s.Id = &v
    return s
}
