package domain


type AlibabaAlihealthSynergyYzwDrugreportOptEntDealResultDTO struct {
    /*
        当操作类型为签收并盖章(signAndSeal)和盖章签收后的报告(sealAfterSign)，会返回盖章的报告id     */
    DrugReportId  *int64 `json:"drug_report_id,omitempty" `

    /*
        操作类型：签收(onlySign),拒收(rejectReceive),签收并盖章(signAndSeal),盖章签收后的报告(sealAfterSign)     */
    Opt  *string `json:"opt,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportOptEntDealResultDTO) SetDrugReportId(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptEntDealResultDTO {
    s.DrugReportId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptEntDealResultDTO) SetOpt(v string) *AlibabaAlihealthSynergyYzwDrugreportOptEntDealResultDTO {
    s.Opt = &v
    return s
}
