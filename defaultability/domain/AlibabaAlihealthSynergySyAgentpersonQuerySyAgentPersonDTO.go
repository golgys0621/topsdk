package domain


type AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO struct {
    /*
        委托人姓名     */
    Username  *string `json:"username,omitempty" `

    /*
        委托人编号     */
    UserNo  *string `json:"user_no,omitempty" `

    /*
        身份证号     */
    IdentityNo  *string `json:"identity_no,omitempty" `

    /*
        身份证有效期至     */
    IdentityExpireDateStr  *string `json:"identity_expire_date_str,omitempty" `

    /*
        联系电话     */
    Tel  *string `json:"tel,omitempty" `

    /*
        授权地区     */
    AuthArea  *string `json:"auth_area,omitempty" `

    /*
        授权品种     */
    AuthProduct  *string `json:"auth_product,omitempty" `

    /*
        过期状态:0未过期1已过期     */
    ExpireStatus  *string `json:"expire_status,omitempty" `

    /*
        归档状态:0未归档1已归档     */
    ArchiveStatus  *string `json:"archive_status,omitempty" `

    /*
        委托人id     */
    Id  *int64 `json:"id,omitempty" `

}

func (s *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO) SetUsername(v string) *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO {
    s.Username = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO) SetUserNo(v string) *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO {
    s.UserNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO) SetIdentityNo(v string) *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO {
    s.IdentityNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO) SetIdentityExpireDateStr(v string) *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO {
    s.IdentityExpireDateStr = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO) SetTel(v string) *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO {
    s.Tel = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO) SetAuthArea(v string) *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO {
    s.AuthArea = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO) SetAuthProduct(v string) *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO {
    s.AuthProduct = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO) SetExpireStatus(v string) *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO {
    s.ExpireStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO) SetArchiveStatus(v string) *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO {
    s.ArchiveStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO) SetId(v int64) *AlibabaAlihealthSynergySyAgentpersonQuerySyAgentPersonDTO {
    s.Id = &v
    return s
}
