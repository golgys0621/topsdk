package request


type AlibabaAlihealthSynergySyOutboxListRequest struct {
    /*
        接收企业refentid     */
    ReceiveRefEntId  *string `json:"receive_ref_ent_id,omitempty" required:"false" `
    /*
        301发送方已撤回,302接收方待查收,303接收方审批中,304接收方全部查收,305接收方部分查收,306接收方全部拒收     */
    ExchangeStatus  *string `json:"exchange_status,omitempty" required:"false" `
    /*
        发送起始时间     */
    SendStartTime  *string `json:"send_start_time,omitempty" required:"false" `
    /*
        本企业refentid     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        0企业资料，1产品资料     */
    Type  *int64 `json:"type,omitempty" required:"false" `
    /*
        接收企业名称     */
    ReceiveEntName  *string `json:"receive_ent_name,omitempty" required:"false" `
    /*
        发送结束时间     */
    SendEndTime  *string `json:"send_end_time,omitempty" required:"false" `
    /*
        签章id     */
    SySignatureId  *int64 `json:"sy_signature_id,omitempty" required:"false" `
    /*
        页     */
    Page  *int64 `json:"page" required:"true" `
    /*
        页大小     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthSynergySyOutboxListRequest) SetReceiveRefEntId(v string) *AlibabaAlihealthSynergySyOutboxListRequest {
    s.ReceiveRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListRequest) SetExchangeStatus(v string) *AlibabaAlihealthSynergySyOutboxListRequest {
    s.ExchangeStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListRequest) SetSendStartTime(v string) *AlibabaAlihealthSynergySyOutboxListRequest {
    s.SendStartTime = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyOutboxListRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListRequest) SetType(v int64) *AlibabaAlihealthSynergySyOutboxListRequest {
    s.Type = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListRequest) SetReceiveEntName(v string) *AlibabaAlihealthSynergySyOutboxListRequest {
    s.ReceiveEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListRequest) SetSendEndTime(v string) *AlibabaAlihealthSynergySyOutboxListRequest {
    s.SendEndTime = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListRequest) SetSySignatureId(v int64) *AlibabaAlihealthSynergySyOutboxListRequest {
    s.SySignatureId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListRequest) SetPage(v int64) *AlibabaAlihealthSynergySyOutboxListRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxListRequest) SetPageSize(v int64) *AlibabaAlihealthSynergySyOutboxListRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthSynergySyOutboxListRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ReceiveRefEntId != nil) {
        paramMap["receive_ref_ent_id"] = *req.ReceiveRefEntId
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
    if(req.ReceiveEntName != nil) {
        paramMap["receive_ent_name"] = *req.ReceiveEntName
    }
    if(req.SendEndTime != nil) {
        paramMap["send_end_time"] = *req.SendEndTime
    }
    if(req.SySignatureId != nil) {
        paramMap["sy_signature_id"] = *req.SySignatureId
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyOutboxListRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}