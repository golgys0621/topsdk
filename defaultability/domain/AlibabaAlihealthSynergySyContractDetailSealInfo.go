package domain


type AlibabaAlihealthSynergySyContractDetailSealInfo struct {
    /*
        印章类型     */
    SealType  *string `json:"seal_type,omitempty" `

    /*
        印章类型名称     */
    SealTypeName  *string `json:"seal_type_name,omitempty" `

    /*
        印章名称     */
    SealName  *string `json:"seal_name,omitempty" `

    /*
        印章描述     */
    SealDesc  *string `json:"seal_desc,omitempty" `

    /*
        统一位置，优先看自定义位置     */
    PositionList  *[]AlibabaAlihealthSynergySyContractDetailPosition `json:"position_list,omitempty" `

    /*
        自定义位置     */
    CustomPositionList  *[]AlibabaAlihealthSynergySyContractDetailPosition `json:"custom_position_list,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractDetailSealInfo) SetSealType(v string) *AlibabaAlihealthSynergySyContractDetailSealInfo {
    s.SealType = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSealInfo) SetSealTypeName(v string) *AlibabaAlihealthSynergySyContractDetailSealInfo {
    s.SealTypeName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSealInfo) SetSealName(v string) *AlibabaAlihealthSynergySyContractDetailSealInfo {
    s.SealName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSealInfo) SetSealDesc(v string) *AlibabaAlihealthSynergySyContractDetailSealInfo {
    s.SealDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSealInfo) SetPositionList(v []AlibabaAlihealthSynergySyContractDetailPosition) *AlibabaAlihealthSynergySyContractDetailSealInfo {
    s.PositionList = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSealInfo) SetCustomPositionList(v []AlibabaAlihealthSynergySyContractDetailPosition) *AlibabaAlihealthSynergySyContractDetailSealInfo {
    s.CustomPositionList = &v
    return s
}
