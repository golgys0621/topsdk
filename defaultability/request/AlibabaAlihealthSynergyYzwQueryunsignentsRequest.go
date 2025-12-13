package request


type AlibabaAlihealthSynergyYzwQueryunsignentsRequest struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthSynergyYzwQueryunsignentsRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwQueryunsignentsRequest {
    s.RefEntId = &v
    return s
}

func (req *AlibabaAlihealthSynergyYzwQueryunsignentsRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergyYzwQueryunsignentsRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}