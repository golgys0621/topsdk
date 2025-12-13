package domain


type AlibabaAlihealthSynergyYzwQuerydrugreportassSealSignatureDTO struct {
    /*
        左下角x     */
    PositionX  *int64 `json:"position_x,omitempty" `

    /*
        左下角y     */
    PositionY  *int64 `json:"position_y,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassSealSignatureDTO) SetPositionX(v int64) *AlibabaAlihealthSynergyYzwQuerydrugreportassSealSignatureDTO {
    s.PositionX = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassSealSignatureDTO) SetPositionY(v int64) *AlibabaAlihealthSynergyYzwQuerydrugreportassSealSignatureDTO {
    s.PositionY = &v
    return s
}
