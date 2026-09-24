package request


type AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest struct {
    /*
        备注     */
    Note  *string `json:"note,omitempty" required:"false" `
    /*
        首营委托人id     */
    SyAgentPersonId  *int64 `json:"sy_agent_person_id" required:"true" `
    /*
        资料名称     */
    ResourceName  *string `json:"resource_name" required:"true" `
    /*
        有效期。ai识别场景必传；非ai识别场景，除"未分类"外必传；长期有效传“长期”     */
    ExpireDate  *string `json:"expire_date,omitempty" required:"false" `
    /*
        refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        资料类别，选择ai识别时无需传递     */
    ResourceTypeCode  *string `json:"resource_type_code,omitempty" required:"false" `
    /*
        文件类型，只支持"pdf", "jpg", "jpeg", "png","tif"     */
    FileType  *string `json:"file_type" required:"true" `
    /*
        资料文件与fileUrl字段填写一个即可     */
    FileContent  *[]byte `json:"file_content,omitempty" required:"false" `
    /*
        上传文件的httpurl地址     */
    FileUrl  *string `json:"file_url,omitempty" required:"false" `
    /*
        1:资料类别由ai识别获取     */
    NeedAi  *string `json:"need_ai,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest) SetNote(v string) *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest {
    s.Note = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest) SetSyAgentPersonId(v int64) *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest {
    s.SyAgentPersonId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest) SetResourceName(v string) *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest {
    s.ResourceName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest) SetExpireDate(v string) *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest {
    s.ExpireDate = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest) SetResourceTypeCode(v string) *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest {
    s.ResourceTypeCode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest) SetFileType(v string) *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest {
    s.FileType = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest) SetFileContent(v []byte) *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest {
    s.FileContent = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest) SetFileUrl(v string) *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest {
    s.FileUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest) SetNeedAi(v string) *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest {
    s.NeedAi = &v
    return s
}

func (req *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Note != nil) {
        paramMap["note"] = *req.Note
    }
    if(req.SyAgentPersonId != nil) {
        paramMap["sy_agent_person_id"] = *req.SyAgentPersonId
    }
    if(req.ResourceName != nil) {
        paramMap["resource_name"] = *req.ResourceName
    }
    if(req.ExpireDate != nil) {
        paramMap["expire_date"] = *req.ExpireDate
    }
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.ResourceTypeCode != nil) {
        paramMap["resource_type_code"] = *req.ResourceTypeCode
    }
    if(req.FileType != nil) {
        paramMap["file_type"] = *req.FileType
    }
    if(req.FileUrl != nil) {
        paramMap["file_url"] = *req.FileUrl
    }
    if(req.NeedAi != nil) {
        paramMap["need_ai"] = *req.NeedAi
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.FileContent != nil {
        fileMap["file_content"] = *req.FileContent
    }
    return fileMap
}