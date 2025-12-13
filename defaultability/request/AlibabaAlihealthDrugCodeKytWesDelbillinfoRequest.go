package request


type AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest struct {
    /*
        当前操作企业唯一标识     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        单据号     */
    BillCode  *string `json:"bill_code" required:"true" `
    /*
        删除原因     */
    DeleteReason  *string `json:"delete_reason" required:"true" `
    /*
        删除人手机号     */
    MobilePhone  *string `json:"mobile_phone" required:"true" `
    /*
        当单据是代理企业上传的，需要填写货主的refEntId     */
    OwnerRefUserId  *string `json:"owner_ref_user_id,omitempty" required:"false" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
    LicenseToken  *string `json:"license_token" required:"true" `
}

func (s *AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest) SetRefEntId(v string) *AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest) SetBillCode(v string) *AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest) SetDeleteReason(v string) *AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest {
    s.DeleteReason = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest) SetMobilePhone(v string) *AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest {
    s.MobilePhone = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest) SetOwnerRefUserId(v string) *AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest {
    s.OwnerRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest {
    s.LicenseToken = &v
    return s
}

func (req *AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.DeleteReason != nil) {
        paramMap["delete_reason"] = *req.DeleteReason
    }
    if(req.MobilePhone != nil) {
        paramMap["mobile_phone"] = *req.MobilePhone
    }
    if(req.OwnerRefUserId != nil) {
        paramMap["owner_ref_user_id"] = *req.OwnerRefUserId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}