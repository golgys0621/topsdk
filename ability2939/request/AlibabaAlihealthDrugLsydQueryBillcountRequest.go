package request


type AlibabaAlihealthDrugLsydQueryBillcountRequest struct {
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
}

func (s *AlibabaAlihealthDrugLsydQueryBillcountRequest) SetRefEntId(v string) *AlibabaAlihealthDrugLsydQueryBillcountRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryBillcountRequest) SetBillType(v int64) *AlibabaAlihealthDrugLsydQueryBillcountRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryBillcountRequest) SetBeginDate(v string) *AlibabaAlihealthDrugLsydQueryBillcountRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydQueryBillcountRequest) SetEndDate(v string) *AlibabaAlihealthDrugLsydQueryBillcountRequest {
    s.EndDate = &v
    return s
}

func (req *AlibabaAlihealthDrugLsydQueryBillcountRequest) ToMap() map[string]interface{} {
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
    return paramMap
}

func (req *AlibabaAlihealthDrugLsydQueryBillcountRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}