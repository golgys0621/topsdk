package request


type AlibabaAlihealthSynergySyInboxDetailRequest struct {
    /*
        本企业refEntId     */
    RefEntId  *string `json:"ref_ent_id,omitempty" required:"false" `
    /*
        收件箱id     */
    InboxId  *int64 `json:"inbox_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergySyInboxDetailRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyInboxDetailRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailRequest) SetInboxId(v int64) *AlibabaAlihealthSynergySyInboxDetailRequest {
    s.InboxId = &v
    return s
}

func (req *AlibabaAlihealthSynergySyInboxDetailRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.InboxId != nil) {
        paramMap["inbox_id"] = *req.InboxId
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyInboxDetailRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}