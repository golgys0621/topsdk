package domain


type AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDetailDTO struct {
    /*
        产品ID     */
    SyProductId  *int64 `json:"sy_product_id,omitempty" `

    /*
        产品名称     */
    ProductName  *string `json:"product_name,omitempty" `

}

func (s *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDetailDTO) SetSyProductId(v int64) *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDetailDTO {
    s.SyProductId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDetailDTO) SetProductName(v string) *AlibabaAlihealthSynergySyReceiverequestDetailSyReceiveRequestDetailDTO {
    s.ProductName = &v
    return s
}
