package domain


type AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO struct {
    /*
        药检报告Id     */
    DrugReportV2Id  *int64 `json:"drug_report_v2_id,omitempty" `

    /*
        药品名称     */
    DrugName  *string `json:"drug_name,omitempty" `

    /*
        包装规格     */
    PkgSpec  *string `json:"pkg_spec,omitempty" `

    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        药品id     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" `

    /*
        药品子类编码     */
    ProdCode  *string `json:"prod_code,omitempty" `

    /*
        包装比例集合     */
    PkgRatioList  *[]string `json:"pkg_ratio_list,omitempty" `

    /*
        报告名称     */
    DrugReportName  *string `json:"drug_report_name,omitempty" `

    /*
        批次号     */
    BatchNo  *string `json:"batch_no,omitempty" `

    /*
        盖章报告链接     */
    SealedReportUrl  *string `json:"sealed_report_url,omitempty" `

    /*
        盖章状态1:未盖章2:已盖章     */
    Status  *int64 `json:"status,omitempty" `

    /*
        生产企业名称     */
    ProduceEntName  *string `json:"produce_ent_name,omitempty" `

    /*
        持有人企业名称     */
    AuthEntName  *string `json:"auth_ent_name,omitempty" `

    /*
        报告编号     */
    ReportNo  *string `json:"report_no,omitempty" `

    /*
        报告日期     */
    ReportDate  *string `json:"report_date,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) SetDrugReportV2Id(v int64) *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO {
    s.DrugReportV2Id = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) SetDrugName(v string) *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO {
    s.DrugName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) SetPkgSpec(v string) *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) SetPrepnSpec(v string) *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) SetProdCode(v string) *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) SetPkgRatioList(v []string) *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO {
    s.PkgRatioList = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) SetDrugReportName(v string) *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO {
    s.DrugReportName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO {
    s.BatchNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) SetSealedReportUrl(v string) *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO {
    s.SealedReportUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) SetStatus(v int64) *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) SetProduceEntName(v string) *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) SetAuthEntName(v string) *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO {
    s.AuthEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) SetReportNo(v string) *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO {
    s.ReportNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO) SetReportDate(v string) *AlibabaAlihealthSynergyYzwDrugreportGetOcrSealDrugReportDTO {
    s.ReportDate = &v
    return s
}
