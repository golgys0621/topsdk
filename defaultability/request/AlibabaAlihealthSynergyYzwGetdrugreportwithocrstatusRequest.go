package request


type AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusRequest struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        任务ID     */
    Id  *int64 `json:"id" required:"true" `
}

func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusRequest) SetId(v int64) *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusRequest {
    s.Id = &v
    return s
}

func (req *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.Id != nil) {
        paramMap["id"] = *req.Id
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}