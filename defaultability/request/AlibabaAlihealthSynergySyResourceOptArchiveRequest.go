package request


type AlibabaAlihealthSynergySyResourceOptArchiveRequest struct {
    /*
        资料所属企业id     */
    RefEntId  *string `json:"ref_ent_id,omitempty" required:"false" `
    /*
        资料类型0:企业资料1:产品资料2:委托人     */
    Type  *string `json:"type,omitempty" required:"false" `
    /*
        资料id     */
    ResourceId  *int64 `json:"resource_id,omitempty" required:"false" `
    /*
        0 未归档，1归档     */
    ArchiveStatus  *int64 `json:"archive_status,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergySyResourceOptArchiveRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyResourceOptArchiveRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyResourceOptArchiveRequest) SetType(v string) *AlibabaAlihealthSynergySyResourceOptArchiveRequest {
    s.Type = &v
    return s
}
func (s *AlibabaAlihealthSynergySyResourceOptArchiveRequest) SetResourceId(v int64) *AlibabaAlihealthSynergySyResourceOptArchiveRequest {
    s.ResourceId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyResourceOptArchiveRequest) SetArchiveStatus(v int64) *AlibabaAlihealthSynergySyResourceOptArchiveRequest {
    s.ArchiveStatus = &v
    return s
}

func (req *AlibabaAlihealthSynergySyResourceOptArchiveRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.ResourceId != nil) {
        paramMap["resource_id"] = *req.ResourceId
    }
    if(req.ArchiveStatus != nil) {
        paramMap["archive_status"] = *req.ArchiveStatus
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyResourceOptArchiveRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}