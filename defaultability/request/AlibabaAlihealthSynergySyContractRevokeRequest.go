package request


type AlibabaAlihealthSynergySyContractRevokeRequest struct {
    /*
        本企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        发送方合同id     */
    Id  *int64 `json:"id" required:"true" `
}

func (s *AlibabaAlihealthSynergySyContractRevokeRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyContractRevokeRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractRevokeRequest) SetId(v int64) *AlibabaAlihealthSynergySyContractRevokeRequest {
    s.Id = &v
    return s
}

func (req *AlibabaAlihealthSynergySyContractRevokeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.Id != nil) {
        paramMap["id"] = *req.Id
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyContractRevokeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}