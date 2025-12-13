package request


type AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusRequest struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        任务ID     */
    Id  *int64 `json:"id" required:"true" `
}

func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusRequest) SetId(v int64) *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusRequest {
    s.Id = &v
    return s
}

func (req *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.Id != nil) {
        paramMap["id"] = *req.Id
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}