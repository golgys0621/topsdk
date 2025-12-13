package request


type AlibabaAlihealthSynergyYzwDrugreportGetRequest struct {
    /*
        报告上传企业refENtId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        报告id     */
    ReportId  *int64 `json:"report_id" required:"true" `
}

func (s *AlibabaAlihealthSynergyYzwDrugreportGetRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwDrugreportGetRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetRequest) SetReportId(v int64) *AlibabaAlihealthSynergyYzwDrugreportGetRequest {
    s.ReportId = &v
    return s
}

func (req *AlibabaAlihealthSynergyYzwDrugreportGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.ReportId != nil) {
        paramMap["report_id"] = *req.ReportId
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergyYzwDrugreportGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}