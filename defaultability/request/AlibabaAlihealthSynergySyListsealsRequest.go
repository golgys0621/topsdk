package request


type AlibabaAlihealthSynergySyListsealsRequest struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        印章类型:ALL:全部,OFFICIAL:公章,CONTRACT:合同章,LEGALPERSON:法人章,CHARGEPERSON:负责人章,OUTBOUND:出库章,QUALITY:质检章.     */
    SealType  *string `json:"seal_type" required:"true" `
}

func (s *AlibabaAlihealthSynergySyListsealsRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyListsealsRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListsealsRequest) SetSealType(v string) *AlibabaAlihealthSynergySyListsealsRequest {
    s.SealType = &v
    return s
}

func (req *AlibabaAlihealthSynergySyListsealsRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.SealType != nil) {
        paramMap["seal_type"] = *req.SealType
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyListsealsRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}