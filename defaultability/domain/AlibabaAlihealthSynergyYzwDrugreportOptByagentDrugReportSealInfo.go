package domain


type AlibabaAlihealthSynergyYzwDrugreportOptByagentDrugReportSealInfo struct {
    /*
        Y坐标，不传值时自动找位置盖章     */
    PositionY  *int64 `json:"position_y,omitempty" `

    /*
        印章名称，当操作类型为signAndSeal和sealAfterSign时，必填     */
    SealName  *string `json:"seal_name,omitempty" `

    /*
        X坐标，不传值时自动找位置盖章     */
    PositionX  *int64 `json:"position_x,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentDrugReportSealInfo) SetPositionY(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptByagentDrugReportSealInfo {
    s.PositionY = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentDrugReportSealInfo) SetSealName(v string) *AlibabaAlihealthSynergyYzwDrugreportOptByagentDrugReportSealInfo {
    s.SealName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentDrugReportSealInfo) SetPositionX(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptByagentDrugReportSealInfo {
    s.PositionX = &v
    return s
}
