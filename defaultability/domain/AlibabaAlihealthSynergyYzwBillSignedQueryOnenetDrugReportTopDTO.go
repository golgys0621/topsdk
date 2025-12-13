package domain


type AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO struct {
    /*
        单据类型     */
    BillType  *int64 `json:"bill_type,omitempty" `

    /*
        单据id     */
    BillId  *string `json:"bill_id,omitempty" `

    /*
        单据明细ID     */
    BillDetailId  *string `json:"bill_detail_id,omitempty" `

    /*
        单据时间     */
    BillTime  *string `json:"bill_time,omitempty" `

    /*
        单据编码     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        批号     */
    ProduceBatchNo  *string `json:"produce_batch_no,omitempty" `

    /*
        药品id     */
    DrugId  *string `json:"drug_id,omitempty" `

    /*
        单据报告url     */
    SealedReportUrl  *string `json:"sealed_report_url,omitempty" `

    /*
        报告盖章状态:0未盖章，5手动盖章，6同批次盖章     */
    ReportSealStatus  *string `json:"report_seal_status,omitempty" `

    /*
        药品通用名     */
    PhysicName  *string `json:"physic_name,omitempty" `

    /*
        上市许可持有人     */
    MahName  *string `json:"mah_name,omitempty" `

    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        包装规格     */
    PkgSpec  *string `json:"pkg_spec,omitempty" `

    /*
        批准文号     */
    ApprovalNo  *string `json:"approval_no,omitempty" `

    /*
        上游企业名称     */
    FromEntName  *string `json:"from_ent_name,omitempty" `

    /*
        签收时间     */
    SignedTime  *string `json:"signed_time,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetBillType(v int64) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetBillId(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.BillId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetBillDetailId(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.BillDetailId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetBillTime(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetBillCode(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetProduceBatchNo(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetDrugId(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.DrugId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetSealedReportUrl(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.SealedReportUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetReportSealStatus(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.ReportSealStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetPhysicName(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetMahName(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.MahName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetPrepnSpec(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetPkgSpec(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetApprovalNo(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.ApprovalNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetFromEntName(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.FromEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO) SetSignedTime(v string) *AlibabaAlihealthSynergyYzwBillSignedQueryOnenetDrugReportTopDTO {
    s.SignedTime = &v
    return s
}
