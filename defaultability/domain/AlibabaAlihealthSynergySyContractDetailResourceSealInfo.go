package domain


type AlibabaAlihealthSynergySyContractDetailResourceSealInfo struct {
    /*
        数量     */
    Mode  *int64 `json:"mode,omitempty" `

    /*
        印章信息列表     */
    SealInfoList  *[]AlibabaAlihealthSynergySyContractDetailSealInfo `json:"seal_info_list,omitempty" `

    /*
        1     */
    Model  *int64 `json:"model,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractDetailResourceSealInfo) SetMode(v int64) *AlibabaAlihealthSynergySyContractDetailResourceSealInfo {
    s.Mode = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailResourceSealInfo) SetSealInfoList(v []AlibabaAlihealthSynergySyContractDetailSealInfo) *AlibabaAlihealthSynergySyContractDetailResourceSealInfo {
    s.SealInfoList = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractDetailResourceSealInfo) SetModel(v int64) *AlibabaAlihealthSynergySyContractDetailResourceSealInfo {
    s.Model = &v
    return s
}
