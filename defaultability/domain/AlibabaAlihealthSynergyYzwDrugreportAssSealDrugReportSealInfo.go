package domain


type AlibabaAlihealthSynergyYzwDrugreportAssSealDrugReportSealInfo struct {
    /*
        Y坐标，不传值时自动找位置盖章     */
    PositionY  *int64 `json:"position_y,omitempty" `

    /*
        授权单位的印章名称,为空表示不盖章     */
    SealName  *string `json:"seal_name,omitempty" `

    /*
        X坐标，不传值时自动找位置盖章     */
    PositionX  *int64 `json:"position_x,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportAssSealDrugReportSealInfo) SetPositionY(v int64) *AlibabaAlihealthSynergyYzwDrugreportAssSealDrugReportSealInfo {
    s.PositionY = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportAssSealDrugReportSealInfo) SetSealName(v string) *AlibabaAlihealthSynergyYzwDrugreportAssSealDrugReportSealInfo {
    s.SealName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportAssSealDrugReportSealInfo) SetPositionX(v int64) *AlibabaAlihealthSynergyYzwDrugreportAssSealDrugReportSealInfo {
    s.PositionX = &v
    return s
}
