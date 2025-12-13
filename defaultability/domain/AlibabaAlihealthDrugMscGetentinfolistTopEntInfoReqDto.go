package domain


type AlibabaAlihealthDrugMscGetentinfolistTopEntInfoReqDto struct {
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

func (s *AlibabaAlihealthDrugMscGetentinfolistTopEntInfoReqDto) SetEntName(v string) *AlibabaAlihealthDrugMscGetentinfolistTopEntInfoReqDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfolistTopEntInfoReqDto) SetOrgCode(v string) *AlibabaAlihealthDrugMscGetentinfolistTopEntInfoReqDto {
    s.OrgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfolistTopEntInfoReqDto) SetMedicalCode(v string) *AlibabaAlihealthDrugMscGetentinfolistTopEntInfoReqDto {
    s.MedicalCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfolistTopEntInfoReqDto) SetParRefEntId(v string) *AlibabaAlihealthDrugMscGetentinfolistTopEntInfoReqDto {
    s.ParRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscGetentinfolistTopEntInfoReqDto) SetEntId(v string) *AlibabaAlihealthDrugMscGetentinfolistTopEntInfoReqDto {
    s.EntId = &v
    return s
}
