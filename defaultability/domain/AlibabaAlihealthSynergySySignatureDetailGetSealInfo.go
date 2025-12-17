package domain


type AlibabaAlihealthSynergySySignatureDetailGetSealInfo struct {
    /*
        印章类型OFFICIAL:公章,CONTRACT:合同章,LEGALPERSON:法人章,CHARGEPERSON:负责人章,OUTBOUND:出库章,QUALITY:质检章.     */
    SealType  *string `json:"seal_type,omitempty" `

    /*
        预留字段：印章名称，暂时不返回，未来支持     */
    SealName  *string `json:"seal_name,omitempty" `

    /*
        x坐标     */
    PointX  *string `json:"point_x,omitempty" `

    /*
        y坐标     */
    PointY  *string `json:"point_y,omitempty" `

    /*
        印章盖在哪些页上     */
    PageRange  *string `json:"page_range,omitempty" `

}

func (s *AlibabaAlihealthSynergySySignatureDetailGetSealInfo) SetSealType(v string) *AlibabaAlihealthSynergySySignatureDetailGetSealInfo {
    s.SealType = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSealInfo) SetSealName(v string) *AlibabaAlihealthSynergySySignatureDetailGetSealInfo {
    s.SealName = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSealInfo) SetPointX(v string) *AlibabaAlihealthSynergySySignatureDetailGetSealInfo {
    s.PointX = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSealInfo) SetPointY(v string) *AlibabaAlihealthSynergySySignatureDetailGetSealInfo {
    s.PointY = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSealInfo) SetPageRange(v string) *AlibabaAlihealthSynergySySignatureDetailGetSealInfo {
    s.PageRange = &v
    return s
}
