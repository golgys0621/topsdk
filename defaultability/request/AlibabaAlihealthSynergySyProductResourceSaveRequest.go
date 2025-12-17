package request

type AlibabaAlihealthSynergySyProductResourceSaveRequest struct {
	/*
	   备注     */
	Note *string `json:"note,omitempty" required:"false" `
	/*
	   首营产品id     */
	SyProductId *int64 `json:"sy_product_id" required:"true" `
	/*
	   资料名称     */
	ResourceName *string `json:"resource_name" required:"true" `
	/*
	   有效期，长期有效传“长期”     */
	ExpireDate *string `json:"expire_date" required:"true" `
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

func (s *AlibabaAlihealthSynergySyProductResourceSaveRequest) SetNote(v string) *AlibabaAlihealthSynergySyProductResourceSaveRequest {
	s.Note = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceSaveRequest) SetSyProductId(v int64) *AlibabaAlihealthSynergySyProductResourceSaveRequest {
	s.SyProductId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceSaveRequest) SetResourceName(v string) *AlibabaAlihealthSynergySyProductResourceSaveRequest {
	s.ResourceName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceSaveRequest) SetExpireDate(v string) *AlibabaAlihealthSynergySyProductResourceSaveRequest {
	s.ExpireDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceSaveRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyProductResourceSaveRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceSaveRequest) SetResourceTypeCode(v string) *AlibabaAlihealthSynergySyProductResourceSaveRequest {
	s.ResourceTypeCode = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceSaveRequest) SetFileType(v string) *AlibabaAlihealthSynergySyProductResourceSaveRequest {
	s.FileType = &v
	return s
}
func (s *AlibabaAlihealthSynergySyProductResourceSaveRequest) SetFileContent(v []byte) *AlibabaAlihealthSynergySyProductResourceSaveRequest {
	s.FileContent = &v
	return s
}

func (req *AlibabaAlihealthSynergySyProductResourceSaveRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Note != nil {
		paramMap["note"] = *req.Note
	}
	if req.SyProductId != nil {
		paramMap["sy_product_id"] = *req.SyProductId
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

func (req *AlibabaAlihealthSynergySyProductResourceSaveRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	if req.FileContent != nil {
		fileMap["file_content"] = *req.FileContent
	}
	return fileMap
}
