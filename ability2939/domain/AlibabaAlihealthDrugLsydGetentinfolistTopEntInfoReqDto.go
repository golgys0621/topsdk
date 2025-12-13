package domain


type AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoReqDto struct {
    /*
        查询参数：企业名称，无其他查询条件时不能为空     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        查询参数：统一社会信用代码，无其他查询条件时不能为空     */
    OrgCode  *string `json:"org_code,omitempty" `

    /*
        查询参数：诊所备案号或医疗单位登记号，无其他查询条件时不能为空     */
    MedicalCode  *string `json:"medical_code,omitempty" `

    /*
        查询参数：企业refEntId     */
    ParRefEntId  *string `json:"par_ref_ent_id,omitempty" `

    /*
        查询参数：企业entId     */
    EntId  *string `json:"ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoReqDto) SetEntName(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoReqDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoReqDto) SetOrgCode(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoReqDto {
    s.OrgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoReqDto) SetMedicalCode(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoReqDto {
    s.MedicalCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoReqDto) SetParRefEntId(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoReqDto {
    s.ParRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoReqDto) SetEntId(v string) *AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoReqDto {
    s.EntId = &v
    return s
}
