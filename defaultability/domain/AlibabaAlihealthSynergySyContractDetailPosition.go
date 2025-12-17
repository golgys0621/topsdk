package domain


type AlibabaAlihealthSynergySyContractDetailPosition struct {
    /*
        x轴     */
    X  *string `json:"x,omitempty" `

    /*
        y轴     */
    Y  *string `json:"y,omitempty" `

    /*
        盖章页范围     */
    PageRange  *string `json:"page_range,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractDetailPosition) SetX(v string) *AlibabaAlihealthSynergySyContractDetailPosition {
    s.X = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailPosition) SetY(v string) *AlibabaAlihealthSynergySyContractDetailPosition {
    s.Y = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailPosition) SetPageRange(v string) *AlibabaAlihealthSynergySyContractDetailPosition {
    s.PageRange = &v
    return s
}
