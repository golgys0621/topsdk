package domain


type AlibabaAlihealthSynergySyContractCreateSealInfo struct {
    /*
        印章类型     */
    SealType  *string `json:"seal_type,omitempty" `

    /*
        盖章范围，全部盖章填all。1. 零散页码以英文逗号分隔，例如输入：1,3,11 2. 页码范围以英文短横线连接，例如输入：5-9 3. 上述两种方式可以组合使用, 例如输入：1,3,5-9,11     */
    PageRange  *string `json:"page_range,omitempty" `

    /*
        盖章位置     */
    PointX  *string `json:"point_x,omitempty" `

    /*
        盖章位置     */
    PointY  *string `json:"point_y,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractCreateSealInfo) SetSealType(v string) *AlibabaAlihealthSynergySyContractCreateSealInfo {
    s.SealType = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateSealInfo) SetPageRange(v string) *AlibabaAlihealthSynergySyContractCreateSealInfo {
    s.PageRange = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateSealInfo) SetPointX(v string) *AlibabaAlihealthSynergySyContractCreateSealInfo {
    s.PointX = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractCreateSealInfo) SetPointY(v string) *AlibabaAlihealthSynergySyContractCreateSealInfo {
    s.PointY = &v
    return s
}
