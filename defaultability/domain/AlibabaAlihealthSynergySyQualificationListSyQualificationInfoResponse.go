package domain


type AlibabaAlihealthSynergySyQualificationListSyQualificationInfoResponse struct {
    /*
        资质集合     */
    SyQualificationInfoList  *[]AlibabaAlihealthSynergySyQualificationListSyQualificationInfoList `json:"sy_qualification_info_list,omitempty" `

    /*
        资料类型:ent:企业资料,product:产品资料,agentUser:委托人资料     */
    Type  *string `json:"type,omitempty" `

}

func (s *AlibabaAlihealthSynergySyQualificationListSyQualificationInfoResponse) SetSyQualificationInfoList(v []AlibabaAlihealthSynergySyQualificationListSyQualificationInfoList) *AlibabaAlihealthSynergySyQualificationListSyQualificationInfoResponse {
    s.SyQualificationInfoList = &v
    return s
}
func (s *AlibabaAlihealthSynergySyQualificationListSyQualificationInfoResponse) SetType(v string) *AlibabaAlihealthSynergySyQualificationListSyQualificationInfoResponse {
    s.Type = &v
    return s
}
