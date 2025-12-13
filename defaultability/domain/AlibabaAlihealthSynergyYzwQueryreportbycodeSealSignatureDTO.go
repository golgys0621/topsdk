package domain


type AlibabaAlihealthSynergyYzwQueryreportbycodeSealSignatureDTO struct {
    /*
        左下角x     */
    PositionX  *int64 `json:"position_x,omitempty" `

    /*
        左下角y     */
    PositionY  *int64 `json:"position_y,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeSealSignatureDTO) SetPositionX(v int64) *AlibabaAlihealthSynergyYzwQueryreportbycodeSealSignatureDTO {
    s.PositionX = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryreportbycodeSealSignatureDTO) SetPositionY(v int64) *AlibabaAlihealthSynergyYzwQueryreportbycodeSealSignatureDTO {
    s.PositionY = &v
    return s
}
