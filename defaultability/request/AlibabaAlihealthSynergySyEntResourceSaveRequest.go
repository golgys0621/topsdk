package request

type AlibabaAlihealthSynergySyEntResourceSaveRequest struct {
	/*
	   上传企业     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   资质分类编码     */
	ResourceTypeCode *string `json:"resource_type_code" required:"true" `
	/*
	   资料文件名称，不能重复     */
	ResourceName *string `json:"resource_name" required:"true" `
	/*
	   资料二进制文件     */
	ResourceData *[]byte `json:"resource_data" required:"true" `
	/*
	   文件类型，支持pdf和图片格式     */
	FileType *string `json:"file_type" required:"true" `
	/*
	   过期时间，格式yyyy-MM-dd或长期     */
	ExpireDate *string `json:"expire_date" required:"true" `
	/*
	   扩展属性，json对象数组，code：资质支持的属性code，value:值     */
	ExtendInfo *string `json:"extend_info,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergySyEntResourceSaveRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyEntResourceSaveRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyEntResourceSaveRequest) SetResourceTypeCode(v string) *AlibabaAlihealthSynergySyEntResourceSaveRequest {
	s.ResourceTypeCode = &v
	return s
}
func (s *AlibabaAlihealthSynergySyEntResourceSaveRequest) SetResourceName(v string) *AlibabaAlihealthSynergySyEntResourceSaveRequest {
	s.ResourceName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyEntResourceSaveRequest) SetResourceData(v []byte) *AlibabaAlihealthSynergySyEntResourceSaveRequest {
	s.ResourceData = &v
	return s
}
func (s *AlibabaAlihealthSynergySyEntResourceSaveRequest) SetFileType(v string) *AlibabaAlihealthSynergySyEntResourceSaveRequest {
	s.FileType = &v
	return s
}
func (s *AlibabaAlihealthSynergySyEntResourceSaveRequest) SetExpireDate(v string) *AlibabaAlihealthSynergySyEntResourceSaveRequest {
	s.ExpireDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyEntResourceSaveRequest) SetExtendInfo(v string) *AlibabaAlihealthSynergySyEntResourceSaveRequest {
	s.ExtendInfo = &v
	return s
}

func (req *AlibabaAlihealthSynergySyEntResourceSaveRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.ResourceTypeCode != nil {
		paramMap["resource_type_code"] = *req.ResourceTypeCode
	}
	if req.ResourceName != nil {
		paramMap["resource_name"] = *req.ResourceName
	}
	if req.FileType != nil {
		paramMap["file_type"] = *req.FileType
	}
	if req.ExpireDate != nil {
		paramMap["expire_date"] = *req.ExpireDate
	}
	if req.ExtendInfo != nil {
		paramMap["extend_info"] = *req.ExtendInfo
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergySyEntResourceSaveRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	if req.ResourceData != nil {
		fileMap["resource_data"] = *req.ResourceData
	}
	return fileMap
}
