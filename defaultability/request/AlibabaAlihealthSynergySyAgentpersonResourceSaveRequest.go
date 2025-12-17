package request

type AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest struct {
	/*
	   备注     */
	Note *string `json:"note,omitempty" required:"false" `
	/*
	   首营委托人id     */
	SyAgentPersonId *int64 `json:"sy_agent_person_id" required:"true" `
	/*
	   资料名称     */
	ResourceName *string `json:"resource_name" required:"true" `
	/*
	   有效期，除"未分类"外必传，长期有效传“长期”     */
	ExpireDate *string `json:"expire_date,omitempty" required:"false" `
	/*
	   refEntId     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   资料类别     */
	ResourceTypeCode *string `json:"resource_type_code" required:"true" `
	/*
	   文件类型，只支持"pdf", "jpg", "jpeg", "png","tif"     */
	FileType *string `json:"file_type" required:"true" `
	/*
	   文件内容     */
	FileContent *[]byte `json:"file_content" required:"true" `
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

func (req *AlibabaAlihealthSynergySyAgentpersonResourceSaveRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Note != nil {
		paramMap["note"] = *req.Note
	}
	if req.SyAgentPersonId != nil {
		paramMap["sy_agent_person_id"] = *req.SyAgentPersonId
	}
	if req.ResourceName != nil {
		paramMap["resource_name"] = *req.ResourceName
	}
	if req.ExpireDate != nil {
		paramMap["expire_date"] = *req.ExpireDate
	}
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.ResourceTypeCode != nil {
		paramMap["resource_type_code"] = *req.ResourceTypeCode
	}
	if req.FileType != nil {
		paramMap["file_type"] = *req.FileType
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
