package domain


type AlibabaAlihealthSynergySyResourceResultOnenetTaskDTO struct {
    /*
        任务id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        任务状态：0待处理,1处理中,2成功,3文件无法下载,4文件合并错误,5其他错误,-1系统错误     */
    TaskStatus  *string `json:"task_status,omitempty" `

    /*
        备注     */
    Remark  *string `json:"remark,omitempty" `

    /*
        扩展参数，resourceId是资料id     */
    ParamJson  *string `json:"param_json,omitempty" `

}

func (s *AlibabaAlihealthSynergySyResourceResultOnenetTaskDTO) SetId(v int64) *AlibabaAlihealthSynergySyResourceResultOnenetTaskDTO {
    s.Id = &v
    return s
}
func (s *AlibabaAlihealthSynergySyResourceResultOnenetTaskDTO) SetTaskStatus(v string) *AlibabaAlihealthSynergySyResourceResultOnenetTaskDTO {
    s.TaskStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergySyResourceResultOnenetTaskDTO) SetRemark(v string) *AlibabaAlihealthSynergySyResourceResultOnenetTaskDTO {
    s.Remark = &v
    return s
}
func (s *AlibabaAlihealthSynergySyResourceResultOnenetTaskDTO) SetParamJson(v string) *AlibabaAlihealthSynergySyResourceResultOnenetTaskDTO {
    s.ParamJson = &v
    return s
}
