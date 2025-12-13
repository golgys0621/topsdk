package request


type AlibabaAlihealthDrugMscListauthsRequest struct {
    /*
        企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        企业名称     */
    EntName  *string `json:"ent_name,omitempty" required:"false" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthDrugMscListauthsRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscListauthsRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsRequest) SetEntName(v string) *AlibabaAlihealthDrugMscListauthsRequest {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsRequest) SetPage(v int64) *AlibabaAlihealthDrugMscListauthsRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListauthsRequest) SetPageSize(v int64) *AlibabaAlihealthDrugMscListauthsRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthDrugMscListauthsRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.EntName != nil) {
        paramMap["ent_name"] = *req.EntName
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMscListauthsRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}