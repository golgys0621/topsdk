package domain


type AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusOnenetTaskDTO struct {
    /*
        0:解析中 1:待提交 2:部分提交 3:已提交 -1:系统错误 -2:文件无法下载 -3:文件无法匹配药品  -5:文件名重复 -6:文件格式转换异常  status in (2,3) 代表人工处理状态，说明文件已经审核完成，可以去查询报告信息。审核不通过的会在remark返回明细，例如该药品批次已有历史药检报告。     */
    TaskStatus  *int64 `json:"task_status,omitempty" `

    /*
        备注     */
    Remark  *string `json:"remark,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusOnenetTaskDTO) SetTaskStatus(v int64) *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusOnenetTaskDTO {
    s.TaskStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusOnenetTaskDTO) SetRemark(v string) *AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusOnenetTaskDTO {
    s.Remark = &v
    return s
}
