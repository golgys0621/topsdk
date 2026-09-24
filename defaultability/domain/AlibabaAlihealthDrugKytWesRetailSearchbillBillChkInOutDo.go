package domain


type AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo struct {
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

func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetBillType(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetBillCode(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetFromUserId(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetRefUserName(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.RefUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetRefUserId(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.RefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetProduceDate(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetUploadFileName(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.UploadFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetFromUserName(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetToUserId(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetProduceEntId(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.ProduceEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetBillTime(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetUserRoleType(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetProcessDate(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetBillId(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.BillId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetToUserName(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetAgentUserName(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.AgentUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetAgentRefUserId(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.AgentRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetBillTypeName(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetToRefUserId(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetFromRefUserId(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetCrtDate(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetAssEntName(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.AssEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetAssEntId(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.AssEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetAssRefEntId(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.AssRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo) SetSubProcessFlag(v string) *AlibabaAlihealthDrugKytWesRetailSearchbillBillChkInOutDo {
    s.SubProcessFlag = &v
    return s
}
