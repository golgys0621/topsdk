package domain


type AlibabaAlihealthSynergySyContractDetailSyContractResourceModel struct {
    /*
        pdf在阿里云oss对象存储中的url     */
    PdfOssKeyUrl  *string `json:"pdf_oss_key_url,omitempty" `

    /*
        页数     */
    Page  *int64 `json:"page,omitempty" `

    /*
        id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        pdf名称     */
    ResourceName  *string `json:"resource_name,omitempty" `

    /*
        接收方签章     */
    ReceiveSealInfos  *AlibabaAlihealthSynergySyContractDetailResourceSealInfo `json:"receive_seal_infos,omitempty" `

    /*
        接受方签章     */
    ReceiveSealLimit  *AlibabaAlihealthSynergySyContractDetailResourceSealInfo `json:"receive_seal_limit,omitempty" `

    /*
        发送方签章信息     */
    SendSealInfos  *AlibabaAlihealthSynergySyContractDetailResourceSealInfo `json:"send_seal_infos,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractDetailSyContractResourceModel) SetPdfOssKeyUrl(v string) *AlibabaAlihealthSynergySyContractDetailSyContractResourceModel {
    s.PdfOssKeyUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractResourceModel) SetPage(v int64) *AlibabaAlihealthSynergySyContractDetailSyContractResourceModel {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractResourceModel) SetId(v int64) *AlibabaAlihealthSynergySyContractDetailSyContractResourceModel {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractResourceModel) SetResourceName(v string) *AlibabaAlihealthSynergySyContractDetailSyContractResourceModel {
    s.ResourceName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractResourceModel) SetReceiveSealInfos(v AlibabaAlihealthSynergySyContractDetailResourceSealInfo) *AlibabaAlihealthSynergySyContractDetailSyContractResourceModel {
    s.ReceiveSealInfos = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractResourceModel) SetReceiveSealLimit(v AlibabaAlihealthSynergySyContractDetailResourceSealInfo) *AlibabaAlihealthSynergySyContractDetailSyContractResourceModel {
    s.ReceiveSealLimit = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractResourceModel) SetSendSealInfos(v AlibabaAlihealthSynergySyContractDetailResourceSealInfo) *AlibabaAlihealthSynergySyContractDetailSyContractResourceModel {
    s.SendSealInfos = &v
    return s
}
