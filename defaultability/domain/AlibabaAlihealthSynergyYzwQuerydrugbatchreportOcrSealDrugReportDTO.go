package domain


type AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO struct {
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
        包装比例     */
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
        报告编号     */
    ReportNo  *string `json:"report_no,omitempty" `

    /*
        报告日期     */
    ReportDate  *string `json:"report_date,omitempty" `

    /*
        盖章状态:1是未盖章,2是已盖章     */
    Status  *string `json:"status,omitempty" `

    /*
        报告id     */
    ReportId  *string `json:"report_id,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO) SetDrugReportV2Id(v int64) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO {
    s.DrugReportV2Id = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO) SetDrugName(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO {
    s.DrugName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO) SetPkgSpec(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO) SetPrepnSpec(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO) SetProdCode(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO) SetPkgRatioList(v []string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO {
    s.PkgRatioList = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO) SetDrugReportName(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO {
    s.DrugReportName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO {
    s.BatchNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO) SetSealedReportUrl(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO {
    s.SealedReportUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO) SetReportNo(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO {
    s.ReportNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO) SetReportDate(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO {
    s.ReportDate = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO) SetStatus(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO {
    s.Status = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO) SetReportId(v string) *AlibabaAlihealthSynergyYzwQuerydrugbatchreportOcrSealDrugReportDTO {
    s.ReportId = &v
    return s
}
