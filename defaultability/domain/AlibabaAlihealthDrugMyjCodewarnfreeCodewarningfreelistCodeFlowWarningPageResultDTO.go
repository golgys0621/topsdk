package domain


type AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO struct {
    /*
        追溯码     */
    Code  *string `json:"code,omitempty" `

    /*
        任务批次编码     */
    TaskId  *int64 `json:"task_id,omitempty" `

    /*
        生成时间     */
    GmtCreate  *string `json:"gmt_create,omitempty" `

    /*
        查询结果状态（处理中，未见异常，异常）     */
    ResultStatusName  *string `json:"result_status_name,omitempty" `

    /*
        码级别     */
    CodeLevel  *int64 `json:"code_level,omitempty" `

    /*
        所含码数量     */
    AllCodeCount  *int64 `json:"all_code_count,omitempty" `

    /*
        药品信息     */
    DrugInfo  *string `json:"drug_info,omitempty" `

    /*
        产品批号     */
    ProduceBatchNo  *string `json:"produce_batch_no,omitempty" `

    /*
        生产日期     */
    ProduceDate  *string `json:"produce_date,omitempty" `

    /*
        有效期至     */
    ValidEndDate  *string `json:"valid_end_date,omitempty" `

    /*
        生产企业名称     */
    ProduceEntName  *string `json:"produce_ent_name,omitempty" `

    /*
        上市许可证持有人     */
    MahRefEntName  *string `json:"mah_ref_ent_name,omitempty" `

    /*
        创建人     */
    OperIcCode  *string `json:"oper_ic_code,omitempty" `

    /*
        任务生成方式（11: 扫码生成单据时生成 12: 上游出库单确认入库时生成 21:查询出入库单时生成 22: 查询上游出库单生成 23: 查询委托方上游出库单生成 3: TOP接口生成 4: 重新查询）     */
    TaskTypeName  *string `json:"task_type_name,omitempty" `

    /*
        来源单据号     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        任务流水编码       */
    Id  *int64 `json:"id,omitempty" `

    /*
        top接口     */
    TaskType  *int64 `json:"task_type,omitempty" `

    /*
        包装规格     */
    PkgSpec  *string `json:"pkg_spec,omitempty" `

    /*
        查询结果状态编码	0查询中 1未见异常 2异常     */
    ResultStatus  *string `json:"result_status,omitempty" `

}

func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetCode(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetTaskId(v int64) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.TaskId = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetGmtCreate(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.GmtCreate = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetResultStatusName(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.ResultStatusName = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetCodeLevel(v int64) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetAllCodeCount(v int64) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.AllCodeCount = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetDrugInfo(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.DrugInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetProduceBatchNo(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetProduceDate(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetValidEndDate(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.ValidEndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetProduceEntName(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetMahRefEntName(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.MahRefEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetOperIcCode(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.OperIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetTaskTypeName(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.TaskTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetBillCode(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetId(v int64) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetTaskType(v int64) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.TaskType = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetPkgSpec(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) SetResultStatus(v string) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO {
    s.ResultStatus = &v
    return s
}
