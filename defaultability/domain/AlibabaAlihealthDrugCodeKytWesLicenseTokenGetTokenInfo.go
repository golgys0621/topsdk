package domain


type AlibabaAlihealthDrugCodeKytWesLicenseTokenGetTokenInfo struct {
    /*
        appkey     */
    Appkey  *string `json:"appkey,omitempty" `

    /*
        调用企业id     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        licenseToken，有效期1天，失效10分钟内调用可以获得新的token     */
    LicenseToken  *string `json:"license_token,omitempty" `

    /*
        token生成时间yyyy-dd-MM hh:mm;ss     */
    CreateTime  *string `json:"create_time,omitempty" `

    /*
        token过期时间yyyy-dd-MM hh:mm;ss     */
    ExpireTime  *string `json:"expire_time,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetTokenInfo) SetAppkey(v string) *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetTokenInfo {
    s.Appkey = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetTokenInfo) SetRefEntId(v string) *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetTokenInfo {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetTokenInfo) SetLicenseToken(v string) *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetTokenInfo {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetTokenInfo) SetCreateTime(v string) *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetTokenInfo {
    s.CreateTime = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetTokenInfo) SetExpireTime(v string) *AlibabaAlihealthDrugCodeKytWesLicenseTokenGetTokenInfo {
    s.ExpireTime = &v
    return s
}
