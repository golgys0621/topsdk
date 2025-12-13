package request


type AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest struct {
    /*
        报告上传企业refENtId     */
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
        报告创建开始时间     */
    BeginTime  *string `json:"begin_time,omitempty" required:"false" `
    /*
        报告创建结束时间     */
    EndTime  *string `json:"end_time,omitempty" required:"false" `
    /*
        页码 defalutValue��1    */
    Page  *int64 `json:"page" required:"true" `
    /*
        页的大小 defalutValue��20    */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest) SetReportId(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest {
    s.ReportId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest) SetDrugId(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest {
    s.DrugId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest {
    s.BatchNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest) SetBeginTime(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest {
    s.BeginTime = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest) SetEndTime(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest {
    s.EndTime = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest) SetPage(v int64) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest) SetPageSize(v int64) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest) ToMap() map[string]interface{} {
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
    if(req.BeginTime != nil) {
        paramMap["begin_time"] = *req.BeginTime
    }
    if(req.EndTime != nil) {
        paramMap["end_time"] = *req.EndTime
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}