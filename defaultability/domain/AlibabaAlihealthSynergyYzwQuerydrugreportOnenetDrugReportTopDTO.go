package domain


type AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO struct {
    /*
        单据类型     */
    BillType  *int64 `json:"bill_type,omitempty" `

    /*
        单据ID     */
    BillId  *string `json:"bill_id,omitempty" `

    /*
        单据明细ID     */
    BillDetailId  *string `json:"bill_detail_id,omitempty" `

    /*
        单据时间     */
    BillTime  *string `json:"bill_time,omitempty" `

    /*
        包装     */
    PkgSpec  *string `json:"pkg_spec,omitempty" `

    /*
        制剂     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        药品名     */
    PhysicName  *string `json:"physic_name,omitempty" `

    /*
        剂型     */
    PrepnTypeDesc  *string `json:"prepn_type_desc,omitempty" `

    /*
        报告状态(0 ：待发送  2：待签收 3：已签收 4：已拒绝 7:对方已签收(更正待处理)  13:对方已拒绝(更正待签收))     */
    DrugReportSignStatus  *string `json:"drug_report_sign_status,omitempty" `

    /*
        单据编码     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        批号     */
    ProduceBatchNo  *string `json:"produce_batch_no,omitempty" `

    /*
        药品ID     */
    DrugId  *string `json:"drug_id,omitempty" `

    /*
        生产日期     */
    ProduceDate  *string `json:"produce_date,omitempty" `

    /*
        生产企业     */
    ProduceEntId  *string `json:"produce_ent_id,omitempty" `

    /*
        委托企业     */
    AssRefEntId  *string `json:"ass_ref_ent_id,omitempty" `

    /*
        报告名称     */
    FileName  *string `json:"file_name,omitempty" `

    /*
        签章位置     */
    SealSignatures  *[]AlibabaAlihealthSynergyYzwQuerydrugreportSealSignatureDTO `json:"seal_signatures,omitempty" `

    /*
        签章URL     */
    SealedReportUrl  *string `json:"sealed_report_url,omitempty" `

    /*
        报告id     */
    DrugReportId  *int64 `json:"drug_report_id,omitempty" `

    /*
        生产企业     */
    ProduceEntName  *string `json:"produce_ent_name,omitempty" `

    /*
        发货企业     */
    FromRefEntId  *string `json:"from_ref_ent_id,omitempty" `

    /*
        发货企业     */
    FromEntName  *string `json:"from_ent_name,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetBillType(v int64) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetBillId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.BillId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetBillDetailId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.BillDetailId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetBillTime(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetPkgSpec(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetPrepnSpec(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetPhysicName(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetPrepnTypeDesc(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.PrepnTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetDrugReportSignStatus(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.DrugReportSignStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetBillCode(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetProduceBatchNo(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetDrugId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.DrugId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetProduceDate(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetProduceEntId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.ProduceEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetAssRefEntId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.AssRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetFileName(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.FileName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetSealSignatures(v []AlibabaAlihealthSynergyYzwQuerydrugreportSealSignatureDTO) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.SealSignatures = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetSealedReportUrl(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.SealedReportUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetDrugReportId(v int64) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.DrugReportId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetProduceEntName(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetFromRefEntId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.FromRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO) SetFromEntName(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportOnenetDrugReportTopDTO {
    s.FromEntName = &v
    return s
}
