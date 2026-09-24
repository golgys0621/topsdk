package request


type AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest struct {
    /*
        企业refentid(填收货单位的ID)     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense；填本企业的或者三方物流企业的license     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        单据号     */
    BillCode  *string `json:"bill_code" required:"true" `
    /*
        发货企业renEntId     */
    FromRefUserId  *string `json:"from_ref_user_id" required:"true" `
    /*
        收货企业refEntId     */
    ToRefUserId  *string `json:"to_ref_user_id" required:"true" `
    /*
        委托企业refentid(如果是三方的模式，需要填三方物流企业的ID)，注意收货企业必须在平台 "授权管理/授权配送单位"里面，授权三方物流，查询本企业的上游出库单     */
    AgentRefEntId  *string `json:"agent_ref_ent_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest) SetBillCode(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest) SetFromRefUserId(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest) SetToRefUserId(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest) SetAgentRefEntId(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest {
    s.AgentRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
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
    if(req.AgentRefEntId != nil) {
        paramMap["agent_ref_ent_id"] = *req.AgentRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}