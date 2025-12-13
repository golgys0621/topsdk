package domain


type AlibabaAlihealthDrugCodeKytWesQuerycodeCodeProduceInfoDto struct {
    /*
        生产信息集合     */
    ProduceInfoList  *[]AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto `json:"produce_info_list,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeProduceInfoDto) SetProduceInfoList(v []AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto) *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeProduceInfoDto {
    s.ProduceInfoList = &v
    return s
}
