package domain


type AlibabaAlihealthDrugKytWesSearchbillDetailDisentinfolist struct {
    /*
        配送单位refentid     */
    DisRefEntId  *string `json:"dis_ref_ent_id,omitempty" `

    /*
        配送单位名称     */
    DisEntName  *string `json:"dis_ent_name,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSearchbillDetailDisentinfolist) SetDisRefEntId(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailDisentinfolist {
    s.DisRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailDisentinfolist) SetDisEntName(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailDisentinfolist {
    s.DisEntName = &v
    return s
}
