package request


type AlibabaAlihealthSynergyYzwDeletedrugreportRequest struct {
    /*
        报告上传企业     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        报告id     */
    ReportId  *string `json:"report_id,omitempty" required:"false" `
    /*
        药品id     */
    DrugId  *string `json:"drug_id,omitempty" required:"false" `
    /*
        批次号     */
    BatchNo  *string `json:"batch_no,omitempty" required:"false" `
    /*
        用户id     */
    UserId  *string `json:"user_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergyYzwDeletedrugreportRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwDeletedrugreportRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDeletedrugreportRequest) SetReportId(v string) *AlibabaAlihealthSynergyYzwDeletedrugreportRequest {
    s.ReportId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDeletedrugreportRequest) SetDrugId(v string) *AlibabaAlihealthSynergyYzwDeletedrugreportRequest {
    s.DrugId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDeletedrugreportRequest) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwDeletedrugreportRequest {
    s.BatchNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDeletedrugreportRequest) SetUserId(v string) *AlibabaAlihealthSynergyYzwDeletedrugreportRequest {
    s.UserId = &v
    return s
}

func (req *AlibabaAlihealthSynergyYzwDeletedrugreportRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.ReportId != nil) {
        paramMap["report_id"] = *req.ReportId
    }
    if(req.DrugId != nil) {
        paramMap["drug_id"] = *req.DrugId
    }
    if(req.BatchNo != nil) {
        paramMap["batch_no"] = *req.BatchNo
    }
    if(req.UserId != nil) {
        paramMap["user_id"] = *req.UserId
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergyYzwDeletedrugreportRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}