package domain


type AlibabaAlihealthDrugtraceTopLsydUploadinoutbillOperationInfoMap struct {
    /*
        json,key是错误类型编码,value是具体得信息,详细参考文档     */
    CodeCheckErrorInfo  *string `json:"code_check_error_info,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillOperationInfoMap) SetCodeCheckErrorInfo(v string) *AlibabaAlihealthDrugtraceTopLsydUploadinoutbillOperationInfoMap {
    s.CodeCheckErrorInfo = &v
    return s
}
