package domain


type AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse struct {
    /*
        资料id     */
    ResourceId  *int64 `json:"resource_id,omitempty" `

    /*
        资质分类编码     */
    ResourceTypeCode  *string `json:"resource_type_code,omitempty" `

    /*
        资质分类描述     */
    ResourceTypeDesc  *string `json:"resource_type_desc,omitempty" `

    /*
        资料文件的http地址，20分钟有效     */
    PdfUrl  *string `json:"pdf_url,omitempty" `

    /*
        签章后的资料http地址，20分钟有效     */
    SignaturePdfUrl  *string `json:"signature_pdf_url,omitempty" `

    /*
        资料签章状态-0:未签章,1:已签章     */
    SignatureStatus  *int64 `json:"signature_status,omitempty" `

    /*
        资料的页数     */
    Page  *int64 `json:"page,omitempty" `

    /*
        盖章成功后不返回。资料的印章信息集合，     */
    SealInfoList  *[]AlibabaAlihealthSynergySySignatureDetailGetSealInfo `json:"seal_info_list,omitempty" `

    /*
        资料名称     */
    ResourceName  *string `json:"resource_name,omitempty" `

}

func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse) SetResourceId(v int64) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse {
    s.ResourceId = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse) SetResourceTypeCode(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse {
    s.ResourceTypeCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse) SetResourceTypeDesc(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse {
    s.ResourceTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse) SetPdfUrl(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse {
    s.PdfUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse) SetSignaturePdfUrl(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse {
    s.SignaturePdfUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse) SetSignatureStatus(v int64) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse {
    s.SignatureStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse) SetPage(v int64) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse) SetSealInfoList(v []AlibabaAlihealthSynergySySignatureDetailGetSealInfo) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse {
    s.SealInfoList = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse) SetResourceName(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse {
    s.ResourceName = &v
    return s
}
