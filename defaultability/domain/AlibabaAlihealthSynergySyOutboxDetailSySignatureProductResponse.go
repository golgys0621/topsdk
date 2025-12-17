package domain


type AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse struct {
    /*
        产品id     */
    ProductId  *int64 `json:"product_id,omitempty" `

    /*
        产品名称     */
    ProductName  *string `json:"product_name,omitempty" `

    /*
        产品分类     */
    ProductTypeDesc  *string `json:"product_type_desc,omitempty" `

    /*
        批准文号/注册证号     */
    ApprovalOrRegisterNo  *string `json:"approval_or_register_no,omitempty" `

    /*
        供应商     */
    SupplyEntName  *string `json:"supply_ent_name,omitempty" `

    /*
        资料数     */
    ResourceCount  *int64 `json:"resource_count,omitempty" `

    /*
        产品资料列表     */
    ProductResources  *[]AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse `json:"product_resources,omitempty" `

}

func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse) SetProductId(v int64) *AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse {
    s.ProductId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse) SetProductName(v string) *AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse {
    s.ProductName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse) SetProductTypeDesc(v string) *AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse {
    s.ProductTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse) SetApprovalOrRegisterNo(v string) *AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse {
    s.ApprovalOrRegisterNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse) SetSupplyEntName(v string) *AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse {
    s.SupplyEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse) SetResourceCount(v int64) *AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse {
    s.ResourceCount = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse) SetProductResources(v []AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse) *AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse {
    s.ProductResources = &v
    return s
}
