package request


type AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest struct {
    /*
        报告上传企业refENtId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        报告创建开始时间     */
    BeginTime  *string `json:"begin_time" required:"true" `
    /*
        报告创建结束时间     */
    EndTime  *string `json:"end_time" required:"true" `
    /*
        签章状态:1是未盖章,2是已盖章     */
    Status  *int64 `json:"status,omitempty" required:"false" `
    /*
        页码 defalutValue��1    */
    Page  *int64 `json:"page" required:"true" `
    /*
        页的大小 defalutValue��20    */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest) SetBeginTime(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest {
    s.BeginTime = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest) SetEndTime(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest {
    s.EndTime = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest) SetStatus(v int64) *AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest) SetPage(v int64) *AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest) SetPageSize(v int64) *AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest) ToMap() map[string]interface{} {
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
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}