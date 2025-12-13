package domain


type AlibabaAlihealthSynergyYzwSavedrugreportwithocrTopFileDTO struct {
    /*
        文件名称     */
    Title  *string `json:"title,omitempty" `

    /*
         文件链接，链接格式支持：pdf, jpeg,jpg, png,  图片：最大边长小于5000像素,  注意：url不能含有空格     */
    Url  *string `json:"url,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwSavedrugreportwithocrTopFileDTO) SetTitle(v string) *AlibabaAlihealthSynergyYzwSavedrugreportwithocrTopFileDTO {
    s.Title = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSavedrugreportwithocrTopFileDTO) SetUrl(v string) *AlibabaAlihealthSynergyYzwSavedrugreportwithocrTopFileDTO {
    s.Url = &v
    return s
}
