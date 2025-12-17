package request


type AlibabaAlihealthSynergySySignatureDetailGetRequest struct {
    /*
        签章id     */
    SignatureId  *int64 `json:"signature_id" required:"true" `
    /*
        发送企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthSynergySySignatureDetailGetRequest) SetSignatureId(v int64) *AlibabaAlihealthSynergySySignatureDetailGetRequest {
    s.SignatureId = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySySignatureDetailGetRequest {
    s.RefEntId = &v
    return s
}

func (req *AlibabaAlihealthSynergySySignatureDetailGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SignatureId != nil) {
        paramMap["signature_id"] = *req.SignatureId
    }
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySySignatureDetailGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}