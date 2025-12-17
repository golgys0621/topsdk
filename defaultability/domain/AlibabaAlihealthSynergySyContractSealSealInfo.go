package domain


type AlibabaAlihealthSynergySyContractSealSealInfo struct {
    /*
        印章位置     */
    PositionList  *[]AlibabaAlihealthSynergySyContractSealPosition `json:"position_list,omitempty" `

    /*
        印章名称     */
    SealName  *string `json:"seal_name,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractSealSealInfo) SetPositionList(v []AlibabaAlihealthSynergySyContractSealPosition) *AlibabaAlihealthSynergySyContractSealSealInfo {
    s.PositionList = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractSealSealInfo) SetSealName(v string) *AlibabaAlihealthSynergySyContractSealSealInfo {
    s.SealName = &v
    return s
}
