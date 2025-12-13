package domain


type AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoReqDto struct {
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

func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoReqDto) SetEntName(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoReqDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoReqDto) SetOrgCode(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoReqDto {
    s.OrgCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoReqDto) SetMedicalCode(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoReqDto {
    s.MedicalCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoReqDto) SetParRefEntId(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoReqDto {
    s.ParRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoReqDto) SetEntId(v string) *AlibabaAlihealthSynergyYzwGetentinfolistTopEntInfoReqDto {
    s.EntId = &v
    return s
}
