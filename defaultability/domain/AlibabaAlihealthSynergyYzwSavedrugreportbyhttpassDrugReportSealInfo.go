package domain


type AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassDrugReportSealInfo struct {
    /*
        印章名称     */
    SealName  *string `json:"seal_name,omitempty" `

    /*
        X坐标     */
    PositionX  *int64 `json:"position_x,omitempty" `

    /*
        Y坐标     */
    PositionY  *int64 `json:"position_y,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassDrugReportSealInfo) SetSealName(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassDrugReportSealInfo {
    s.SealName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassDrugReportSealInfo) SetPositionX(v int64) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassDrugReportSealInfo {
    s.PositionX = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassDrugReportSealInfo) SetPositionY(v int64) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassDrugReportSealInfo {
    s.PositionY = &v
    return s
}
