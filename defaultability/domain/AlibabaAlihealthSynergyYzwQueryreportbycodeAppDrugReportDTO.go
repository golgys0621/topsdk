package domain


type AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO struct {
    /*
        报告状态(0 ：待发送  2：待签收 3：已签收 4：已拒绝)     */
    DrugReportSignStatus  *string `json:"drug_report_sign_status,omitempty" `

    /*
        单据id     */
    BillId  *string `json:"bill_id,omitempty" `

    /*
        单据号     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        单据类型     */
    BillType  *int64 `json:"bill_type,omitempty" `

    /*
        单据时间     */
    BillTimeDesc  *string `json:"bill_time_desc,omitempty" `

    /*
        药品id     */
    DrugId  *string `json:"drug_id,omitempty" `

    /*
        批号     */
    ProduceBatchNo  *string `json:"produce_batch_no,omitempty" `

    /*
        药品信息     */
    DrugInfoDesc  *string `json:"drug_info_desc,omitempty" `

    /*
        发货企业     */
    FromRefEntId  *string `json:"from_ref_ent_id,omitempty" `

    /*
        发货企业     */
    FromEntName  *string `json:"from_ent_name,omitempty" `

    /*
        签章报告URL     */
    SealReportUrl  *string `json:"seal_report_url,omitempty" `

    /*
        委托企业     */
    AssEntName  *string `json:"ass_ent_name,omitempty" `

    /*
        生产企业     */
    ProduceEntName  *string `json:"produce_ent_name,omitempty" `

    /*
        批准文号     */
    ApprovalLicenceNo  *string `json:"approval_licence_no,omitempty" `

    /*
        mah     */
    AuthorizedEntName  *string `json:"authorized_ent_name,omitempty" `

    /*
        单据detail     */
    BillDetailId  *string `json:"bill_detail_id,omitempty" `

    /*
        签章坐标     */
    SealSignatures  *[]AlibabaAlihealthSynergyYzwQueryreportbycodeSealSignatureDTO `json:"seal_signatures,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetDrugReportSignStatus(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.DrugReportSignStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetBillId(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.BillId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetBillCode(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetBillType(v int64) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetBillTimeDesc(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.BillTimeDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetDrugId(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.DrugId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetProduceBatchNo(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetDrugInfoDesc(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.DrugInfoDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetFromRefEntId(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.FromRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetFromEntName(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.FromEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetSealReportUrl(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.SealReportUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetAssEntName(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.AssEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetProduceEntName(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetApprovalLicenceNo(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.ApprovalLicenceNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetAuthorizedEntName(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.AuthorizedEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetBillDetailId(v string) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.BillDetailId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO) SetSealSignatures(v []AlibabaAlihealthSynergyYzwQueryreportbycodeSealSignatureDTO) *AlibabaAlihealthSynergyYzwQueryreportbycodeAppDrugReportDTO {
    s.SealSignatures = &v
    return s
}
