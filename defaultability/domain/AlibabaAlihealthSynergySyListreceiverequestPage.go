package domain


type AlibabaAlihealthSynergySyListreceiverequestPage struct {
    /*
        总记录数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        结果     */
    Result  *[]AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO `json:"result,omitempty" `

}

func (s *AlibabaAlihealthSynergySyListreceiverequestPage) SetTotalNum(v int64) *AlibabaAlihealthSynergySyListreceiverequestPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthSynergySyListreceiverequestPage) SetResult(v []AlibabaAlihealthSynergySyListreceiverequestSyReceiveRequestDTO) *AlibabaAlihealthSynergySyListreceiverequestPage {
    s.Result = &v
    return s
}
