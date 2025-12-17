package domain


type AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO struct {
    /*
        企业（本企业id）     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        0:企业资料 1：产品资料     */
    Type  *int64 `json:"type,omitempty" `

    /*
        索取企业ID     */
    SendRefEntId  *string `json:"send_ref_ent_id,omitempty" `

    /*
        开始日期     */
    StartDate  *string `json:"start_date,omitempty" `

    /*
        结束日期     */
    EndDate  *string `json:"end_date,omitempty" `

    /*
        索取企业名称     */
    SendEntName  *string `json:"send_ent_name,omitempty" `

    /*
        分页     */
    Page  *int64 `json:"page,omitempty" `

    /*
        分页大小，最大分页为10     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        索取状态     */
    RequestStatus  *string `json:"request_status,omitempty" `

}

func (s *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO) SetRefEntId(v string) *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO) SetType(v int64) *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO {
    s.Type = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO) SetSendRefEntId(v string) *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO {
    s.SendRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO) SetStartDate(v string) *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO {
    s.StartDate = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO) SetEndDate(v string) *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO) SetSendEntName(v string) *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO {
    s.SendEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO) SetPage(v int64) *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO) SetPageSize(v int64) *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO) SetRequestStatus(v string) *AlibabaAlihealthSynergySyListreceiverequestSyRequestDTO {
    s.RequestStatus = &v
    return s
}
