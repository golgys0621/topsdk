package request


type AlibabaAlihealthSynergyYzwListsealassRequest struct {
    /*
        本企业     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        委托企业     */
    AssRefEntId  *string `json:"ass_ref_ent_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergyYzwListsealassRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwListsealassRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwListsealassRequest) SetAssRefEntId(v string) *AlibabaAlihealthSynergyYzwListsealassRequest {
    s.AssRefEntId = &v
    return s
}

func (req *AlibabaAlihealthSynergyYzwListsealassRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.AssRefEntId != nil) {
        paramMap["ass_ref_ent_id"] = *req.AssRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergyYzwListsealassRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}