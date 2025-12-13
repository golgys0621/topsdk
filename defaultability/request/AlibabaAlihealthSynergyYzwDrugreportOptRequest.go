package request


type AlibabaAlihealthSynergyYzwDrugreportOptRequest struct {
    /*
        调用接口的企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        用户id     */
    UserId  *string `json:"user_id" required:"true" `
    /*
        操作类型：签收(onlySign),拒收(rejectReceive),签收并盖章(signAndSeal),盖章签收后的报告(sealAfterSign)     */
    OperType  *string `json:"oper_type" required:"true" `
    /*
        单据id     */
    BillId  *int64 `json:"bill_id" required:"true" `
    /*
        单据明细id     */
    DetailBillId  *int64 `json:"detail_bill_id" required:"true" `
    /*
        单据类型     */
    BillType  *int64 `json:"bill_type" required:"true" `
    /*
        印章名称，当操作类型为signAndSeal和sealAfterSign时，必填     */
    SealName  *string `json:"seal_name,omitempty" required:"false" `
    /*
        X坐标     */
    PositionX  *int64 `json:"position_x,omitempty" required:"false" `
    /*
        Y坐标     */
    PositionY  *int64 `json:"position_y,omitempty" required:"false" `
    /*
        拒绝原因     */
    OperNote  *string `json:"oper_note,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergyYzwDrugreportOptRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwDrugreportOptRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptRequest) SetUserId(v string) *AlibabaAlihealthSynergyYzwDrugreportOptRequest {
    s.UserId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptRequest) SetOperType(v string) *AlibabaAlihealthSynergyYzwDrugreportOptRequest {
    s.OperType = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptRequest) SetBillId(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptRequest {
    s.BillId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptRequest) SetDetailBillId(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptRequest {
    s.DetailBillId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptRequest) SetBillType(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptRequest) SetSealName(v string) *AlibabaAlihealthSynergyYzwDrugreportOptRequest {
    s.SealName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptRequest) SetPositionX(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptRequest {
    s.PositionX = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptRequest) SetPositionY(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptRequest {
    s.PositionY = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptRequest) SetOperNote(v string) *AlibabaAlihealthSynergyYzwDrugreportOptRequest {
    s.OperNote = &v
    return s
}

func (req *AlibabaAlihealthSynergyYzwDrugreportOptRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.UserId != nil) {
        paramMap["user_id"] = *req.UserId
    }
    if(req.OperType != nil) {
        paramMap["oper_type"] = *req.OperType
    }
    if(req.BillId != nil) {
        paramMap["bill_id"] = *req.BillId
    }
    if(req.DetailBillId != nil) {
        paramMap["detail_bill_id"] = *req.DetailBillId
    }
    if(req.BillType != nil) {
        paramMap["bill_type"] = *req.BillType
    }
    if(req.SealName != nil) {
        paramMap["seal_name"] = *req.SealName
    }
    if(req.PositionX != nil) {
        paramMap["position_x"] = *req.PositionX
    }
    if(req.PositionY != nil) {
        paramMap["position_y"] = *req.PositionY
    }
    if(req.OperNote != nil) {
        paramMap["oper_note"] = *req.OperNote
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergyYzwDrugreportOptRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}