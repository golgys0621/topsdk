package domain


type AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse struct {
    /*
        产品id     */
    ProductId  *int64 `json:"product_id,omitempty" `

    /*
        产品名称     */
    ProductName  *string `json:"product_name,omitempty" `

    /*
        产品大类     */
    ProductTypeDesc  *string `json:"product_type_desc,omitempty" `

    /*
        批准文号/注册证号     */
    ApprovalOrRegisterNo  *string `json:"approval_or_register_no,omitempty" `

    /*
        供应商名字     */
    SupplyEntName  *string `json:"supply_ent_name,omitempty" `

    /*
        产品所属资料的总页数     */
    ResourceCount  *int64 `json:"resource_count,omitempty" `

    /*
        产品资料集合     */
    ProductResourcesList  *[]AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse `json:"product_resources_list,omitempty" `

}

func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse) SetProductId(v int64) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse {
    s.ProductId = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse) SetProductName(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse {
    s.ProductName = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse) SetProductTypeDesc(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse {
    s.ProductTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse) SetApprovalOrRegisterNo(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse {
    s.ApprovalOrRegisterNo = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse) SetSupplyEntName(v string) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse {
    s.SupplyEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse) SetResourceCount(v int64) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse {
    s.ResourceCount = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse) SetProductResourcesList(v []AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse {
    s.ProductResourcesList = &v
    return s
}
