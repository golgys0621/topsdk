package request


type AlibabaAlihealthSynergyYzwQueryreportbycodeRequest struct {
    /*
        本企业     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        追溯码     */
    Code  *string `json:"code" required:"true" `
}

func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeRequest) SetCode(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeRequest {
    s.Code = &v
    return s
}

func (req *AlibabaAlihealthSynergyYzwQueryreportbycodeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergyYzwQueryreportbycodeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}