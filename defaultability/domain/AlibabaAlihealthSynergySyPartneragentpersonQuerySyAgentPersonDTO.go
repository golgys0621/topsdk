package domain


type AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO struct {
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

    /*
        合作企业refEntId     */
    PartnerRefEntId  *string `json:"partner_ref_ent_id,omitempty" `

    /*
        合作企业entId     */
    PartnerEntId  *string `json:"partner_ent_id,omitempty" `

}

func (s *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO) SetUsername(v string) *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO {
    s.Username = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO) SetUserNo(v string) *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO {
    s.UserNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO) SetIdentityNo(v string) *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO {
    s.IdentityNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO) SetIdentityExpireDateStr(v string) *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO {
    s.IdentityExpireDateStr = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO) SetTel(v string) *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO {
    s.Tel = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO) SetAuthArea(v string) *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO {
    s.AuthArea = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO) SetAuthProduct(v string) *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO {
    s.AuthProduct = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO) SetExpireStatus(v string) *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO {
    s.ExpireStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO) SetArchiveStatus(v string) *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO {
    s.ArchiveStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO) SetId(v int64) *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO) SetPartnerRefEntId(v string) *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO {
    s.PartnerRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO) SetPartnerEntId(v string) *AlibabaAlihealthSynergySyPartneragentpersonQuerySyAgentPersonDTO {
    s.PartnerEntId = &v
    return s
}
