package request


type AlibabaAlihealthSynergySyEntpartnerResourceListRequest struct {
    /*
        企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        合作企业refEntId     */
    PartnerRefEntId  *string `json:"partner_ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthSynergySyEntpartnerResourceListRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyEntpartnerResourceListRequest) SetPartnerRefEntId(v string) *AlibabaAlihealthSynergySyEntpartnerResourceListRequest {
    s.PartnerRefEntId = &v
    return s
}

func (req *AlibabaAlihealthSynergySyEntpartnerResourceListRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.PartnerRefEntId != nil) {
        paramMap["partner_ref_ent_id"] = *req.PartnerRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyEntpartnerResourceListRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}