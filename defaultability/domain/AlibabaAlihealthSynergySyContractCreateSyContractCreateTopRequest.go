package domain


type AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest struct {
    /*
        合同接收方refEntId     */
    ReceiveRefEntId  *string `json:"receive_ref_ent_id,omitempty" `

    /*
        0质量保障协议/1销售合同/2采购合同     */
    ContractType  *int64 `json:"contract_type,omitempty" `

    /*
        合同编号     */
    ContractNo  *string `json:"contract_no,omitempty" `

    /*
        资料列表     */
    Resources  *[]AlibabaAlihealthSynergySyContractCreateSyContractResource `json:"resources,omitempty" `

    /*
        合同名称     */
    ContractName  *string `json:"contract_name,omitempty" `

    /*
        过期时间     */
    ExpireDate  *string `json:"expire_date,omitempty" `

    /*
        本企业refentId     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        限制对方盖章，0不限制、1限制印章类型、2限制印章类型和位置     */
    ReceiverSealLimitType  *int64 `json:"receiver_seal_limit_type,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest) SetReceiveRefEntId(v string) *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest {
    s.ReceiveRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest) SetContractType(v int64) *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest {
    s.ContractType = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest) SetContractNo(v string) *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest {
    s.ContractNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest) SetResources(v []AlibabaAlihealthSynergySyContractCreateSyContractResource) *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest {
    s.Resources = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest) SetContractName(v string) *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest {
    s.ContractName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest) SetExpireDate(v string) *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest {
    s.ExpireDate = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest) SetReceiverSealLimitType(v int64) *AlibabaAlihealthSynergySyContractCreateSyContractCreateTopRequest {
    s.ReceiverSealLimitType = &v
    return s
}
