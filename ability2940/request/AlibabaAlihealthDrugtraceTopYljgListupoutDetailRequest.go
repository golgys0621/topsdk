package request


type AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest struct {
    /*
        企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        单据编码     */
    BillCode  *string `json:"bill_code" required:"true" `
    /*
        发货企业refEntId【发货企业和委托企业必须填写一个，否则查询不出来数据】     */
    FromRefUserId  *string `json:"from_ref_user_id,omitempty" required:"false" `
    /*
        收货企业refEntId     */
    ToRefUserId  *string `json:"to_ref_user_id" required:"true" `
    /*
        委托企业refEntId【发货企业和委托企业必须填写一个，否则查询不出来数据】     */
    AssRefEntId  *string `json:"ass_ref_ent_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) SetRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) SetBillCode(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) SetFromRefUserId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) SetToRefUserId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) SetAssRefEntId(v string) *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest {
    s.AssRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) ToMap() map[string]interface{} {
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

func (req *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}