package domain


type AlibabaAlihealthDrugKytWesUploadcircubillOperationInfoMap struct {
    /*
        json,key是错误类型编码,value是具体得信息,详细参考文档     */
    CodeCheckErrorInfo  *string `json:"code_check_error_info,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesUploadcircubillOperationInfoMap) SetCodeCheckErrorInfo(v string) *AlibabaAlihealthDrugKytWesUploadcircubillOperationInfoMap {
    s.CodeCheckErrorInfo = &v
    return s
}
