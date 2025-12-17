package request


type AlibabaAlihealthSynergySyQualificationListRequest struct {
    /*
        企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        资料类型，all:全部，ent:企业资料,product:药品产品资料,agentUser:委托人资料,productDrug、productEquip、productOther已废弃     */
    Type  *string `json:"type" required:"true" `
}

func (s *AlibabaAlihealthSynergySyQualificationListRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyQualificationListRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyQualificationListRequest) SetType(v string) *AlibabaAlihealthSynergySyQualificationListRequest {
    s.Type = &v
    return s
}

func (req *AlibabaAlihealthSynergySyQualificationListRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyQualificationListRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}