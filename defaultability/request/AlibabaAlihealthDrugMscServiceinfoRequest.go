package request


type AlibabaAlihealthDrugMscServiceinfoRequest struct {
    /*
        企业的refEntId     */
    RefEntId  *string `json:"ref_ent_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugMscServiceinfoRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscServiceinfoRequest {
    s.RefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugMscServiceinfoRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMscServiceinfoRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}