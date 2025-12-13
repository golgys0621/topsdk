package request


type AlibabaAlihealthEntYzwAuthGetdruginfoRequest struct {
    /*
        查询企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        授权企业refEntId     */
    AuthRefEntId  *string `json:"auth_ref_ent_id" required:"true" `
    /*
        药品上市许可持有人的refEntId     */
    AuthorizerRefEntId  *string `json:"authorizer_ref_ent_id,omitempty" required:"false" `
    /*
        药品通用名     */
    PhysicName  *string `json:"physic_name,omitempty" required:"false" `
    /*
        药品生产企业refEntId     */
    ProduceRefEntId  *string `json:"produce_ref_ent_id,omitempty" required:"false" `
    /*
        批准文号     */
    ApprovalLicenceNo  *string `json:"approval_licence_no,omitempty" required:"false" `
    /*
        每页数量     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthEntYzwAuthGetdruginfoRequest) SetRefEntId(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoRequest) SetAuthRefEntId(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoRequest {
    s.AuthRefEntId = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoRequest) SetAuthorizerRefEntId(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoRequest {
    s.AuthorizerRefEntId = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoRequest) SetPhysicName(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoRequest {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoRequest) SetProduceRefEntId(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoRequest {
    s.ProduceRefEntId = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoRequest) SetApprovalLicenceNo(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoRequest {
    s.ApprovalLicenceNo = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoRequest) SetPageSize(v int64) *AlibabaAlihealthEntYzwAuthGetdruginfoRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoRequest) SetPage(v int64) *AlibabaAlihealthEntYzwAuthGetdruginfoRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthEntYzwAuthGetdruginfoRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.AuthRefEntId != nil) {
        paramMap["auth_ref_ent_id"] = *req.AuthRefEntId
    }
    if(req.AuthorizerRefEntId != nil) {
        paramMap["authorizer_ref_ent_id"] = *req.AuthorizerRefEntId
    }
    if(req.PhysicName != nil) {
        paramMap["physic_name"] = *req.PhysicName
    }
    if(req.ProduceRefEntId != nil) {
        paramMap["produce_ref_ent_id"] = *req.ProduceRefEntId
    }
    if(req.ApprovalLicenceNo != nil) {
        paramMap["approval_licence_no"] = *req.ApprovalLicenceNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthEntYzwAuthGetdruginfoRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}