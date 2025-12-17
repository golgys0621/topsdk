package domain


type AlibabaAlihealthSynergySyInboxDetailSyInboxDetailResponse struct {
    /*
        企业资料列表     */
    EntResources  *[]AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse `json:"ent_resources,omitempty" `

    /*
        收件箱基本信息     */
    InboxDetail  *AlibabaAlihealthSynergySyInboxDetailSyInboxModel `json:"inbox_detail,omitempty" `

    /*
        产品列表     */
    Products  *[]AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse `json:"products,omitempty" `

    /*
        委托人列表     */
    AgentPersons  *[]AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse `json:"agent_persons,omitempty" `

}

func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxDetailResponse) SetEntResources(v []AlibabaAlihealthSynergySyInboxDetailSySignatureResourceResponse) *AlibabaAlihealthSynergySyInboxDetailSyInboxDetailResponse {
    s.EntResources = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxDetailResponse) SetInboxDetail(v AlibabaAlihealthSynergySyInboxDetailSyInboxModel) *AlibabaAlihealthSynergySyInboxDetailSyInboxDetailResponse {
    s.InboxDetail = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxDetailResponse) SetProducts(v []AlibabaAlihealthSynergySyInboxDetailSySignatureProductResponse) *AlibabaAlihealthSynergySyInboxDetailSyInboxDetailResponse {
    s.Products = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyInboxDetailResponse) SetAgentPersons(v []AlibabaAlihealthSynergySyInboxDetailSySignatureAgentPersonResponse) *AlibabaAlihealthSynergySyInboxDetailSyInboxDetailResponse {
    s.AgentPersons = &v
    return s
}
