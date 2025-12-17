package domain


type AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse struct {
    /*
        资料id     */
    ResourceId  *int64 `json:"resource_id,omitempty" `

    /*
        资料类型code     */
    ResourceTypeCode  *string `json:"resource_type_code,omitempty" `

    /*
        资料类型     */
    ResourceTypeDesc  *string `json:"resource_type_desc,omitempty" `

    /*
        资料URL     */
    SignatureOssKeyUrl  *string `json:"signature_oss_key_url,omitempty" `

    /*
        资料页数     */
    Page  *int64 `json:"page,omitempty" `

    /*
        资料查收状态，0不通过1通过     */
    ApproveStatus  *int64 `json:"approve_status,omitempty" `

    /*
        资料过期时间     */
    ExpireDate  *string `json:"expire_date,omitempty" `

    /*
        资料名称     */
    ResourceName  *string `json:"resource_name,omitempty" `

}

func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse) SetResourceId(v int64) *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse {
    s.ResourceId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse) SetResourceTypeCode(v string) *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse {
    s.ResourceTypeCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse) SetResourceTypeDesc(v string) *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse {
    s.ResourceTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse) SetSignatureOssKeyUrl(v string) *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse {
    s.SignatureOssKeyUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse) SetPage(v int64) *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse) SetApproveStatus(v int64) *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse {
    s.ApproveStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse) SetExpireDate(v string) *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse {
    s.ExpireDate = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse) SetResourceName(v string) *AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse {
    s.ResourceName = &v
    return s
}
