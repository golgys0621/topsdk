package domain


type AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail struct {
    /*
        签章id     */
    SignatureId  *int64 `json:"signature_id,omitempty" `

    /*
        创建时间     */
    GmtCreate  *string `json:"gmt_create,omitempty" `

    /*
        接收资料的企业集合     */
    ReceiveEntNameList  *[]string `json:"receive_ent_name_list,omitempty" `

    /*
        签章资料类型：企业资料，产品资料     */
    TypeDesc  *string `json:"type_desc,omitempty" `

    /*
        创建签章的操作人     */
    SubmitPerson  *string `json:"submit_person,omitempty" `

    /*
        签章状态-200:待提交,202:待签章,204:签章中,205:待发送,206:发送中,207:已发送,208:签章失败     */
    SignatureStatus  *int64 `json:"signature_status,omitempty" `

    /*
        状态描述     */
    SignatureStatusDesc  *string `json:"signature_status_desc,omitempty" `

    /*
        签章时间     */
    SignatureFinishTime  *string `json:"signature_finish_time,omitempty" `

}

func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail) SetSignatureId(v int64) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail {
    s.SignatureId = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail) SetGmtCreate(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail {
    s.GmtCreate = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail) SetReceiveEntNameList(v []string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail {
    s.ReceiveEntNameList = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail) SetTypeDesc(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail {
    s.TypeDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail) SetSubmitPerson(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail {
    s.SubmitPerson = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail) SetSignatureStatus(v int64) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail {
    s.SignatureStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail) SetSignatureStatusDesc(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail {
    s.SignatureStatusDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail) SetSignatureFinishTime(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail {
    s.SignatureFinishTime = &v
    return s
}
