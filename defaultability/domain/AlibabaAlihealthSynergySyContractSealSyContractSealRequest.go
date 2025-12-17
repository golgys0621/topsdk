package domain


type AlibabaAlihealthSynergySyContractSealSyContractSealRequest struct {
    /*
        合同id     */
    ContractId  *int64 `json:"contract_id,omitempty" `

    /*
        本企业refEntId     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        印章信息     */
    SealInfos  *[]AlibabaAlihealthSynergySyContractSealSealInfo `json:"seal_infos,omitempty" `

    /*
        合同资料id     */
    ContractResourceId  *int64 `json:"contract_resource_id,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractSealSyContractSealRequest) SetContractId(v int64) *AlibabaAlihealthSynergySyContractSealSyContractSealRequest {
    s.ContractId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractSealSyContractSealRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyContractSealSyContractSealRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractSealSyContractSealRequest) SetSealInfos(v []AlibabaAlihealthSynergySyContractSealSealInfo) *AlibabaAlihealthSynergySyContractSealSyContractSealRequest {
    s.SealInfos = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractSealSyContractSealRequest) SetContractResourceId(v int64) *AlibabaAlihealthSynergySyContractSealSyContractSealRequest {
    s.ContractResourceId = &v
    return s
}
