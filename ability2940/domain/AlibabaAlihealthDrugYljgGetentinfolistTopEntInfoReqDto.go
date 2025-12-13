package domain


type AlibabaAlihealthDrugYljgGetentinfolistTopEntInfoReqDto struct {
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

func (s *AlibabaAlihealthDrugYljgGetentinfolistTopEntInfoReqDto) SetEntName(v string) *AlibabaAlihealthDrugYljgGetentinfolistTopEntInfoReqDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgGetentinfolistTopEntInfoReqDto) SetOrgCode(v string) *AlibabaAlihealthDrugYljgGetentinfolistTopEntInfoReqDto {
    s.OrgCode = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgGetentinfolistTopEntInfoReqDto) SetMedicalCode(v string) *AlibabaAlihealthDrugYljgGetentinfolistTopEntInfoReqDto {
    s.MedicalCode = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgGetentinfolistTopEntInfoReqDto) SetParRefEntId(v string) *AlibabaAlihealthDrugYljgGetentinfolistTopEntInfoReqDto {
    s.ParRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgGetentinfolistTopEntInfoReqDto) SetEntId(v string) *AlibabaAlihealthDrugYljgGetentinfolistTopEntInfoReqDto {
    s.EntId = &v
    return s
}
