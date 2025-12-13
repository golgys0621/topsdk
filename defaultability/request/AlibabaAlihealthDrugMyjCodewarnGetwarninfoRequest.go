package request


type AlibabaAlihealthDrugMyjCodewarnGetwarninfoRequest struct {
    /*
        企业ref_ent_id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        单个追溯码     */
    Code  *string `json:"code" required:"true" `
}

func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoRequest) SetCode(v string) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoRequest {
    s.Code = &v
    return s
}

func (req *AlibabaAlihealthDrugMyjCodewarnGetwarninfoRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMyjCodewarnGetwarninfoRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}