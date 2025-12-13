package domain


type AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassOnenetOcrPdfDTO struct {
    /*
        文件名称 后缀仅支持:pdf,jpg,png     */
    Title  *string `json:"title,omitempty" `

    /*
        文件URL 注意：url不能含有空格     */
    Url  *string `json:"url,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassOnenetOcrPdfDTO) SetTitle(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassOnenetOcrPdfDTO {
    s.Title = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassOnenetOcrPdfDTO) SetUrl(v string) *AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassOnenetOcrPdfDTO {
    s.Url = &v
    return s
}
