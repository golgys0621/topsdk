package domain


type AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfDrugReportSealInfo struct {
    /*
        Y坐标，不传值时自动找位置盖章     */
    PositionY  *int64 `json:"position_y,omitempty" `

    /*
        授权企业的印章名称,为空表示不盖章     */
    SealName  *string `json:"seal_name,omitempty" `

    /*
        授权企业的refEntId     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        X坐标，不传值时自动找位置盖章     */
    PositionX  *int64 `json:"position_x,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfDrugReportSealInfo) SetPositionY(v int64) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfDrugReportSealInfo {
    s.PositionY = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfDrugReportSealInfo) SetSealName(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfDrugReportSealInfo {
    s.SealName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfDrugReportSealInfo) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfDrugReportSealInfo {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfDrugReportSealInfo) SetPositionX(v int64) *AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfDrugReportSealInfo {
    s.PositionX = &v
    return s
}
