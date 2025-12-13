package domain


type AlibabaAlihealthDrugKytWesListbillprocesspartsuccessBillProcessPartSuccessExcelDTO struct {
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

func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetLastProcessDateDesc(v string) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.LastProcessDateDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetCode(v string) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetProcessStatusReasonDesc(v string) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.ProcessStatusReasonDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetErrorCodeDesc(v string) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.ErrorCodeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessBillProcessPartSuccessExcelDTO) SetErrorCode(v string) *AlibabaAlihealthDrugKytWesListbillprocesspartsuccessBillProcessPartSuccessExcelDTO {
    s.ErrorCode = &v
    return s
}
