package domain


type AlibabaAlihealthSynergySyQualificationListSyQualificationInfoList struct {
    /*
        资质编码     */
    Code  *string `json:"code,omitempty" `

    /*
        资质名称     */
    Name  *string `json:"name,omitempty" `

    /*
        该资质下的特殊资料属性，json数组     */
    ExtAttr  *string `json:"ext_attr,omitempty" `

}

func (s *AlibabaAlihealthSynergySyQualificationListSyQualificationInfoList) SetCode(v string) *AlibabaAlihealthSynergySyQualificationListSyQualificationInfoList {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthSynergySyQualificationListSyQualificationInfoList) SetName(v string) *AlibabaAlihealthSynergySyQualificationListSyQualificationInfoList {
    s.Name = &v
    return s
}
func (s *AlibabaAlihealthSynergySyQualificationListSyQualificationInfoList) SetExtAttr(v string) *AlibabaAlihealthSynergySyQualificationListSyQualificationInfoList {
    s.ExtAttr = &v
    return s
}
