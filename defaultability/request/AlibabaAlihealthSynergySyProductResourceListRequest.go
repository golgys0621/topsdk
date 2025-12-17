package request


type AlibabaAlihealthSynergySyProductResourceListRequest struct {
    /*
        企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        产品id     */
    SyProductId  *int64 `json:"sy_product_id" required:"true" `
    /*
        使用状态，0已过期1使用中     */
    UseStatus  *int64 `json:"use_status,omitempty" required:"false" `
    /*
        归档状态，0未归档1已归档     */
    ArchiveStatus  *int64 `json:"archive_status,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergySyProductResourceListRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyProductResourceListRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListRequest) SetSyProductId(v int64) *AlibabaAlihealthSynergySyProductResourceListRequest {
    s.SyProductId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListRequest) SetUseStatus(v int64) *AlibabaAlihealthSynergySyProductResourceListRequest {
    s.UseStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyProductResourceListRequest) SetArchiveStatus(v int64) *AlibabaAlihealthSynergySyProductResourceListRequest {
    s.ArchiveStatus = &v
    return s
}

func (req *AlibabaAlihealthSynergySyProductResourceListRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.SyProductId != nil) {
        paramMap["sy_product_id"] = *req.SyProductId
    }
    if(req.UseStatus != nil) {
        paramMap["use_status"] = *req.UseStatus
    }
    if(req.ArchiveStatus != nil) {
        paramMap["archive_status"] = *req.ArchiveStatus
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyProductResourceListRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}