package domain


type AlibabaAlihealthDrugCodeKytWesCheckcoderelationallWesCodeRelationDTOV2 struct {
    /*
        存在上下级关系时返回下级码     */
    ResultList  *[]AlibabaAlihealthDrugCodeKytWesCheckcoderelationallWesCodeRelationDTO `json:"result_list,omitempty" `

    /*
        错误信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallWesCodeRelationDTOV2) SetResultList(v []AlibabaAlihealthDrugCodeKytWesCheckcoderelationallWesCodeRelationDTO) *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallWesCodeRelationDTOV2 {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallWesCodeRelationDTOV2) SetMsgInfo(v string) *AlibabaAlihealthDrugCodeKytWesCheckcoderelationallWesCodeRelationDTOV2 {
    s.MsgInfo = &v
    return s
}
