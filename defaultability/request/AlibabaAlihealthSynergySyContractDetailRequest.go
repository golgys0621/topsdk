package request


type AlibabaAlihealthSynergySyContractDetailRequest struct {
    /*
        本企业refentId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        合同id     */
    Id  *int64 `json:"id" required:"true" `
}

func (s *AlibabaAlihealthSynergySyContractDetailRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyContractDetailRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailRequest) SetId(v int64) *AlibabaAlihealthSynergySyContractDetailRequest {
    s.Id = &v
    return s
}

func (req *AlibabaAlihealthSynergySyContractDetailRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.Id != nil) {
        paramMap["id"] = *req.Id
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyContractDetailRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}