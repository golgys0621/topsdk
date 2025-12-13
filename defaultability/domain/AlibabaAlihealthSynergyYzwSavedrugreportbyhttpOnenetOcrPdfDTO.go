package domain


type AlibabaAlihealthSynergyYzwSavedrugreportbyhttpOnenetOcrPdfDTO struct {
    /*
        文件名称 后缀仅支持:pdf,jpg,png     */
    Title  *string `json:"title,omitempty" `

    /*
        文件URL 注意：url不能含有空格     */
    Url  *string `json:"url,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpOnenetOcrPdfDTO) SetTitle(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpOnenetOcrPdfDTO {
    s.Title = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpOnenetOcrPdfDTO) SetUrl(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpOnenetOcrPdfDTO {
    s.Url = &v
    return s
}
