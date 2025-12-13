package request


type AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest struct {
    /*
        调用企业的refEntId     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        授权企业refEntId，报告属于该企业     */
    AssRefEntId  *string `json:"ass_ref_ent_id" required:"true" `
    /*
        生产日期起始     */
    ProduceDateBegin  *string `json:"produce_date_begin,omitempty" required:"false" `
    /*
        生产日期终止     */
    ProduceDateEnd  *string `json:"produce_date_end,omitempty" required:"false" `
    /*
        药品id     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" required:"false" `
    /*
        批次号     */
    ProduceBatchNo  *string `json:"produce_batch_no,omitempty" required:"false" `
    /*
        药品生产企业refEntId     */
    ProduceRefEntId  *string `json:"produce_ref_ent_id,omitempty" required:"false" `
    /*
        药品委托企业refEntId     */
    AuthRefEntId  *string `json:"auth_ref_ent_id,omitempty" required:"false" `
    /*
        报告状态: 1 未上传 ; 2 已上传未盖章 ; 3 已盖章     */
    Type  *string `json:"type,omitempty" required:"false" `
    /*
        页码 defalutValue��1    */
    Page  *int64 `json:"page" required:"true" `
    /*
        页的大小 defalutValue��20    */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest) SetAssRefEntId(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest {
    s.AssRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest) SetProduceDateBegin(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest {
    s.ProduceDateBegin = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest) SetProduceDateEnd(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest {
    s.ProduceDateEnd = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest) SetProduceBatchNo(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest) SetProduceRefEntId(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest {
    s.ProduceRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest) SetAuthRefEntId(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest {
    s.AuthRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest) SetType(v string) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest {
    s.Type = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest) SetPage(v int64) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest) SetPageSize(v int64) *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest {
    s.PageSize = &v
    return s
}

func (req *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.AssRefEntId != nil) {
        paramMap["ass_ref_ent_id"] = *req.AssRefEntId
    }
    if(req.ProduceDateBegin != nil) {
        paramMap["produce_date_begin"] = *req.ProduceDateBegin
    }
    if(req.ProduceDateEnd != nil) {
        paramMap["produce_date_end"] = *req.ProduceDateEnd
    }
    if(req.DrugEntBaseInfoId != nil) {
        paramMap["drug_ent_base_info_id"] = *req.DrugEntBaseInfoId
    }
    if(req.ProduceBatchNo != nil) {
        paramMap["produce_batch_no"] = *req.ProduceBatchNo
    }
    if(req.ProduceRefEntId != nil) {
        paramMap["produce_ref_ent_id"] = *req.ProduceRefEntId
    }
    if(req.AuthRefEntId != nil) {
        paramMap["auth_ref_ent_id"] = *req.AuthRefEntId
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}