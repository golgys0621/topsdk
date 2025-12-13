package request


type AlibabaAlihealthDrugLsydQueryUpoutbillcountRequest struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        单据类型     */
    BillType  *int64 `json:"bill_type,omitempty" required:"false" `
    /*
        单据上传开始日期     */
    BeginDate  *string `json:"begin_date" required:"true" `
    /*
        单据上传结束日期     */
    EndDate  *string `json:"end_date" required:"true" `
    /*
        委托企业Id     */
    AssRefEntId  *string `json:"ass_ref_ent_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugLsydQueryUpoutbillcountRequest) SetRefEntId(v string) *AlibabaAlihealthDrugLsydQueryUpoutbillcountRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryUpoutbillcountRequest) SetBillType(v int64) *AlibabaAlihealthDrugLsydQueryUpoutbillcountRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryUpoutbillcountRequest) SetBeginDate(v string) *AlibabaAlihealthDrugLsydQueryUpoutbillcountRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryUpoutbillcountRequest) SetEndDate(v string) *AlibabaAlihealthDrugLsydQueryUpoutbillcountRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryUpoutbillcountRequest) SetAssRefEntId(v string) *AlibabaAlihealthDrugLsydQueryUpoutbillcountRequest {
    s.AssRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugLsydQueryUpoutbillcountRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.BillType != nil) {
        paramMap["bill_type"] = *req.BillType
    }
    if(req.BeginDate != nil) {
        paramMap["begin_date"] = *req.BeginDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.AssRefEntId != nil) {
        paramMap["ass_ref_ent_id"] = *req.AssRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugLsydQueryUpoutbillcountRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}