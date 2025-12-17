package domain


type AlibabaAlihealthSynergySySignatureResourceSealSealInfo struct {
    /*
        印章名称     */
    SealName  *string `json:"seal_name,omitempty" `

    /*
        盖章范围，all所有页都盖     */
    PageRange  *string `json:"page_range,omitempty" `

    /*
        印章x坐标     */
    PointX  *string `json:"point_x,omitempty" `

    /*
        印章y坐标     */
    PointY  *string `json:"point_y,omitempty" `

}

func (s *AlibabaAlihealthSynergySySignatureResourceSealSealInfo) SetSealName(v string) *AlibabaAlihealthSynergySySignatureResourceSealSealInfo {
    s.SealName = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureResourceSealSealInfo) SetPageRange(v string) *AlibabaAlihealthSynergySySignatureResourceSealSealInfo {
    s.PageRange = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureResourceSealSealInfo) SetPointX(v string) *AlibabaAlihealthSynergySySignatureResourceSealSealInfo {
    s.PointX = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureResourceSealSealInfo) SetPointY(v string) *AlibabaAlihealthSynergySySignatureResourceSealSealInfo {
    s.PointY = &v
    return s
}
