package domain


type AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO struct {
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

func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetCode(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetTaskId(v int64) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.TaskId = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetGmtCreate(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.GmtCreate = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetResultStatusName(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.ResultStatusName = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetCodeLevel(v int64) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetAllCodeCount(v int64) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.AllCodeCount = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetDrugInfo(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.DrugInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetProduceBatchNo(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetProduceDate(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetValidEndDate(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.ValidEndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetProduceEntName(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetMahRefEntName(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.MahRefEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetOperIcCode(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.OperIcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetTaskTypeName(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.TaskTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetBillCode(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetId(v int64) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetTaskType(v int64) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.TaskType = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetPkgSpec(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) SetResultStatus(v string) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO {
    s.ResultStatus = &v
    return s
}
