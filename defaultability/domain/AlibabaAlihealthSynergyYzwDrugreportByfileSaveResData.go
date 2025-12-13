package domain


type AlibabaAlihealthSynergyYzwDrugreportByfileSaveResData struct {
    /*
        报告id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        盖章报告链接     */
    SealedReportUrl  *string `json:"sealed_report_url,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveResData) SetId(v int64) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveResData {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveResData) SetSealedReportUrl(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveResData {
    s.SealedReportUrl = &v
    return s
}
