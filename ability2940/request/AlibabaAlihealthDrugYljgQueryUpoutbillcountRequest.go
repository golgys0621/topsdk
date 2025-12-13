package request


type AlibabaAlihealthDrugYljgQueryUpoutbillcountRequest struct {
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

func (s *AlibabaAlihealthDrugYljgQueryUpoutbillcountRequest) SetRefEntId(v string) *AlibabaAlihealthDrugYljgQueryUpoutbillcountRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryUpoutbillcountRequest) SetBillType(v int64) *AlibabaAlihealthDrugYljgQueryUpoutbillcountRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryUpoutbillcountRequest) SetBeginDate(v string) *AlibabaAlihealthDrugYljgQueryUpoutbillcountRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryUpoutbillcountRequest) SetEndDate(v string) *AlibabaAlihealthDrugYljgQueryUpoutbillcountRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgQueryUpoutbillcountRequest) SetAssRefEntId(v string) *AlibabaAlihealthDrugYljgQueryUpoutbillcountRequest {
    s.AssRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugYljgQueryUpoutbillcountRequest) ToMap() map[string]interface{} {
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

func (req *AlibabaAlihealthDrugYljgQueryUpoutbillcountRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}