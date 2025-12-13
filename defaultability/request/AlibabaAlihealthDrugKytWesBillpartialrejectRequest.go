package request


type AlibabaAlihealthDrugKytWesBillpartialrejectRequest struct {
    /*
        企业唯一标识     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        licenseToken     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        上游企业出库单的标识，通过listupout接口出参的bill_out_id查询。     */
    BillOutId  *string `json:"bill_out_id" required:"true" `
    /*
        操作类型（1整单拒收、2批次拒收、 3追溯码拒收）     */
    OperateType  *int64 `json:"operate_type" required:"true" `
    /*
        单据明细ID（多个ID用逗号分隔的字符串），当操作类型为按批次拒收时该字段必填     */
    BillDetailIds  *string `json:"bill_detail_ids,omitempty" required:"false" `
    /*
        单据中的码（多个码用逗号分隔的字符串），当操作类型为按追溯码拒收时该字段必填     */
    Codes  *string `json:"codes,omitempty" required:"false" `
    /*
        操作人     */
    Operator  *string `json:"operator" required:"true" `
    /*
        操作人联系方式     */
    Contact  *string `json:"contact" required:"true" `
    /*
        拒收原因     */
    RejectReason  *string `json:"reject_reason" required:"true" `
    /*
        委托企业id     */
    AssRefEntId  *string `json:"ass_ref_ent_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugKytWesBillpartialrejectRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesBillpartialrejectRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesBillpartialrejectRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesBillpartialrejectRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesBillpartialrejectRequest) SetBillOutId(v string) *AlibabaAlihealthDrugKytWesBillpartialrejectRequest {
    s.BillOutId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesBillpartialrejectRequest) SetOperateType(v int64) *AlibabaAlihealthDrugKytWesBillpartialrejectRequest {
    s.OperateType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesBillpartialrejectRequest) SetBillDetailIds(v string) *AlibabaAlihealthDrugKytWesBillpartialrejectRequest {
    s.BillDetailIds = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesBillpartialrejectRequest) SetCodes(v string) *AlibabaAlihealthDrugKytWesBillpartialrejectRequest {
    s.Codes = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesBillpartialrejectRequest) SetOperator(v string) *AlibabaAlihealthDrugKytWesBillpartialrejectRequest {
    s.Operator = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesBillpartialrejectRequest) SetContact(v string) *AlibabaAlihealthDrugKytWesBillpartialrejectRequest {
    s.Contact = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesBillpartialrejectRequest) SetRejectReason(v string) *AlibabaAlihealthDrugKytWesBillpartialrejectRequest {
    s.RejectReason = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesBillpartialrejectRequest) SetAssRefEntId(v string) *AlibabaAlihealthDrugKytWesBillpartialrejectRequest {
    s.AssRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesBillpartialrejectRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    if(req.BillOutId != nil) {
        paramMap["bill_out_id"] = *req.BillOutId
    }
    if(req.OperateType != nil) {
        paramMap["operate_type"] = *req.OperateType
    }
    if(req.BillDetailIds != nil) {
        paramMap["bill_detail_ids"] = *req.BillDetailIds
    }
    if(req.Codes != nil) {
        paramMap["codes"] = *req.Codes
    }
    if(req.Operator != nil) {
        paramMap["operator"] = *req.Operator
    }
    if(req.Contact != nil) {
        paramMap["contact"] = *req.Contact
    }
    if(req.RejectReason != nil) {
        paramMap["reject_reason"] = *req.RejectReason
    }
    if(req.AssRefEntId != nil) {
        paramMap["ass_ref_ent_id"] = *req.AssRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesBillpartialrejectRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}