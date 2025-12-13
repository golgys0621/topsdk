package domain


type AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto struct {
    /*
        激活信息     */
    CodeActiveInfoDTO  *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto `json:"code_active_info_d_t_o,omitempty" `

    /*
        码关联关系     */
    CodeRelationList  *[]AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeInfo `json:"code_relation_list,omitempty" `

    /*
        是否是最小包装     */
    IsSmallest  *string `json:"is_smallest,omitempty" `

    /*
        药品包装信息     */
    PkgInfoDTO  *AlibabaAlihealthDrugCodeKytWesQuerycoderelationPkgInfoDto `json:"pkg_info_d_t_o,omitempty" `

    /*
        药品基础信息     */
    BaseInfosDTO  *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfosDto `json:"base_infos_d_t_o,omitempty" `

    /*
        生产信息     */
    ProduceInfoList  *[]AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto `json:"produce_info_list,omitempty" `

    /*
        80012800128001280012     */
    Code  *string `json:"code,omitempty" `

    /*
        异常错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        异常错误信息     */
    ErrorInfo  *string `json:"error_info,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto) SetCodeActiveInfoDTO(v AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto {
    s.CodeActiveInfoDTO = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto) SetCodeRelationList(v []AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeInfo) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto {
    s.CodeRelationList = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto) SetIsSmallest(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto {
    s.IsSmallest = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto) SetPkgInfoDTO(v AlibabaAlihealthDrugCodeKytWesQuerycoderelationPkgInfoDto) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto {
    s.PkgInfoDTO = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto) SetBaseInfosDTO(v AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfosDto) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto {
    s.BaseInfosDTO = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto) SetProduceInfoList(v []AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto {
    s.ProduceInfoList = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto) SetCode(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto) SetErrorCode(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto {
    s.ErrorCode = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto) SetErrorInfo(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeRelationDto {
    s.ErrorInfo = &v
    return s
}
