package request


type AlibabaAlihealthDrugMyjCodewarnfreeCreatecodewarnfreetaskRequest struct {
    /*
        企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        追溯码列表（多个追溯码，以逗号拼接）     */
    Codes  *string `json:"codes" required:"true" `
    /*
        来源单据号     */
    BillCode  *string `json:"bill_code" required:"true" `
    /*
        服务商id     */
    IsvRefEntId  *string `json:"isv_ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugMyjCodewarnfreeCreatecodewarnfreetaskRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCreatecodewarnfreetaskRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCreatecodewarnfreetaskRequest) SetCodes(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCreatecodewarnfreetaskRequest {
    s.Codes = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCreatecodewarnfreetaskRequest) SetBillCode(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCreatecodewarnfreetaskRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCreatecodewarnfreetaskRequest) SetIsvRefEntId(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCreatecodewarnfreetaskRequest {
    s.IsvRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugMyjCodewarnfreeCreatecodewarnfreetaskRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.Codes != nil) {
        paramMap["codes"] = *req.Codes
    }
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.IsvRefEntId != nil) {
        paramMap["isv_ref_ent_id"] = *req.IsvRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMyjCodewarnfreeCreatecodewarnfreetaskRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}