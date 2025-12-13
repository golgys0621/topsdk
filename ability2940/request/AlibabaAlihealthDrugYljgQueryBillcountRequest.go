package request


type AlibabaAlihealthDrugYljgQueryBillcountRequest struct {
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

func (s *AlibabaAlihealthDrugYljgQueryBillcountRequest) SetRefEntId(v string) *AlibabaAlihealthDrugYljgQueryBillcountRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryBillcountRequest) SetBillType(v int64) *AlibabaAlihealthDrugYljgQueryBillcountRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryBillcountRequest) SetBeginDate(v string) *AlibabaAlihealthDrugYljgQueryBillcountRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryBillcountRequest) SetEndDate(v string) *AlibabaAlihealthDrugYljgQueryBillcountRequest {
    s.EndDate = &v
    return s
}

func (req *AlibabaAlihealthDrugYljgQueryBillcountRequest) ToMap() map[string]interface{} {
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

func (req *AlibabaAlihealthDrugYljgQueryBillcountRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}