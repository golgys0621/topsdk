package domain


type AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog struct {
    /*
        报告所属企业id     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        报告所属企业名称     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        批次号     */
    BatchNo  *string `json:"batch_no,omitempty" `

    /*
        药品id     */
    DrugId  *string `json:"drug_id,omitempty" `

    /*
        报告操作记录唯一值     */
    Uuid  *string `json:"uuid,omitempty" `

    /*
        操作类型insertReport：新增报告，deleteReport：删除报告，sealReport：报告签章，updateReport：报告更新     */
    OptType  *string `json:"opt_type,omitempty" `

    /*
        操作来源web：客户端，top：top平台，legao：小二操作     */
    OptFrom  *string `json:"opt_from,omitempty" `

    /*
        操作时间:yyyy-MM-dd HH:mm:ss     */
    OptTime  *string `json:"opt_time,omitempty" `

    /*
        操作企业id     */
    OptRefEntId  *string `json:"opt_ref_ent_id,omitempty" `

    /*
        操作企业名称     */
    OptEntName  *string `json:"opt_ent_name,omitempty" `

    /*
        药检报告id     */
    DrugReportV2Id  *string `json:"drug_report_v2_id,omitempty" `

    /*
        操作人     */
    OptUser  *string `json:"opt_user,omitempty" `

    /*
        报告日期yyyy-MM-dd     */
    ReportDate  *string `json:"report_date,omitempty" `

    /*
        ocr上传报告名字     */
    ReportName  *string `json:"report_name,omitempty" `

    /*
        报告编号     */
    ReportNo  *string `json:"report_no,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) SetEntName(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog {
    s.BatchNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) SetDrugId(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog {
    s.DrugId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) SetUuid(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog {
    s.Uuid = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) SetOptType(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog {
    s.OptType = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) SetOptFrom(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog {
    s.OptFrom = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) SetOptTime(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog {
    s.OptTime = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) SetOptRefEntId(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog {
    s.OptRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) SetOptEntName(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog {
    s.OptEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) SetDrugReportV2Id(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog {
    s.DrugReportV2Id = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) SetOptUser(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog {
    s.OptUser = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) SetReportDate(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog {
    s.ReportDate = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) SetReportName(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog {
    s.ReportName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog) SetReportNo(v string) *AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllTopDrugReportOptLog {
    s.ReportNo = &v
    return s
}
