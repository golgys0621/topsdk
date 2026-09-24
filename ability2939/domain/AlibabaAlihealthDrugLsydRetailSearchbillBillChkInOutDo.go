package domain


type AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo struct {
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

func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetBillType(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetBillCode(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetFromUserId(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetRefUserName(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.RefUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetRefUserId(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.RefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetProduceDate(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetUploadFileName(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.UploadFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetFromUserName(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetToUserId(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetProduceEntId(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.ProduceEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetBillTime(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetUserRoleType(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetProcessDate(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetBillId(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.BillId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetToUserName(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetAgentUserName(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.AgentUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetAgentRefUserId(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.AgentRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetBillTypeName(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetToRefUserId(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetFromRefUserId(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo) SetSubProcessFlag(v string) *AlibabaAlihealthDrugLsydRetailSearchbillBillChkInOutDo {
    s.SubProcessFlag = &v
    return s
}
