package domain


type AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO struct {
    /*
        药品信息     */
    DrugBaseInfo  *string `json:"drug_base_info,omitempty" `

    /*
        产品批号     */
    ProduceBatchNo  *string `json:"produce_batch_no,omitempty" `

    /*
        药检报告上传状态: 0无效; 1有效      */
    Status  *string `json:"status,omitempty" `

    /*
        盖章状态:0无效  1有效     */
    SealStatus  *string `json:"seal_status,omitempty" `

    /*
        批准文号     */
    ApprovalNo  *string `json:"approval_no,omitempty" `

    /*
        生产日期     */
    ProduceDate  *string `json:"produce_date,omitempty" `

    /*
        有效期至     */
    ExpireDate  *string `json:"expire_date,omitempty" `

    /*
        报告创建时间     */
    ReportCreateTime  *string `json:"report_create_time,omitempty" `

    /*
        药品持有人企业名称     */
    AuthEntName  *string `json:"auth_ent_name,omitempty" `

    /*
        生产企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        药检报告id     */
    DrugReportV2Id  *string `json:"drug_report_v2_id,omitempty" `

    /*
        药品id     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" `

    /*
        报告id     */
    ReportId  *string `json:"report_id,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO) SetDrugBaseInfo(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO {
    s.DrugBaseInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO) SetProduceBatchNo(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO) SetStatus(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO) SetSealStatus(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO {
    s.SealStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO) SetApprovalNo(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO {
    s.ApprovalNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO) SetProduceDate(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO) SetExpireDate(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO {
    s.ExpireDate = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO) SetReportCreateTime(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO {
    s.ReportCreateTime = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO) SetAuthEntName(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO {
    s.AuthEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO) SetEntName(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO) SetDrugReportV2Id(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO {
    s.DrugReportV2Id = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO) SetReportId(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoDrugTestReportDTO {
    s.ReportId = &v
    return s
}
