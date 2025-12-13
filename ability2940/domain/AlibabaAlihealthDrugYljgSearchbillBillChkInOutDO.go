package domain


type AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO struct {
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

func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetBillType(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetBillCode(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetFromUserId(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetRefUserName(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.RefUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetRefUserId(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.RefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetProduceDate(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetUploadFileName(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.UploadFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetFromUserName(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetToUserId(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetProduceEntId(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.ProduceEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetBillTime(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetUserRoleType(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.UserRoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetProcessDate(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetBillId(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.BillId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetToUserName(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetAgentUserName(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.AgentUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetAgentRefUserId(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.AgentRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetBillTypeName(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetToRefUserId(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetFromRefUserId(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO) SetSubProcessFlag(v string) *AlibabaAlihealthDrugYljgSearchbillBillChkInOutDO {
    s.SubProcessFlag = &v
    return s
}
