package domain


type AlibabaAlihealthSynergySyListreceiverequestSyRequestDetailDTO struct {
    /*
        产品ID     */
    SyProductId  *int64 `json:"sy_product_id,omitempty" `

    /*
        资料名称     */
    ProductName  *string `json:"product_name,omitempty" `

}

func (s *AlibabaAlihealthSynergySyListreceiverequestSyRequestDetailDTO) SetSyProductId(v int64) *AlibabaAlihealthSynergySyListreceiverequestSyRequestDetailDTO {
    s.SyProductId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyRequestDetailDTO) SetProductName(v string) *AlibabaAlihealthSynergySyListreceiverequestSyRequestDetailDTO {
    s.ProductName = &v
    return s
}
