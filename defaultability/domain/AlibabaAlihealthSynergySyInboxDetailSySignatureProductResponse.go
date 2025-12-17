package domain


type AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse struct {
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
    ProductResources  *[]AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse `json:"product_resources,omitempty" `

    /*
        产品详情     */
    ProductDetail  *AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO `json:"product_detail,omitempty" `

}

func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse) SetProductId(v int64) *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse {
    s.ProductId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse) SetProductName(v string) *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse {
    s.ProductName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse) SetProductTypeDesc(v string) *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse {
    s.ProductTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse) SetApprovalOrRegisterNo(v string) *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse {
    s.ApprovalOrRegisterNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse) SetSupplyEntName(v string) *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse {
    s.SupplyEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse) SetResourceCount(v int64) *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse {
    s.ResourceCount = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse) SetProductResources(v []AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse) *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse {
    s.ProductResources = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse) SetProductDetail(v AlibabaAlihealthSynergySyInboxDetailSyProductMapperDTO) *AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse {
    s.ProductDetail = &v
    return s
}
