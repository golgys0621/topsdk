package domain


type TaobaoTmcQueueGetTmcQueueInfo struct {
    /*
        当前队列当天读取量     */
    GetTotal  *int64 `json:"get_total,omitempty" `

    /*
        当前队列当天写入量     */
    PutToal  *int64 `json:"put_toal,omitempty" `

    /*
        TMC组名     */
    Name  *string `json:"name,omitempty" `

    /*
        消息队列Broker名称     */
    BrokerName  *string `json:"broker_name,omitempty" `

}

func (s *TaobaoTmcQueueGetTmcQueueInfo) SetGetTotal(v int64) *TaobaoTmcQueueGetTmcQueueInfo {
    s.GetTotal = &v
    return s
}
func (s *TaobaoTmcQueueGetTmcQueueInfo) SetPutToal(v int64) *TaobaoTmcQueueGetTmcQueueInfo {
    s.PutToal = &v
    return s
}
func (s *TaobaoTmcQueueGetTmcQueueInfo) SetName(v string) *TaobaoTmcQueueGetTmcQueueInfo {
    s.Name = &v
    return s
}
func (s *TaobaoTmcQueueGetTmcQueueInfo) SetBrokerName(v string) *TaobaoTmcQueueGetTmcQueueInfo {
    s.BrokerName = &v
    return s
}
