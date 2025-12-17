package domain


type AlibabaAlihealthSynergySyOutboxDetailSyOutboxDetailResponse struct {
    /*
        企业资料列表     */
    EntResources  *[]AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse `json:"ent_resources,omitempty" `

    /*
        发件箱基本信息     */
    OutboxDetail  *AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel `json:"outbox_detail,omitempty" `

    /*
        产品列表     */
    Products  *[]AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse `json:"products,omitempty" `

    /*
        委托人列表     */
    AgentPersons  *[]AlibabaAlihealthSynergySyOutboxDetailSySignatureAgentPersonResponse `json:"agent_persons,omitempty" `

}

func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxDetailResponse) SetEntResources(v []AlibabaAlihealthSynergySyOutboxDetailSySignatureResourceResponse) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxDetailResponse {
    s.EntResources = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxDetailResponse) SetOutboxDetail(v AlibabaAlihealthSynergySyOutboxDetailSyOutboxModel) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxDetailResponse {
    s.OutboxDetail = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxDetailResponse) SetProducts(v []AlibabaAlihealthSynergySyOutboxDetailSySignatureProductResponse) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxDetailResponse {
    s.Products = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailSyOutboxDetailResponse) SetAgentPersons(v []AlibabaAlihealthSynergySyOutboxDetailSySignatureAgentPersonResponse) *AlibabaAlihealthSynergySyOutboxDetailSyOutboxDetailResponse {
    s.AgentPersons = &v
    return s
}
