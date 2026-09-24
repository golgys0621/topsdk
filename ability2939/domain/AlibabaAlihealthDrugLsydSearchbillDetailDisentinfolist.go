package domain


type AlibabaAlihealthDrugLsydSearchbillDetailDisentinfolist struct {
    /*
        配送单位refentid     */
    DisRefEntId  *string `json:"dis_ref_ent_id,omitempty" `

    /*
        配送单位名称     */
    DisEntName  *string `json:"dis_ent_name,omitempty" `

}

func (s *AlibabaAlihealthDrugLsydSearchbillDetailDisentinfolist) SetDisRefEntId(v string) *AlibabaAlihealthDrugLsydSearchbillDetailDisentinfolist {
    s.DisRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailDisentinfolist) SetDisEntName(v string) *AlibabaAlihealthDrugLsydSearchbillDetailDisentinfolist {
    s.DisEntName = &v
    return s
}
