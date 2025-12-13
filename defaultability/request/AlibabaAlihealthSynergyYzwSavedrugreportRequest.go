package request

type AlibabaAlihealthSynergyYzwSavedrugreportRequest struct {
	/*
	   企业ID     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   药品ID     */
	DrugEntBaseInfoId *string `json:"drug_ent_base_info_id" required:"true" `
	/*
	   批次号     */
	BatchNo *string `json:"batch_no" required:"true" `
	/*
	   报告数据     */
	ReportData *[]byte `json:"report_data" required:"true" `
	/*
	   文件类型支持:pdf,jpg,jpeg, png     */
	FileType *string `json:"file_type" required:"true" `
	/*
	   用户id     */
	UserId *string `json:"user_id" required:"true" `
	/*
	   印章名称     */
	SealName *string `json:"seal_name" required:"true" `
	/*
	   药检报告号     */
	ReportNo *string `json:"report_no,omitempty" required:"false" `
	/*
	   报告日期     */
	ReportDate *string `json:"report_date,omitempty" required:"false" `
	/*
	   生产日期     */
	ProduceDate *string `json:"produce_date,omitempty" required:"false" `
	/*
	   X坐标     */
	PositionX *int64 `json:"position_x,omitempty" required:"false" `
	/*
	   Y坐标     */
	PositionY *int64 `json:"position_y,omitempty" required:"false" `
	/*
	   是否要显示日期 defalutValue��false    */
	NoDate *bool `json:"no_date,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergyYzwSavedrugreportRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwSavedrugreportRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportRequest) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthSynergyYzwSavedrugreportRequest {
	s.DrugEntBaseInfoId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportRequest) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwSavedrugreportRequest {
	s.BatchNo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportRequest) SetReportData(v []byte) *AlibabaAlihealthSynergyYzwSavedrugreportRequest {
	s.ReportData = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportRequest) SetFileType(v string) *AlibabaAlihealthSynergyYzwSavedrugreportRequest {
	s.FileType = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportRequest) SetUserId(v string) *AlibabaAlihealthSynergyYzwSavedrugreportRequest {
	s.UserId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportRequest) SetSealName(v string) *AlibabaAlihealthSynergyYzwSavedrugreportRequest {
	s.SealName = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportRequest) SetReportNo(v string) *AlibabaAlihealthSynergyYzwSavedrugreportRequest {
	s.ReportNo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportRequest) SetReportDate(v string) *AlibabaAlihealthSynergyYzwSavedrugreportRequest {
	s.ReportDate = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportRequest) SetProduceDate(v string) *AlibabaAlihealthSynergyYzwSavedrugreportRequest {
	s.ProduceDate = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportRequest) SetPositionX(v int64) *AlibabaAlihealthSynergyYzwSavedrugreportRequest {
	s.PositionX = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportRequest) SetPositionY(v int64) *AlibabaAlihealthSynergyYzwSavedrugreportRequest {
	s.PositionY = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportRequest) SetNoDate(v bool) *AlibabaAlihealthSynergyYzwSavedrugreportRequest {
	s.NoDate = &v
	return s
}

func (req *AlibabaAlihealthSynergyYzwSavedrugreportRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.DrugEntBaseInfoId != nil {
		paramMap["drug_ent_base_info_id"] = *req.DrugEntBaseInfoId
	}
	if req.BatchNo != nil {
		paramMap["batch_no"] = *req.BatchNo
	}
	if req.FileType != nil {
		paramMap["file_type"] = *req.FileType
	}
	if req.UserId != nil {
		paramMap["user_id"] = *req.UserId
	}
	if req.SealName != nil {
		paramMap["seal_name"] = *req.SealName
	}
	if req.ReportNo != nil {
		paramMap["report_no"] = *req.ReportNo
	}
	if req.ReportDate != nil {
		paramMap["report_date"] = *req.ReportDate
	}
	if req.ProduceDate != nil {
		paramMap["produce_date"] = *req.ProduceDate
	}
	if req.PositionX != nil {
		paramMap["position_x"] = *req.PositionX
	}
	if req.PositionY != nil {
		paramMap["position_y"] = *req.PositionY
	}
	if req.NoDate != nil {
		paramMap["no_date"] = *req.NoDate
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergyYzwSavedrugreportRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	if req.ReportData != nil {
		fileMap["report_data"] = *req.ReportData
	}
	return fileMap
}
