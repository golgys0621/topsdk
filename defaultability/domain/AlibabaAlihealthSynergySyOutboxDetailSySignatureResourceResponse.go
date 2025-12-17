package domain


type AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse struct {
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
        页数     */
    Page  *int64 `json:"page,omitempty" `

    /*
        资料查收状态，0不通过1通过     */
    ApproveStatus  *int64 `json:"approve_status,omitempty" `

    /*
        外部自定义Id     */
    OutId  *string `json:"out_id,omitempty" `

    /*
        资料名称     */
    ResourceName  *string `json:"resource_name,omitempty" `

}

func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse) SetResourceId(v int64) *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse {
    s.ResourceId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse) SetResourceTypeCode(v string) *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse {
    s.ResourceTypeCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse) SetResourceTypeDesc(v string) *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse {
    s.ResourceTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse) SetSignatureOssKeyUrl(v string) *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse {
    s.SignatureOssKeyUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse) SetPage(v int64) *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse) SetApproveStatus(v int64) *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse {
    s.ApproveStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse) SetOutId(v string) *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse {
    s.OutId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse) SetResourceName(v string) *AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse {
    s.ResourceName = &v
    return s
}
