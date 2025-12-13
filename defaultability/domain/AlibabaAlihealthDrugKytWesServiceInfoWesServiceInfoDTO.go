package domain


type AlibabaAlihealthDrugKytWesServiceInfoWesServiceInfoDTO struct {
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

func (s *AlibabaAlihealthDrugKytWesServiceInfoWesServiceInfoDTO) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesServiceInfoWesServiceInfoDTO {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesServiceInfoWesServiceInfoDTO) SetStartDate(v string) *AlibabaAlihealthDrugKytWesServiceInfoWesServiceInfoDTO {
    s.StartDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesServiceInfoWesServiceInfoDTO) SetEndDate(v string) *AlibabaAlihealthDrugKytWesServiceInfoWesServiceInfoDTO {
    s.EndDate = &v
    return s
}
