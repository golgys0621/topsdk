package request


type AlibabaAlihealthDrugYljgSearchbillDetailRequest struct {
    /*
        企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        货主     */
    AuthRefUserId  *string `json:"auth_ref_user_id,omitempty" required:"false" `
    /*
        单据号     */
    BillCode  *string `json:"bill_code" required:"true" `
    /*
        是否显示单据中的码( 1：显示    0：不显示 )     */
    ShowCode  *string `json:"show_code" required:"true" `
}

func (s *AlibabaAlihealthDrugYljgSearchbillDetailRequest) SetRefEntId(v string) *AlibabaAlihealthDrugYljgSearchbillDetailRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailRequest) SetAuthRefUserId(v string) *AlibabaAlihealthDrugYljgSearchbillDetailRequest {
    s.AuthRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailRequest) SetBillCode(v string) *AlibabaAlihealthDrugYljgSearchbillDetailRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailRequest) SetShowCode(v string) *AlibabaAlihealthDrugYljgSearchbillDetailRequest {
    s.ShowCode = &v
    return s
}

func (req *AlibabaAlihealthDrugYljgSearchbillDetailRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.AuthRefUserId != nil) {
        paramMap["auth_ref_user_id"] = *req.AuthRefUserId
    }
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.ShowCode != nil) {
        paramMap["show_code"] = *req.ShowCode
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugYljgSearchbillDetailRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}