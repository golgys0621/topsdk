package request


type AlibabaAlihealthSynergySyContractRefuseRequest struct {
    /*
        本企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        合同id     */
    ContractId  *int64 `json:"contract_id" required:"true" `
    /*
        拒签理由     */
    Reason  *string `json:"reason" required:"true" `
}

func (s *AlibabaAlihealthSynergySyContractRefuseRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyContractRefuseRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractRefuseRequest) SetContractId(v int64) *AlibabaAlihealthSynergySyContractRefuseRequest {
    s.ContractId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractRefuseRequest) SetReason(v string) *AlibabaAlihealthSynergySyContractRefuseRequest {
    s.Reason = &v
    return s
}

func (req *AlibabaAlihealthSynergySyContractRefuseRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.ContractId != nil) {
        paramMap["contract_id"] = *req.ContractId
    }
    if(req.Reason != nil) {
        paramMap["reason"] = *req.Reason
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyContractRefuseRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}