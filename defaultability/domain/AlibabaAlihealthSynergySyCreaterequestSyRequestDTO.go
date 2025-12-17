package domain


type AlibabaAlihealthSynergySyCreaterequestSyRequestDTO struct {
    /*
        企业（本企业id）     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        被索取企业, 逗号隔开。  不支持同时索取企业资料和产品资料，需要分别调用。索取产品资料时只填一个企业ID。     */
    ReceiveEntId  *string `json:"receive_ent_id,omitempty" `

    /*
        被索取产品ID, 逗号隔开     */
    ReceiveProductId  *string `json:"receive_product_id,omitempty" `

    /*
        索取理由     */
    RequestNote  *string `json:"request_note,omitempty" `

    /*
        索取人     */
    Operator  *string `json:"operator,omitempty" `

    /*
        0:企业资料 1：产品资料     */
    Type  *int64 `json:"type,omitempty" `

}

func (s *AlibabaAlihealthSynergySyCreaterequestSyRequestDTO) SetRefEntId(v string) *AlibabaAlihealthSynergySyCreaterequestSyRequestDTO {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyCreaterequestSyRequestDTO) SetReceiveEntId(v string) *AlibabaAlihealthSynergySyCreaterequestSyRequestDTO {
    s.ReceiveEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyCreaterequestSyRequestDTO) SetReceiveProductId(v string) *AlibabaAlihealthSynergySyCreaterequestSyRequestDTO {
    s.ReceiveProductId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyCreaterequestSyRequestDTO) SetRequestNote(v string) *AlibabaAlihealthSynergySyCreaterequestSyRequestDTO {
    s.RequestNote = &v
    return s
}
func (s *AlibabaAlihealthSynergySyCreaterequestSyRequestDTO) SetOperator(v string) *AlibabaAlihealthSynergySyCreaterequestSyRequestDTO {
    s.Operator = &v
    return s
}
func (s *AlibabaAlihealthSynergySyCreaterequestSyRequestDTO) SetType(v int64) *AlibabaAlihealthSynergySyCreaterequestSyRequestDTO {
    s.Type = &v
    return s
}
