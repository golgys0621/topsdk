package domain


type AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo struct {
    /*
        单据类型     */
    BillType  *string `json:"bill_type,omitempty" `

    /*
        单号号码     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        发货企业ID     */
    FromUserId  *string `json:"from_user_id,omitempty" `

    /*
        企业名称     */
    RefUserName  *string `json:"ref_user_name,omitempty" `

    /*
        企业ID     */
    RefUserId  *string `json:"ref_user_id,omitempty" `

    /*
        生产日期     */
    ProduceDate  *string `json:"produce_date,omitempty" `

    /*
        上传文件名     */
    UploadFileName  *string `json:"upload_file_name,omitempty" `

    /*
        发货单位     */
    FromUserName  *string `json:"from_user_name,omitempty" `

    /*
        收货单位     */
    ToUserId  *string `json:"to_user_id,omitempty" `

    /*
        生产企业ID     */
    ProduceEntId  *string `json:"produce_ent_id,omitempty" `

    /*
        单据时间     */
    BillTime  *string `json:"bill_time,omitempty" `

    /*
        角色类型     */
    UserRoleType  *string `json:"user_role_type,omitempty" `

    /*
        处理日期     */
    ProcessDate  *string `json:"process_date,omitempty" `

    /*
        单据ID     */
    BillId  *string `json:"bill_id,omitempty" `

    /*
        收货单位     */
    ToUserName  *string `json:"to_user_name,omitempty" `

    /*
        代理企业     */
    AgentUserName  *string `json:"agent_user_name,omitempty" `

    /*
        代理企业ID     */
    AgentRefUserId  *string `json:"agent_ref_user_id,omitempty" `

    /*
        单据类型     */
    BillTypeName  *string `json:"bill_type_name,omitempty" `

    /*
        收货单位ID     */
    ToRefUserId  *string `json:"to_ref_user_id,omitempty" `

    /*
        发货单位ID     */
    FromRefUserId  *string `json:"from_ref_user_id,omitempty" `

    /*
        上传日期     */
    CrtDate  *string `json:"crt_date,omitempty" `

    /*
        委托企业     */
    AssEntName  *string `json:"ass_ent_name,omitempty" `

    /*
        委托企业entId     */
    AssEntId  *string `json:"ass_ent_id,omitempty" `

    /*
        委托企业assRefEntId     */
    AssRefEntId  *string `json:"ass_ref_ent_id,omitempty" `

    /*
        51全部成功 52部分成功     */
    SubProcessFlag  *string `json:"sub_process_flag,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetBillType(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetBillCode(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetFromUserId(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetRefUserName(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.RefUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetRefUserId(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.RefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetProduceDate(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetUploadFileName(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.UploadFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetFromUserName(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetToUserId(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetProduceEntId(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.ProduceEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetBillTime(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetUserRoleType(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetProcessDate(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetBillId(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.BillId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetToUserName(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetAgentUserName(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.AgentUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetAgentRefUserId(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.AgentRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetBillTypeName(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetToRefUserId(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetFromRefUserId(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetCrtDate(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetAssEntName(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.AssEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetAssEntId(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.AssEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetAssRefEntId(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.AssRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo) SetSubProcessFlag(v string) *AlibabaAlihealthDrugKytWesSearchbillBillChkInOutDo {
    s.SubProcessFlag = &v
    return s
}
