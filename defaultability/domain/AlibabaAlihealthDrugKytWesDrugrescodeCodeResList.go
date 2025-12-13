package domain


type AlibabaAlihealthDrugKytWesDrugrescodeCodeResList struct {
    /*
        码前缀     */
    CodePrefix  *string `json:"code_prefix,omitempty" `

    /*
        资源码     */
    ResCode  *string `json:"res_code,omitempty" `

    /*
        层级     */
    CodeLevel  *string `json:"code_level,omitempty" `

    /*
        包装比例     */
    PkgRatio  *string `json:"pkg_ratio,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesDrugrescodeCodeResList) SetCodePrefix(v string) *AlibabaAlihealthDrugKytWesDrugrescodeCodeResList {
    s.CodePrefix = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeCodeResList) SetResCode(v string) *AlibabaAlihealthDrugKytWesDrugrescodeCodeResList {
    s.ResCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeCodeResList) SetCodeLevel(v string) *AlibabaAlihealthDrugKytWesDrugrescodeCodeResList {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeCodeResList) SetPkgRatio(v string) *AlibabaAlihealthDrugKytWesDrugrescodeCodeResList {
    s.PkgRatio = &v
    return s
}
