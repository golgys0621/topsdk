package request


type AlibabaAlihealthSynergySyEntResourceListRequest struct {
    /*
        归档状态，0未归档1已归档     */
    ArchiveStatus  *int64 `json:"archive_status,omitempty" required:"false" `
    /*
        企业refEntId     */
    RefEntId  *string `json:"ref_ent_id,omitempty" required:"false" `
    /*
        使用状态，0已过期1使用中     */
    UseStatus  *int64 `json:"use_status,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergySyEntResourceListRequest) SetArchiveStatus(v int64) *AlibabaAlihealthSynergySyEntResourceListRequest {
    s.ArchiveStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyEntResourceListRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntResourceListRequest) SetUseStatus(v int64) *AlibabaAlihealthSynergySyEntResourceListRequest {
    s.UseStatus = &v
    return s
}

func (req *AlibabaAlihealthSynergySyEntResourceListRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ArchiveStatus != nil) {
        paramMap["archive_status"] = *req.ArchiveStatus
    }
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.UseStatus != nil) {
        paramMap["use_status"] = *req.UseStatus
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyEntResourceListRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}