package request


type AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest struct {
    /*
        企业标识     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        生成时间--开始日期     */
    GmtCreateBegin  *string `json:"gmt_create_begin,omitempty" required:"false" `
    /*
        生成时间--结束日期     */
    GmtCreateEnd  *string `json:"gmt_create_end,omitempty" required:"false" `
    /*
        查询结果状态编码	0查询中 1未见异常 2异常     */
    ResultStatus  *string `json:"result_status,omitempty" required:"false" `
    /*
        追溯码     */
    Code  *string `json:"code,omitempty" required:"false" `
    /*
        药品ID     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" required:"false" `
    /*
        产品批号     */
    ProduceBatchNo  *string `json:"produce_batch_no,omitempty" required:"false" `
    /*
        生产企业ID     */
    ProduceRefEntId  *string `json:"produce_ref_ent_id,omitempty" required:"false" `
    /*
        上市许可持有人企业ID     */
    MahRefEntId  *string `json:"mah_ref_ent_id,omitempty" required:"false" `
    /*
         21:查询出入库单时生成  22: 查询上游出库单生成  23: 查询委托方上游出库单生成  3: TOP接口生成  4: 重新查询     */
    TaskType  *string `json:"task_type,omitempty" required:"false" `
    /*
        任务批次编码     */
    TaskId  *int64 `json:"task_id,omitempty" required:"false" `
    /*
        来源单据号     */
    BillCode  *string `json:"bill_code,omitempty" required:"false" `
    /*
        页数     */
    Page  *int64 `json:"page" required:"true" `
    /*
        每页条数     */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        服务商id     */
    IsvRefEntId  *string `json:"isv_ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) SetGmtCreateBegin(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest {
    s.GmtCreateBegin = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) SetGmtCreateEnd(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest {
    s.GmtCreateEnd = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) SetResultStatus(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest {
    s.ResultStatus = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) SetCode(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) SetProduceBatchNo(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) SetProduceRefEntId(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest {
    s.ProduceRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) SetMahRefEntId(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest {
    s.MahRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) SetTaskType(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest {
    s.TaskType = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) SetTaskId(v int64) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest {
    s.TaskId = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) SetBillCode(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) SetPage(v int64) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest {
    s.Page = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) SetPageSize(v int64) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) SetIsvRefEntId(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest {
    s.IsvRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.GmtCreateBegin != nil) {
        paramMap["gmt_create_begin"] = *req.GmtCreateBegin
    }
    if(req.GmtCreateEnd != nil) {
        paramMap["gmt_create_end"] = *req.GmtCreateEnd
    }
    if(req.ResultStatus != nil) {
        paramMap["result_status"] = *req.ResultStatus
    }
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
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
    if(req.MahRefEntId != nil) {
        paramMap["mah_ref_ent_id"] = *req.MahRefEntId
    }
    if(req.TaskType != nil) {
        paramMap["task_type"] = *req.TaskType
    }
    if(req.TaskId != nil) {
        paramMap["task_id"] = *req.TaskId
    }
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.IsvRefEntId != nil) {
        paramMap["isv_ref_ent_id"] = *req.IsvRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}