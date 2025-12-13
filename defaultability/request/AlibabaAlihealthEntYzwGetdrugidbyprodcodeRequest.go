package request


type AlibabaAlihealthEntYzwGetdrugidbyprodcodeRequest struct {
    /*
        企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        药品子类编码（prodCode）     */
    ProdCode  *string `json:"prod_code" required:"true" `
    /*
        包装规格     */
    PkgSpec  *string `json:"pkg_spec" required:"true" `
}

func (s *AlibabaAlihealthEntYzwGetdrugidbyprodcodeRequest) SetRefEntId(v string) *AlibabaAlihealthEntYzwGetdrugidbyprodcodeRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthEntYzwGetdrugidbyprodcodeRequest) SetProdCode(v string) *AlibabaAlihealthEntYzwGetdrugidbyprodcodeRequest {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthEntYzwGetdrugidbyprodcodeRequest) SetPkgSpec(v string) *AlibabaAlihealthEntYzwGetdrugidbyprodcodeRequest {
    s.PkgSpec = &v
    return s
}

func (req *AlibabaAlihealthEntYzwGetdrugidbyprodcodeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.ProdCode != nil) {
        paramMap["prod_code"] = *req.ProdCode
    }
    if(req.PkgSpec != nil) {
        paramMap["pkg_spec"] = *req.PkgSpec
    }
    return paramMap
}

func (req *AlibabaAlihealthEntYzwGetdrugidbyprodcodeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}