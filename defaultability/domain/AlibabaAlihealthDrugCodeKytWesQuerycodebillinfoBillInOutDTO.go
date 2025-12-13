package domain


type AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO struct {
    /*
        单据类型     */
    BillType  *string `json:"bill_type,omitempty" `

    /*
        单据号码     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        单据时间     */
    BillTime  *string `json:"bill_time,omitempty" `

    /*
        发货企业id     */
    FromUserId  *string `json:"from_user_id,omitempty" `

    /*
        发货企业refEntId     */
    FromRefUserId  *string `json:"from_ref_user_id,omitempty" `

    /*
        发货企业名称     */
    FromUserName  *string `json:"from_user_name,omitempty" `

    /*
        收货企业id     */
    ToUserId  *string `json:"to_user_id,omitempty" `

    /*
        收货企业refEntId     */
    ToRefUserId  *string `json:"to_ref_user_id,omitempty" `

    /*
        收货企业名称     */
    ToUserName  *string `json:"to_user_name,omitempty" `

    /*
        代理企业id     */
    AssEntId  *string `json:"ass_ent_id,omitempty" `

    /*
        代理企业refEntId     */
    AssRefEntId  *string `json:"ass_ref_ent_id,omitempty" `

    /*
        代理企业名称     */
    AssEntName  *string `json:"ass_ent_name,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO) SetBillType(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO) SetBillCode(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO) SetBillTime(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO) SetFromUserId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO) SetFromRefUserId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO) SetFromUserName(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO) SetToUserId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO) SetToRefUserId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO) SetToUserName(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO) SetAssEntId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO {
    s.AssEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO) SetAssRefEntId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO {
    s.AssRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO) SetAssEntName(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoBillInOutDTO {
    s.AssEntName = &v
    return s
}
