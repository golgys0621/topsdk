package domain


type AlibabaAlihealthSynergySyContractDetailSyContractDetail struct {
    /*
        合同信息     */
    SyContractModel  *AlibabaAlihealthSynergySyContractDetailSyContractModel `json:"sy_contract_model,omitempty" `

    /*
        材料信息     */
    SyContractResourceModel  *[]AlibabaAlihealthSynergySyContractDetailSyContractResourceModel `json:"sy_contract_resource_model,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractDetailSyContractDetail) SetSyContractModel(v AlibabaAlihealthSynergySyContractDetailSyContractModel) *AlibabaAlihealthSynergySyContractDetailSyContractDetail {
    s.SyContractModel = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractDetail) SetSyContractResourceModel(v []AlibabaAlihealthSynergySyContractDetailSyContractResourceModel) *AlibabaAlihealthSynergySyContractDetailSyContractDetail {
    s.SyContractResourceModel = &v
    return s
}
