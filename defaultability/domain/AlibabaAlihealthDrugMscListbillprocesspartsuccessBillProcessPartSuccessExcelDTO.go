package domain


type AlibabaAlihealthDrugMscListbillprocesspartsuccessBillProcessPartSuccessExcelDTO struct {
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

func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetLastProcessDateDesc(v string) *AlibabaAlihealthDrugMscListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.LastProcessDateDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetCode(v string) *AlibabaAlihealthDrugMscListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetProcessStatusReasonDesc(v string) *AlibabaAlihealthDrugMscListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.ProcessStatusReasonDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetErrorCodeDesc(v string) *AlibabaAlihealthDrugMscListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.ErrorCodeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetErrorCode(v string) *AlibabaAlihealthDrugMscListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.ErrorCode = &v
    return s
}
