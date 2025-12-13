package domain


type AlibabaAlihealthDrugtraceTopYljgUploadinoutbillOperationInfoMap struct {
    /*
        json,key是错误类型编码,value是具体得信息,详细参考文档     */
    CodeCheckErrorInfo  *string `json:"code_check_error_info,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillOperationInfoMap) SetCodeCheckErrorInfo(v string) *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillOperationInfoMap {
    s.CodeCheckErrorInfo = &v
    return s
}
