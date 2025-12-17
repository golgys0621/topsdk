package domain


type AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetailResponse struct {
    /*
        签章基本信息     */
    BasicInfo  *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail `json:"basic_info,omitempty" `

    /*
        企业资料集合     */
    EntResourceList  *[]AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse `json:"ent_resource_list,omitempty" `

    /*
        产品列表     */
    ProductList  *[]AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse `json:"product_list,omitempty" `

    /*
        委托人列表     */
    AgentPersonList  *[]AlibabaAlihealthSynergySySignatureDetailGetSySignatureAgentPersonResponse `json:"agent_person_list,omitempty" `

}

func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetailResponse) SetBasicInfo(v AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetail) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetailResponse {
    s.BasicInfo = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetailResponse) SetEntResourceList(v []AlibabaAlihealthSynergySySignatureDetailGetSySignatureResourceResponse) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetailResponse {
    s.EntResourceList = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetailResponse) SetProductList(v []AlibabaAlihealthSynergySySignatureDetailGetSySignatureProductResponse) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetailResponse {
    s.ProductList = &v
    return s
}
func (s *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetailResponse) SetAgentPersonList(v []AlibabaAlihealthSynergySySignatureDetailGetSySignatureAgentPersonResponse) *AlibabaAlihealthSynergySySignatureDetailGetSySignatureDetailResponse {
    s.AgentPersonList = &v
    return s
}
