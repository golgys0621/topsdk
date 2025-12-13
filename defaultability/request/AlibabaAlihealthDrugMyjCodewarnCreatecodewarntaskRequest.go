package request


type AlibabaAlihealthDrugMyjCodewarnCreatecodewarntaskRequest struct {
    /*
        企业refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        追溯码列表（多个追溯码，以逗号拼接）     */
    Codes  *string `json:"codes" required:"true" `
    /*
        来源单据号     */
    BillCode  *string `json:"bill_code,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugMyjCodewarnCreatecodewarntaskRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMyjCodewarnCreatecodewarntaskRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCreatecodewarntaskRequest) SetCodes(v string) *AlibabaAlihealthDrugMyjCodewarnCreatecodewarntaskRequest {
    s.Codes = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCreatecodewarntaskRequest) SetBillCode(v string) *AlibabaAlihealthDrugMyjCodewarnCreatecodewarntaskRequest {
    s.BillCode = &v
    return s
}

func (req *AlibabaAlihealthDrugMyjCodewarnCreatecodewarntaskRequest) ToMap() map[string]interface{} {
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
    return paramMap
}

func (req *AlibabaAlihealthDrugMyjCodewarnCreatecodewarntaskRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}