package request


type AlibabaAlihealthSynergySyListrequestRequest struct {
    /*
        企业的refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        索取人     */
    RequestPerson  *string `json:"request_person,omitempty" required:"false" `
    /*
        0企业资料1产品资料     */
    Type  *string `json:"type,omitempty" required:"false" `
    /*
        被索取的企业ref_ent_id     */
    ReceiveRefEntId  *string `json:"receive_ref_ent_id,omitempty" required:"false" `
    /*
        索取状态     */
    RequestStatus  *string `json:"request_status,omitempty" required:"false" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthSynergySyListrequestRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyListrequestRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListrequestRequest) SetRequestPerson(v string) *AlibabaAlihealthSynergySyListrequestRequest {
    s.RequestPerson = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListrequestRequest) SetType(v string) *AlibabaAlihealthSynergySyListrequestRequest {
    s.Type = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListrequestRequest) SetReceiveRefEntId(v string) *AlibabaAlihealthSynergySyListrequestRequest {
    s.ReceiveRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListrequestRequest) SetRequestStatus(v string) *AlibabaAlihealthSynergySyListrequestRequest {
    s.RequestStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListrequestRequest) SetPage(v int64) *AlibabaAlihealthSynergySyListrequestRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListrequestRequest) SetPageSize(v int64) *AlibabaAlihealthSynergySyListrequestRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthSynergySyListrequestRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.RequestPerson != nil) {
        paramMap["request_person"] = *req.RequestPerson
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.ReceiveRefEntId != nil) {
        paramMap["receive_ref_ent_id"] = *req.ReceiveRefEntId
    }
    if(req.RequestStatus != nil) {
        paramMap["request_status"] = *req.RequestStatus
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyListrequestRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}