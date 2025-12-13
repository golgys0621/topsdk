package request


type AlibabaAlihealthEntYzwHasuploaddrugreportRequest struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        药品ID     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id" required:"true" `
    /*
        批次号     */
    BatchNo  *string `json:"batch_no" required:"true" `
}

func (s *AlibabaAlihealthEntYzwHasuploaddrugreportRequest) SetRefEntId(v string) *AlibabaAlihealthEntYzwHasuploaddrugreportRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthEntYzwHasuploaddrugreportRequest) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthEntYzwHasuploaddrugreportRequest {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthEntYzwHasuploaddrugreportRequest) SetBatchNo(v string) *AlibabaAlihealthEntYzwHasuploaddrugreportRequest {
    s.BatchNo = &v
    return s
}

func (req *AlibabaAlihealthEntYzwHasuploaddrugreportRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.DrugEntBaseInfoId != nil) {
        paramMap["drug_ent_base_info_id"] = *req.DrugEntBaseInfoId
    }
    if(req.BatchNo != nil) {
        paramMap["batch_no"] = *req.BatchNo
    }
    return paramMap
}

func (req *AlibabaAlihealthEntYzwHasuploaddrugreportRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}