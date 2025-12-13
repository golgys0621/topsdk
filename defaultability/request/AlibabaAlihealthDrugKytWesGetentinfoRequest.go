package request


type AlibabaAlihealthDrugKytWesGetentinfoRequest struct {
    /*
        企业名称     */
    EntName  *string `json:"ent_name" required:"true" `
}

func (s *AlibabaAlihealthDrugKytWesGetentinfoRequest) SetEntName(v string) *AlibabaAlihealthDrugKytWesGetentinfoRequest {
    s.EntName = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesGetentinfoRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.EntName != nil) {
        paramMap["ent_name"] = *req.EntName
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesGetentinfoRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}