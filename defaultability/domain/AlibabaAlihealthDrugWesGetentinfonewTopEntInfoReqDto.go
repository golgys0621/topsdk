package domain


type AlibabaAlihealthDrugWesGetentinfonewTopEntInfoReqDto struct {
    /*
        查询参数：企业名称，无其他查询条件时不能为空     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        查询参数：统一社会信用代码，无其他查询条件时不能为空     */
    OrgCode  *string `json:"org_code,omitempty" `

    /*
        查询参数：诊所备案号或医疗单位登记号，无其他查询条件时不能为空     */
    MedicalCode  *string `json:"medical_code,omitempty" `

}

func (s *AlibabaAlihealthDrugWesGetentinfonewTopEntInfoReqDto) SetEntName(v string) *AlibabaAlihealthDrugWesGetentinfonewTopEntInfoReqDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfonewTopEntInfoReqDto) SetOrgCode(v string) *AlibabaAlihealthDrugWesGetentinfonewTopEntInfoReqDto {
    s.OrgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugWesGetentinfonewTopEntInfoReqDto) SetMedicalCode(v string) *AlibabaAlihealthDrugWesGetentinfonewTopEntInfoReqDto {
    s.MedicalCode = &v
    return s
}
