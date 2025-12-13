package domain


type AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO struct {
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
        报告id     */
    ReportId  *string `json:"report_id,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO) SetDrugReportV2Id(v int64) *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO {
    s.DrugReportV2Id = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO) SetDrugName(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO {
    s.DrugName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO) SetPkgSpec(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO) SetPrepnSpec(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO) SetProdCode(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO) SetPkgRatioList(v []string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO {
    s.PkgRatioList = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO) SetDrugReportName(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO {
    s.DrugReportName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO {
    s.BatchNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO) SetSealedReportUrl(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO {
    s.SealedReportUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO) SetReportNo(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO {
    s.ReportNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO) SetReportDate(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO {
    s.ReportDate = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO) SetReportId(v string) *AlibabaAlihealthSynergyYzwQuerysealdrugreportOcrSealDrugReportDTO {
    s.ReportId = &v
    return s
}
