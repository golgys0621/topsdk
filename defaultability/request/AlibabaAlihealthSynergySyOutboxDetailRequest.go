package request


type AlibabaAlihealthSynergySyOutboxDetailRequest struct {
    /*
        本企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        发件箱id     */
    OutboxId  *int64 `json:"outbox_id" required:"true" `
}

func (s *AlibabaAlihealthSynergySyOutboxDetailRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyOutboxDetailRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyOutboxDetailRequest) SetOutboxId(v int64) *AlibabaAlihealthSynergySyOutboxDetailRequest {
    s.OutboxId = &v
    return s
}

func (req *AlibabaAlihealthSynergySyOutboxDetailRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.OutboxId != nil) {
        paramMap["outbox_id"] = *req.OutboxId
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyOutboxDetailRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}