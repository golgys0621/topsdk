package domain


type AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResData struct {
    /*
        报告id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        盖章报告链接     */
    SealedReportUrl  *string `json:"sealed_report_url,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResData) SetId(v int64) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResData {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResData) SetSealedReportUrl(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResData {
    s.SealedReportUrl = &v
    return s
}
