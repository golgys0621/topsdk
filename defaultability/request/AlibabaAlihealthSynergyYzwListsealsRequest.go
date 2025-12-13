package request


type AlibabaAlihealthSynergyYzwListsealsRequest struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthSynergyYzwListsealsRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwListsealsRequest {
    s.RefEntId = &v
    return s
}

func (req *AlibabaAlihealthSynergyYzwListsealsRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergyYzwListsealsRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}