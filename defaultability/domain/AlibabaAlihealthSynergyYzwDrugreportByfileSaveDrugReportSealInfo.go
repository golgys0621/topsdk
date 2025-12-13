package domain


type AlibabaAlihealthSynergyYzwDrugreportByfileSaveDrugReportSealInfo struct {
    /*
        Y坐标，不传值时自动找位置盖章     */
    PositionY  *int64 `json:"position_y,omitempty" `

    /*
        调用企业的印章名称,为空表示不盖章     */
    SealName  *string `json:"seal_name,omitempty" `

    /*
        X坐标，不传值时自动找位置盖章     */
    PositionX  *int64 `json:"position_x,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveDrugReportSealInfo) SetPositionY(v int64) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveDrugReportSealInfo {
    s.PositionY = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveDrugReportSealInfo) SetSealName(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveDrugReportSealInfo {
    s.SealName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveDrugReportSealInfo) SetPositionX(v int64) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveDrugReportSealInfo {
    s.PositionX = &v
    return s
}
