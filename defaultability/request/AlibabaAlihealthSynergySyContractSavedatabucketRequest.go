package request


type AlibabaAlihealthSynergySyContractSavedatabucketRequest struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        文件字节数组转base64:  1）读取文件转成byte[], 然后拆分字节数组分批上传。类似分页，每一页字节数组大小建议设置2000000(即2M),之后转成base64,调用此接口上传。  2）循环N次后，把接口返回的结果使用逗号拼接隔开，传给其他接口作为文件URL。      */
    Base64Data  *string `json:"base64_data" required:"true" `
    /*
        文件名，支持pdf , png ,jpg 等     */
    FileName  *string `json:"file_name" required:"true" `
}

func (s *AlibabaAlihealthSynergySyContractSavedatabucketRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyContractSavedatabucketRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractSavedatabucketRequest) SetBase64Data(v string) *AlibabaAlihealthSynergySyContractSavedatabucketRequest {
    s.Base64Data = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractSavedatabucketRequest) SetFileName(v string) *AlibabaAlihealthSynergySyContractSavedatabucketRequest {
    s.FileName = &v
    return s
}

func (req *AlibabaAlihealthSynergySyContractSavedatabucketRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.Base64Data != nil) {
        paramMap["base64_data"] = *req.Base64Data
    }
    if(req.FileName != nil) {
        paramMap["file_name"] = *req.FileName
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyContractSavedatabucketRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}