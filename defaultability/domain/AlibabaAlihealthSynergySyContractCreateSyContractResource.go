package domain


type AlibabaAlihealthSynergySyContractCreateSyContractResource struct {
    /*
        接收方印章限制     */
    LimitSealInfos  *[]AlibabaAlihealthSynergySyContractCreateSealInfo `json:"limit_seal_infos,omitempty" `

    /*
        资料名称     */
    ResourceName  *string `json:"resource_name,omitempty" `

    /*
        文件类型，目前只支持A4大小pdf     */
    FileType  *string `json:"file_type,omitempty" `

    /*
        文件url     */
    Url  *string `json:"url,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractCreateSyContractResource) SetLimitSealInfos(v []AlibabaAlihealthSynergySyContractCreateSealInfo) *AlibabaAlihealthSynergySyContractCreateSyContractResource {
    s.LimitSealInfos = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateSyContractResource) SetResourceName(v string) *AlibabaAlihealthSynergySyContractCreateSyContractResource {
    s.ResourceName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateSyContractResource) SetFileType(v string) *AlibabaAlihealthSynergySyContractCreateSyContractResource {
    s.FileType = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateSyContractResource) SetUrl(v string) *AlibabaAlihealthSynergySyContractCreateSyContractResource {
    s.Url = &v
    return s
}
