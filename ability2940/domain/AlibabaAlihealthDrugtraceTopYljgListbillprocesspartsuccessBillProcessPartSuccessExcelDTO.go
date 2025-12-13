package domain


type AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessBillProcessPartSuccessExcelDTO struct {
    /*
        最后一次重新处理时间     */
    LastProcessDateDesc  *string `json:"last_process_date_desc,omitempty" `

    /*
        追溯码     */
    Code  *string `json:"code,omitempty" `

    /*
        处理失败原因描述     */
    ProcessStatusReasonDesc  *string `json:"process_status_reason_desc,omitempty" `

    /*
        错误类型描述     */
    ErrorCodeDesc  *string `json:"error_code_desc,omitempty" `

    /*
        错误类型     */
    ErrorCode  *string `json:"error_code,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetLastProcessDateDesc(v string) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.LastProcessDateDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetCode(v string) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetProcessStatusReasonDesc(v string) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.ProcessStatusReasonDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetErrorCodeDesc(v string) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.ErrorCodeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetErrorCode(v string) *AlibabaAlihealthDrugtraceTopYljgListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.ErrorCode = &v
    return s
}
