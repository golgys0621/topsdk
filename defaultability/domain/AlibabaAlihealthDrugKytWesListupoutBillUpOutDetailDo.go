package domain


type AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo struct {
    /*
        发货单位     */
    FromEntName  *string `json:"from_ent_name,omitempty" `

    /*
        最小码量     */
    CodeCount  *int64 `json:"code_count,omitempty" `

    /*
        失效日期     */
    ExprieDate  *string `json:"exprie_date,omitempty" `

    /*
        厂商     */
    ProduceEntName  *string `json:"produce_ent_name,omitempty" `

    /*
        生产日期     */
    ProduceDate  *string `json:"produce_date,omitempty" `

    /*
        生产批号     */
    ProduceBatchNo  *string `json:"produce_batch_no,omitempty" `

    /*
        包装规格     */
    PkgSpec  *string `json:"pkg_spec,omitempty" `

    /*
        药品信息     */
    PhysicInfo  *string `json:"physic_info,omitempty" `

    /*
        药品名称     */
    PhysicName  *string `json:"physic_name,omitempty" `

    /*
        制剂数量     */
    PrepnCount  *int64 `json:"prepn_count,omitempty" `

    /*
        发货单位REF_ENT_ID     */
    FromRefUserId  *string `json:"from_ref_user_id,omitempty" `

    /*
        收货单位REF_ENT_ID     */
    ToRefUserId  *string `json:"to_ref_user_id,omitempty" `

    /*
        单据时间     */
    BillTime  *string `json:"bill_time,omitempty" `

    /*
        单据码     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        单据类型     */
    BillType  *string `json:"bill_type,omitempty" `

    /*
        收货企业     */
    ToUserName  *string `json:"to_user_name,omitempty" `

    /*
        发货企业     */
    FromUserName  *string `json:"from_user_name,omitempty" `

    /*
        失效日期格式化     */
    ExprieDateFormat  *string `json:"exprie_date_format,omitempty" `

    /*
        单据时间格式化     */
    BillTimeFormat  *string `json:"bill_time_format,omitempty" `

    /*
        单据ID     */
    BillOutId  *int64 `json:"bill_out_id,omitempty" `

    /*
        制剂单位     */
    PrepnUnit  *int64 `json:"prepn_unit,omitempty" `

    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        药品ID     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" `

    /*
        生产日期格式化     */
    ProduceDateFormat  *string `json:"produce_date_format,omitempty" `

    /*
        确认状态1未确认2已确认     */
    Status  *string `json:"status,omitempty" `

    /*
        收货企业ent_id     */
    ToUserId  *string `json:"to_user_id,omitempty" `

    /*
        发货企业ent_id     */
    FromUserId  *string `json:"from_user_id,omitempty" `

    /*
        单据明细ID     */
    BillDetailId  *string `json:"bill_detail_id,omitempty" `

    /*
        委托企业名称     */
    AssEntName  *string `json:"ass_ent_name,omitempty" `

    /*
        委托企业refEntId     */
    AssRefEntId  *string `json:"ass_ref_ent_id,omitempty" `

    /*
        委托企业entId     */
    AssEntId  *string `json:"ass_ent_id,omitempty" `

    /*
        批准文号     */
    ApprovalLicenceNo  *string `json:"approval_licence_no,omitempty" `

    /*
        制剂单位规格描述     */
    PrepnTypeDesc  *string `json:"prepn_type_desc,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetFromEntName(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.FromEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetCodeCount(v int64) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.CodeCount = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetExprieDate(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.ExprieDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetProduceEntName(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetProduceDate(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetProduceBatchNo(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetPkgSpec(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetPhysicInfo(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.PhysicInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetPhysicName(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetPrepnCount(v int64) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.PrepnCount = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetFromRefUserId(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetToRefUserId(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetBillTime(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetBillCode(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetBillType(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetToUserName(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetFromUserName(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetExprieDateFormat(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.ExprieDateFormat = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetBillTimeFormat(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.BillTimeFormat = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetBillOutId(v int64) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.BillOutId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetPrepnUnit(v int64) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.PrepnUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetPrepnSpec(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetProduceDateFormat(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.ProduceDateFormat = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetStatus(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetToUserId(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetFromUserId(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetBillDetailId(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.BillDetailId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetAssEntName(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.AssEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetAssRefEntId(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.AssRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetAssEntId(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.AssEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetApprovalLicenceNo(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.ApprovalLicenceNo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo) SetPrepnTypeDesc(v string) *AlibabaAlihealthDrugKytWesListupoutBillUpOutDetailDo {
    s.PrepnTypeDesc = &v
    return s
}
