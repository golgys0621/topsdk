package request


type AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest struct {
    /*
        调用企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        操作开始时间     */
    BeginTime  *string `json:"begin_time" required:"true" `
    /*
        操作结束时间     */
    EndTime  *string `json:"end_time" required:"true" `
    /*
        批次号     */
    BatchNo  *string `json:"batch_no,omitempty" required:"false" `
    /*
        药品id     */
    DrugId  *string `json:"drug_id,omitempty" required:"false" `
    /*
        页码 defalutValue��1    */
    Page  *int64 `json:"page" required:"true" `
    /*
        页的大小 defalutValue��50    */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest) SetBeginTime(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest {
    s.BeginTime = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest) SetEndTime(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest {
    s.EndTime = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest {
    s.BatchNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest) SetDrugId(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest {
    s.DrugId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest) SetPage(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest) SetPageSize(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.BeginTime != nil) {
        paramMap["begin_time"] = *req.BeginTime
    }
    if(req.EndTime != nil) {
        paramMap["end_time"] = *req.EndTime
    }
    if(req.BatchNo != nil) {
        paramMap["batch_no"] = *req.BatchNo
    }
    if(req.DrugId != nil) {
        paramMap["drug_id"] = *req.DrugId
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}