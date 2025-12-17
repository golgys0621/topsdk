package domain


type AlibabaAlihealthSynergySyContractSealPosition struct {
    /*
        盖章范围，全部盖章填all。1. 零散页码以英文逗号分隔，例如输入：1,3,11 2. 页码范围以英文短横线连接，例如输入：5-9 3. 上述两种方式可以组合使用, 例如输入：1,3,5-9,11     */
    PageRange  *string `json:"page_range,omitempty" `

    /*
        盖章位置     */
    X  *string `json:"x,omitempty" `

    /*
        盖章位置     */
    Y  *string `json:"y,omitempty" `

}

func (s *AlibabaAlihealthSynergySyContractSealPosition) SetPageRange(v string) *AlibabaAlihealthSynergySyContractSealPosition {
    s.PageRange = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractSealPosition) SetX(v string) *AlibabaAlihealthSynergySyContractSealPosition {
    s.X = &v
    return s
}
func (s *AlibabaAlihealthSynergySyContractSealPosition) SetY(v string) *AlibabaAlihealthSynergySyContractSealPosition {
    s.Y = &v
    return s
}
