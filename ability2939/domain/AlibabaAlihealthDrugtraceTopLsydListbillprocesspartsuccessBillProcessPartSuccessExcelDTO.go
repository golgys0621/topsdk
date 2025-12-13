package domain


type AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessBillProcessPartSuccessExcelDTO struct {
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

func (s *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetLastProcessDateDesc(v string) *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.LastProcessDateDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetCode(v string) *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetProcessStatusReasonDesc(v string) *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.ProcessStatusReasonDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetErrorCodeDesc(v string) *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.ErrorCodeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetErrorCode(v string) *AlibabaAlihealthDrugtraceTopLsydListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.ErrorCode = &v
    return s
}
