package request


type AlibabaAlihealthSynergyYzwDruginfoQueryRequest struct {
    /*
        mah企业refEntId     */
    MahRefEntId  *string `json:"mah_ref_ent_id" required:"true" `
    /*
        批次号     */
    BatchNo  *string `json:"batch_no" required:"true" `
    /*
        生产企业refEntId     */
    ProduceRefEntId  *string `json:"produce_ref_ent_id" required:"true" `
    /*
        调用者企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthSynergyYzwDruginfoQueryRequest) SetMahRefEntId(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryRequest {
    s.MahRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryRequest) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryRequest {
    s.BatchNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryRequest) SetProduceRefEntId(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryRequest {
    s.ProduceRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryRequest {
    s.RefEntId = &v
    return s
}

func (req *AlibabaAlihealthSynergyYzwDruginfoQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.MahRefEntId != nil) {
        paramMap["mah_ref_ent_id"] = *req.MahRefEntId
    }
    if(req.BatchNo != nil) {
        paramMap["batch_no"] = *req.BatchNo
    }
    if(req.ProduceRefEntId != nil) {
        paramMap["produce_ref_ent_id"] = *req.ProduceRefEntId
    }
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergyYzwDruginfoQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}