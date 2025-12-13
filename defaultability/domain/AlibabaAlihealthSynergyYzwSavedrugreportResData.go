package domain


type AlibabaAlihealthSynergyYzwSavedrugreportResData struct {
    /*
        报告id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        盖章报告链接     */
    SealedReportUrl  *string `json:"sealed_report_url,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwSavedrugreportResData) SetId(v int64) *AlibabaAlihealthSynergyYzwSavedrugreportResData {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportResData) SetSealedReportUrl(v string) *AlibabaAlihealthSynergyYzwSavedrugreportResData {
    s.SealedReportUrl = &v
    return s
}
