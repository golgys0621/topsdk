package domain


type AlibabaAlihealthSynergyYzwQuerydrugreportSealSignatureDTO struct {
    /*
        左下角x     */
    PositionX  *int64 `json:"position_x,omitempty" `

    /*
        左下角y     */
    PositionY  *int64 `json:"position_y,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQuerydrugreportSealSignatureDTO) SetPositionX(v int64) *AlibabaAlihealthSynergyYzwQuerydrugreportSealSignatureDTO {
    s.PositionX = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportSealSignatureDTO) SetPositionY(v int64) *AlibabaAlihealthSynergyYzwQuerydrugreportSealSignatureDTO {
    s.PositionY = &v
    return s
}
