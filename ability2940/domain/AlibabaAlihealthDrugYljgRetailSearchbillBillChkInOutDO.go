package domain


type AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO struct {
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

}

func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetBillType(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetBillCode(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetFromUserId(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetRefUserName(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.RefUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetRefUserId(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.RefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetProduceDate(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetUploadFileName(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.UploadFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetFromUserName(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetToUserId(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetProduceEntId(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.ProduceEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetBillTime(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetUserRoleType(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetProcessDate(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetBillId(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.BillId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetToUserName(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetAgentUserName(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.AgentUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetAgentRefUserId(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.AgentRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetBillTypeName(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetToRefUserId(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetFromRefUserId(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO) SetSubProcessFlag(v string) *AlibabaAlihealthDrugYljgRetailSearchbillBillChkInOutDO {
    s.SubProcessFlag = &v
    return s
}
