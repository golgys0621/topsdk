package request


type AlibabaAlihealthSynergySyOutboxWithdrawRequest struct {
    /*
        本企业refEntId     */
    RefEntId  *string `json:"ref_ent_id,omitempty" required:"false" `
    /*
        发件箱id     */
    OutboxId  *int64 `json:"outbox_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergySyOutboxWithdrawRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyOutboxWithdrawRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxWithdrawRequest) SetOutboxId(v int64) *AlibabaAlihealthSynergySyOutboxWithdrawRequest {
    s.OutboxId = &v
    return s
}

func (req *AlibabaAlihealthSynergySyOutboxWithdrawRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.OutboxId != nil) {
        paramMap["outbox_id"] = *req.OutboxId
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyOutboxWithdrawRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}