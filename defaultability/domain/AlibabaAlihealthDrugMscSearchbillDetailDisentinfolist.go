package domain


type AlibabaAlihealthDrugMscSearchbillDetailDisentinfolist struct {
    /*
        配送单位refentid     */
    DisRefEntId  *string `json:"dis_ref_ent_id,omitempty" `

    /*
        配送单位名称     */
    DisEntName  *string `json:"dis_ent_name,omitempty" `

}

func (s *AlibabaAlihealthDrugMscSearchbillDetailDisentinfolist) SetDisRefEntId(v string) *AlibabaAlihealthDrugMscSearchbillDetailDisentinfolist {
    s.DisRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscSearchbillDetailDisentinfolist) SetDisEntName(v string) *AlibabaAlihealthDrugMscSearchbillDetailDisentinfolist {
    s.DisEntName = &v
    return s
}
