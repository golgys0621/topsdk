package request


type AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoRequest struct {
    /*
        查询的码，必须有上下游企业单据或本企业单据     */
    Code  *string `json:"code" required:"true" `
    /*
        服务商refEntId     */
    IsvRefEntId  *string `json:"isv_ref_ent_id" required:"true" `
    /*
        本企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoRequest) SetCode(v string) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoRequest {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoRequest) SetIsvRefEntId(v string) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoRequest {
    s.IsvRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoRequest {
    s.RefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    if(req.IsvRefEntId != nil) {
        paramMap["isv_ref_ent_id"] = *req.IsvRefEntId
    }
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}