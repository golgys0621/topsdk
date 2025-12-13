package request


type AlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest struct {
    /*
        企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        单据编码     */
    BillCode  *string `json:"bill_code" required:"true" `
    /*
        发货企业renEntId     */
    FromRefUserId  *string `json:"from_ref_user_id" required:"true" `
    /*
        收货企业refEntId     */
    ToRefUserId  *string `json:"to_ref_user_id" required:"true" `
    /*
        委托企业Id     */
    AssRefEntId  *string `json:"ass_ref_ent_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest) SetFromRefUserId(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest) SetToRefUserId(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest) SetAssRefEntId(v string) *AlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest {
    s.AssRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.FromRefUserId != nil) {
        paramMap["from_ref_user_id"] = *req.FromRefUserId
    }
    if(req.ToRefUserId != nil) {
        paramMap["to_ref_user_id"] = *req.ToRefUserId
    }
    if(req.AssRefEntId != nil) {
        paramMap["ass_ref_ent_id"] = *req.AssRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}