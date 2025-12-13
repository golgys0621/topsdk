package domain


type AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusOnenetTaskDTO struct {
    /*
        任务状态：0待处理,1处理中,2成功,3文件无法下载,4文件合并错误,5其他错误,-1系统错误     */
    TaskStatus  *int64 `json:"task_status,omitempty" `

    /*
        备注     */
    Remark  *string `json:"remark,omitempty" `

    /*
        药检报告URL     */
    ReportUrl  *string `json:"report_url,omitempty" `

    /*
        签章后的药检报告URL     */
    SealReportUrl  *string `json:"seal_report_url,omitempty" `

    /*
        报告ID     */
    Id  *int64 `json:"id,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusOnenetTaskDTO) SetTaskStatus(v int64) *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusOnenetTaskDTO {
    s.TaskStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusOnenetTaskDTO) SetRemark(v string) *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusOnenetTaskDTO {
    s.Remark = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusOnenetTaskDTO) SetReportUrl(v string) *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusOnenetTaskDTO {
    s.ReportUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusOnenetTaskDTO) SetSealReportUrl(v string) *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusOnenetTaskDTO {
    s.SealReportUrl = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusOnenetTaskDTO) SetId(v int64) *AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusOnenetTaskDTO {
    s.Id = &v
    return s
}
