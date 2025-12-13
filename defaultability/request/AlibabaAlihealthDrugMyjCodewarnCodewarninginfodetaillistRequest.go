package request


type AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistRequest struct {
    /*
        企业标识     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        任务流水编码（码预警任务列表中的ID）     */
    CodeFlowWarningId  *int64 `json:"code_flow_warning_id" required:"true" `
    /*
        页数     */
    Page  *int64 `json:"page" required:"true" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistRequest) SetCodeFlowWarningId(v int64) *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistRequest {
    s.CodeFlowWarningId = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistRequest) SetPage(v int64) *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistRequest) SetPageSize(v int64) *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.CodeFlowWarningId != nil) {
        paramMap["code_flow_warning_id"] = *req.CodeFlowWarningId
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}