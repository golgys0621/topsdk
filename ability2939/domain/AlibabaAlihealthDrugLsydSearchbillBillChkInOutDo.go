package domain


type AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo struct {
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
        日期     */
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

}

func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetBillType(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetBillCode(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetFromUserId(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetRefUserName(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.RefUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetRefUserId(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.RefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetProduceDate(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetUploadFileName(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.UploadFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetFromUserName(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetToUserId(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetProduceEntId(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.ProduceEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetBillTime(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetUserRoleType(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetProcessDate(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetBillId(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.BillId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetToUserName(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetAgentUserName(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.AgentUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetAgentRefUserId(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.AgentRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetBillTypeName(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetToRefUserId(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetFromRefUserId(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo) SetSubProcessFlag(v string) *AlibabaAlihealthDrugLsydSearchbillBillChkInOutDo {
    s.SubProcessFlag = &v
    return s
}
