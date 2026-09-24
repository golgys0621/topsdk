package request


type AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistRequest struct {
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
    /*
        服务商id     */
    IsvRefEntId  *string `json:"isv_ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistRequest) SetCodeFlowWarningId(v int64) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistRequest {
    s.CodeFlowWarningId = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistRequest) SetPage(v int64) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistRequest) SetPageSize(v int64) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistRequest) SetIsvRefEntId(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistRequest {
    s.IsvRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistRequest) ToMap() map[string]interface{} {
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
    if(req.IsvRefEntId != nil) {
        paramMap["isv_ref_ent_id"] = *req.IsvRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}