package domain


type AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfosDto struct {
    /*
        药品基础信息     */
    BaseInfoList  *[]AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto `json:"base_info_list,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfosDto) SetBaseInfoList(v []AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfoDto) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationBaseInfosDto {
    s.BaseInfoList = &v
    return s
}
