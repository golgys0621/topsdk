package request


type AlibabaAlihealthSynergySyEntpartnerPageRequest struct {
    /*
        当前登录企业的refEntId     */
    LoginRefEntId  *string `json:"login_ref_ent_id" required:"true" `
    /*
        合作企业的refEntId     */
    PartnerRefEntId  *string `json:"partner_ref_ent_id,omitempty" required:"false" `
    /*
        过期状态0未过期1已过期     */
    ExpireStatus  *string `json:"expire_status,omitempty" required:"false" `
    /*
        归档状态0未归档1已归档     */
    ArchiveStatus  *string `json:"archive_status,omitempty" required:"false" `
    /*
        使用状态:0曾用信息1使用中     */
    UseStatus  *string `json:"use_status,omitempty" required:"false" `
    /*
        标签ID列表     */
    TagIds  *string `json:"tag_ids,omitempty" required:"false" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthSynergySyEntpartnerPageRequest) SetLoginRefEntId(v string) *AlibabaAlihealthSynergySyEntpartnerPageRequest {
    s.LoginRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageRequest) SetPartnerRefEntId(v string) *AlibabaAlihealthSynergySyEntpartnerPageRequest {
    s.PartnerRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageRequest) SetExpireStatus(v string) *AlibabaAlihealthSynergySyEntpartnerPageRequest {
    s.ExpireStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageRequest) SetArchiveStatus(v string) *AlibabaAlihealthSynergySyEntpartnerPageRequest {
    s.ArchiveStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageRequest) SetUseStatus(v string) *AlibabaAlihealthSynergySyEntpartnerPageRequest {
    s.UseStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageRequest) SetTagIds(v string) *AlibabaAlihealthSynergySyEntpartnerPageRequest {
    s.TagIds = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageRequest) SetPage(v int64) *AlibabaAlihealthSynergySyEntpartnerPageRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerPageRequest) SetPageSize(v int64) *AlibabaAlihealthSynergySyEntpartnerPageRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthSynergySyEntpartnerPageRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.LoginRefEntId != nil) {
        paramMap["login_ref_ent_id"] = *req.LoginRefEntId
    }
    if(req.PartnerRefEntId != nil) {
        paramMap["partner_ref_ent_id"] = *req.PartnerRefEntId
    }
    if(req.ExpireStatus != nil) {
        paramMap["expire_status"] = *req.ExpireStatus
    }
    if(req.ArchiveStatus != nil) {
        paramMap["archive_status"] = *req.ArchiveStatus
    }
    if(req.UseStatus != nil) {
        paramMap["use_status"] = *req.UseStatus
    }
    if(req.TagIds != nil) {
        paramMap["tag_ids"] = *req.TagIds
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyEntpartnerPageRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}