package request


type AlibabaAlihealthSynergySyPartneragentpersonQueryRequest struct {
    /*
        refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        委托人姓名     */
    Username  *string `json:"username,omitempty" required:"false" `
    /*
        委托人编号     */
    UserNo  *string `json:"user_no,omitempty" required:"false" `
    /*
        过期状态:0未过期1已过期     */
    ExpireStatus  *string `json:"expire_status,omitempty" required:"false" `
    /*
        归档状态:0未归档1已归档     */
    ArchiveStatus  *string `json:"archive_status,omitempty" required:"false" `
    /*
        页码 defalutValue��1    */
    Page  *int64 `json:"page" required:"true" `
    /*
        页的大小 defalutValue��10    */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        合作企业的refEntId     */
    PartnerRefEntId  *string `json:"partner_ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest) SetUsername(v string) *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest {
    s.Username = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest) SetUserNo(v string) *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest {
    s.UserNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest) SetExpireStatus(v string) *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest {
    s.ExpireStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest) SetArchiveStatus(v string) *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest {
    s.ArchiveStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest) SetPage(v int64) *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest) SetPageSize(v int64) *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest) SetPartnerRefEntId(v string) *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest {
    s.PartnerRefEntId = &v
    return s
}

func (req *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.Username != nil) {
        paramMap["username"] = *req.Username
    }
    if(req.UserNo != nil) {
        paramMap["user_no"] = *req.UserNo
    }
    if(req.ExpireStatus != nil) {
        paramMap["expire_status"] = *req.ExpireStatus
    }
    if(req.ArchiveStatus != nil) {
        paramMap["archive_status"] = *req.ArchiveStatus
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PartnerRefEntId != nil) {
        paramMap["partner_ref_ent_id"] = *req.PartnerRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyPartneragentpersonQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}