package request


type AlibabaAlihealthSynergySyListrequestdetailRequest struct {
    /*
        企业的refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        发出索取单父表的id     */
    RequestId  *int64 `json:"request_id" required:"true" `
}

func (s *AlibabaAlihealthSynergySyListrequestdetailRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyListrequestdetailRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListrequestdetailRequest) SetRequestId(v int64) *AlibabaAlihealthSynergySyListrequestdetailRequest {
    s.RequestId = &v
    return s
}

func (req *AlibabaAlihealthSynergySyListrequestdetailRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.RequestId != nil) {
        paramMap["request_id"] = *req.RequestId
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyListrequestdetailRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}