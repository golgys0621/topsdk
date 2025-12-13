package domain


type AlibabaAlihealthDrugMscServiceinfoMSCServiceInfoDTO struct {
    /*
        企业的refEntId     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        服务生效日期     */
    StartDate  *string `json:"start_date,omitempty" `

    /*
        服务到期日期     */
    EndDate  *string `json:"end_date,omitempty" `

}

func (s *AlibabaAlihealthDrugMscServiceinfoMSCServiceInfoDTO) SetRefEntId(v string) *AlibabaAlihealthDrugMscServiceinfoMSCServiceInfoDTO {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscServiceinfoMSCServiceInfoDTO) SetStartDate(v string) *AlibabaAlihealthDrugMscServiceinfoMSCServiceInfoDTO {
    s.StartDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscServiceinfoMSCServiceInfoDTO) SetEndDate(v string) *AlibabaAlihealthDrugMscServiceinfoMSCServiceInfoDTO {
    s.EndDate = &v
    return s
}
