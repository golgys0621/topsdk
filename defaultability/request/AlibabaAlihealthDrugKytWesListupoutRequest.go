package request


type AlibabaAlihealthDrugKytWesListupoutRequest struct {
    /*
        货主企业唯一标识（一般情况下填写自已企业）     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        licenseToken     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        单据时间的开始日期（不写时分秒），格式：yyyy-MM-dd     */
    BeginDate  *string `json:"begin_date,omitempty" required:"false" `
    /*
        单据时间的结束日期（不写时分秒），格式：yyyy-MM-dd     */
    EndDate  *string `json:"end_date,omitempty" required:"false" `
    /*
        发货单位ent_id     */
    FromUserId  *string `json:"from_user_id,omitempty" required:"false" `
    /*
        生产批号     */
    ProduceBatchNo  *string `json:"produce_batch_no,omitempty" required:"false" `
    /*
        药品ID     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" required:"false" `
    /*
        单据类型     */
    BillType  *string `json:"bill_type,omitempty" required:"false" `
    /*
        药品类型     */
    PhysicType  *string `json:"physic_type,omitempty" required:"false" `
    /*
        状态     */
    Status  *string `json:"status,omitempty" required:"false" `
    /*
        单据号     */
    BillCode  *string `json:"bill_code,omitempty" required:"false" `
    /*
        第三方物流企业唯一标识（只有代查其它企业数据时填写）     */
    AgentRefEntId  *string `json:"agent_ref_ent_id,omitempty" required:"false" `
    /*
        单据上传的开始日期，格式：yyyy-MM-dd HH:mm:ss     */
    UploadTimeBegin  *string `json:"upload_time_begin,omitempty" required:"false" `
    /*
        单据上传的结束日期，格式：yyyy-MM-dd HH:mm:ss     */
    UploadTimeEnd  *string `json:"upload_time_end,omitempty" required:"false" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码     */
    Page  *int64 `json:"page" required:"true" `
}

func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetBeginDate(v string) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetEndDate(v string) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetFromUserId(v string) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetProduceBatchNo(v string) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetBillType(v string) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetPhysicType(v string) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetStatus(v string) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetBillCode(v string) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetAgentRefEntId(v string) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.AgentRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetUploadTimeBegin(v string) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.UploadTimeBegin = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetUploadTimeEnd(v string) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.UploadTimeEnd = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetPageSize(v int64) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutRequest) SetPage(v int64) *AlibabaAlihealthDrugKytWesListupoutRequest {
    s.Page = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesListupoutRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    if(req.BeginDate != nil) {
        paramMap["begin_date"] = *req.BeginDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.FromUserId != nil) {
        paramMap["from_user_id"] = *req.FromUserId
    }
    if(req.ProduceBatchNo != nil) {
        paramMap["produce_batch_no"] = *req.ProduceBatchNo
    }
    if(req.DrugEntBaseInfoId != nil) {
        paramMap["drug_ent_base_info_id"] = *req.DrugEntBaseInfoId
    }
    if(req.BillType != nil) {
        paramMap["bill_type"] = *req.BillType
    }
    if(req.PhysicType != nil) {
        paramMap["physic_type"] = *req.PhysicType
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.AgentRefEntId != nil) {
        paramMap["agent_ref_ent_id"] = *req.AgentRefEntId
    }
    if(req.UploadTimeBegin != nil) {
        paramMap["upload_time_begin"] = *req.UploadTimeBegin
    }
    if(req.UploadTimeEnd != nil) {
        paramMap["upload_time_end"] = *req.UploadTimeEnd
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesListupoutRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}