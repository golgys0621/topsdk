package request


type AlibabaAlihealthSynergySyInboxListRequest struct {
    /*
        发送企业名称     */
    SendEntName  *string `json:"send_ent_name,omitempty" required:"false" `
    /*
        301发送方已撤回,302接收方待查收,303接收方审批中,304接收方全部查收,305接收方部分查收,306接收方全部拒收     */
    ExchangeStatus  *int64 `json:"exchange_status,omitempty" required:"false" `
    /*
        发送起始时间     */
    SendStartTime  *string `json:"send_start_time,omitempty" required:"false" `
    /*
        本企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        0企业资料，1产品资料     */
    Type  *int64 `json:"type,omitempty" required:"false" `
    /*
        发送企业refEntId     */
    SendRefEntId  *string `json:"send_ref_ent_id,omitempty" required:"false" `
    /*
        发送结束时间     */
    SendEndTime  *string `json:"send_end_time,omitempty" required:"false" `
    /*
        页     */
    Page  *int64 `json:"page" required:"true" `
    /*
        页     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthSynergySyInboxListRequest) SetSendEntName(v string) *AlibabaAlihealthSynergySyInboxListRequest {
    s.SendEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxListRequest) SetExchangeStatus(v int64) *AlibabaAlihealthSynergySyInboxListRequest {
    s.ExchangeStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxListRequest) SetSendStartTime(v string) *AlibabaAlihealthSynergySyInboxListRequest {
    s.SendStartTime = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxListRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyInboxListRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxListRequest) SetType(v int64) *AlibabaAlihealthSynergySyInboxListRequest {
    s.Type = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxListRequest) SetSendRefEntId(v string) *AlibabaAlihealthSynergySyInboxListRequest {
    s.SendRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxListRequest) SetSendEndTime(v string) *AlibabaAlihealthSynergySyInboxListRequest {
    s.SendEndTime = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxListRequest) SetPage(v int64) *AlibabaAlihealthSynergySyInboxListRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxListRequest) SetPageSize(v int64) *AlibabaAlihealthSynergySyInboxListRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthSynergySyInboxListRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SendEntName != nil) {
        paramMap["send_ent_name"] = *req.SendEntName
    }
    if(req.ExchangeStatus != nil) {
        paramMap["exchange_status"] = *req.ExchangeStatus
    }
    if(req.SendStartTime != nil) {
        paramMap["send_start_time"] = *req.SendStartTime
    }
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.SendRefEntId != nil) {
        paramMap["send_ref_ent_id"] = *req.SendRefEntId
    }
    if(req.SendEndTime != nil) {
        paramMap["send_end_time"] = *req.SendEndTime
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyInboxListRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}