package request


type AlibabaAlihealthSynergySyContractListRequest struct {
    /*
        创建时间的开始日期     */
    CreateStartTime  *string `json:"create_start_time,omitempty" required:"false" `
    /*
        创建时间的结束日期     */
    CreateEndTime  *string `json:"create_end_time,omitempty" required:"false" `
    /*
        收发时间的开始日期     */
    SendStartTime  *string `json:"send_start_time,omitempty" required:"false" `
    /*
        收发时间的结束日期     */
    SendEndTime  *string `json:"send_end_time,omitempty" required:"false" `
    /*
        双方签署完成时间的开始日期     */
    CompleteStartTime  *string `json:"complete_start_time,omitempty" required:"false" `
    /*
        双方签署完成时间的结束日期     */
    CompleteEndTime  *string `json:"complete_end_time,omitempty" required:"false" `
    /*
        合同名称     */
    ContractName  *string `json:"contract_name,omitempty" required:"false" `
    /*
        合同类型：0质量保障协议、1销售合同、2采购合同     */
    ContractType  *int64 `json:"contract_type,omitempty" required:"false" `
    /*
        合同状态：1:发起方待选章提交, 2:发起方审批中, 3:发起方审批不通过, 4:发起方待签章发送, 5:发起方签章失败, 6:发起方待发送, 7:接收方待查收, 8:接收方审批中, 9:接收方拒签, 10:接收方待签章, 11:接收方签章失败, 12:双方签署成功, 13:发起方已撤回     */
    ContractStatus  *int64 `json:"contract_status,omitempty" required:"false" `
    /*
        合同来源：0已方发起、1对方发起     */
    Source  *int64 `json:"source,omitempty" required:"false" `
    /*
        发件企业refentid     */
    SendRefEntId  *string `json:"send_ref_ent_id,omitempty" required:"false" `
    /*
        发件企业名称     */
    SendEntName  *string `json:"send_ent_name,omitempty" required:"false" `
    /*
        收件企业refentid     */
    ReceiveRefEntId  *string `json:"receive_ref_ent_id,omitempty" required:"false" `
    /*
        收件企业名称     */
    ReceiveEntName  *string `json:"receive_ent_name,omitempty" required:"false" `
    /*
        往来单位的refentid     */
    PartnerRefEntId  *string `json:"partner_ref_ent_id,omitempty" required:"false" `
    /*
        往来单位的企业id     */
    PartnerEntId  *string `json:"partner_ent_id,omitempty" required:"false" `
    /*
        企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        页     */
    Page  *int64 `json:"page" required:"true" `
    /*
        页     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthSynergySyContractListRequest) SetCreateStartTime(v string) *AlibabaAlihealthSynergySyContractListRequest {
    s.CreateStartTime = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetCreateEndTime(v string) *AlibabaAlihealthSynergySyContractListRequest {
    s.CreateEndTime = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetSendStartTime(v string) *AlibabaAlihealthSynergySyContractListRequest {
    s.SendStartTime = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetSendEndTime(v string) *AlibabaAlihealthSynergySyContractListRequest {
    s.SendEndTime = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetCompleteStartTime(v string) *AlibabaAlihealthSynergySyContractListRequest {
    s.CompleteStartTime = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetCompleteEndTime(v string) *AlibabaAlihealthSynergySyContractListRequest {
    s.CompleteEndTime = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetContractName(v string) *AlibabaAlihealthSynergySyContractListRequest {
    s.ContractName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetContractType(v int64) *AlibabaAlihealthSynergySyContractListRequest {
    s.ContractType = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetContractStatus(v int64) *AlibabaAlihealthSynergySyContractListRequest {
    s.ContractStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetSource(v int64) *AlibabaAlihealthSynergySyContractListRequest {
    s.Source = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetSendRefEntId(v string) *AlibabaAlihealthSynergySyContractListRequest {
    s.SendRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetSendEntName(v string) *AlibabaAlihealthSynergySyContractListRequest {
    s.SendEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetReceiveRefEntId(v string) *AlibabaAlihealthSynergySyContractListRequest {
    s.ReceiveRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetReceiveEntName(v string) *AlibabaAlihealthSynergySyContractListRequest {
    s.ReceiveEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetPartnerRefEntId(v string) *AlibabaAlihealthSynergySyContractListRequest {
    s.PartnerRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetPartnerEntId(v string) *AlibabaAlihealthSynergySyContractListRequest {
    s.PartnerEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyContractListRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetPage(v int64) *AlibabaAlihealthSynergySyContractListRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractListRequest) SetPageSize(v int64) *AlibabaAlihealthSynergySyContractListRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthSynergySyContractListRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CreateStartTime != nil) {
        paramMap["create_start_time"] = *req.CreateStartTime
    }
    if(req.CreateEndTime != nil) {
        paramMap["create_end_time"] = *req.CreateEndTime
    }
    if(req.SendStartTime != nil) {
        paramMap["send_start_time"] = *req.SendStartTime
    }
    if(req.SendEndTime != nil) {
        paramMap["send_end_time"] = *req.SendEndTime
    }
    if(req.CompleteStartTime != nil) {
        paramMap["complete_start_time"] = *req.CompleteStartTime
    }
    if(req.CompleteEndTime != nil) {
        paramMap["complete_end_time"] = *req.CompleteEndTime
    }
    if(req.ContractName != nil) {
        paramMap["contract_name"] = *req.ContractName
    }
    if(req.ContractType != nil) {
        paramMap["contract_type"] = *req.ContractType
    }
    if(req.ContractStatus != nil) {
        paramMap["contract_status"] = *req.ContractStatus
    }
    if(req.Source != nil) {
        paramMap["source"] = *req.Source
    }
    if(req.SendRefEntId != nil) {
        paramMap["send_ref_ent_id"] = *req.SendRefEntId
    }
    if(req.SendEntName != nil) {
        paramMap["send_ent_name"] = *req.SendEntName
    }
    if(req.ReceiveRefEntId != nil) {
        paramMap["receive_ref_ent_id"] = *req.ReceiveRefEntId
    }
    if(req.ReceiveEntName != nil) {
        paramMap["receive_ent_name"] = *req.ReceiveEntName
    }
    if(req.PartnerRefEntId != nil) {
        paramMap["partner_ref_ent_id"] = *req.PartnerRefEntId
    }
    if(req.PartnerEntId != nil) {
        paramMap["partner_ent_id"] = *req.PartnerEntId
    }
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyContractListRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}