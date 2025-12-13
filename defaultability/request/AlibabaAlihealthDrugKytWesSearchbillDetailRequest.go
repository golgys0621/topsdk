package request


type AlibabaAlihealthDrugKytWesSearchbillDetailRequest struct {
    /*
        企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        licenseToken     */
    LicenseToken  *string `json:"license_token" required:"true" `
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

func (s *AlibabaAlihealthDrugKytWesSearchbillDetailRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailRequest) SetAuthRefUserId(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailRequest {
    s.AuthRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailRequest) SetBillCode(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailRequest) SetShowCode(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailRequest {
    s.ShowCode = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesSearchbillDetailRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
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

func (req *AlibabaAlihealthDrugKytWesSearchbillDetailRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}