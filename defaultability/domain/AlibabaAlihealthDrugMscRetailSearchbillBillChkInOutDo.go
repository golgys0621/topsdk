package domain


type AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo struct {
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
        51全部成功 52部分成功     */
    SubProcessFlag  *string `json:"sub_process_flag,omitempty" `

    /*
        委托企业refEntId     */
    AssRefEntId  *string `json:"ass_ref_ent_id,omitempty" `

    /*
        委托企业EntId     */
    AssEntId  *string `json:"ass_ent_id,omitempty" `

    /*
        委托企业名称     */
    AssEntName  *string `json:"ass_ent_name,omitempty" `

}

func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetBillType(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetBillCode(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetFromUserId(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetRefUserName(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.RefUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetRefUserId(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.RefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetProduceDate(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetUploadFileName(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.UploadFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetFromUserName(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetToUserId(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetProduceEntId(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.ProduceEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetBillTime(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetUserRoleType(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetProcessDate(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetBillId(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.BillId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetToUserName(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetAgentUserName(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.AgentUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetAgentRefUserId(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.AgentRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetBillTypeName(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetToRefUserId(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetFromRefUserId(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetSubProcessFlag(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.SubProcessFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetAssRefEntId(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.AssRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetAssEntId(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.AssEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo) SetAssEntName(v string) *AlibabaAlihealthDrugMscRetailSearchbillBillChkInOutDo {
    s.AssEntName = &v
    return s
}
