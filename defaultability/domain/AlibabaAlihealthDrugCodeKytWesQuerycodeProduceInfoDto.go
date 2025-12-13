package domain


type AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto struct {
    /*
        生产日期     */
    ProduceDateStr  *string `json:"produce_date_str,omitempty" `

    /*
        最小包装数量     */
    PkgAmount  *string `json:"pkg_amount,omitempty" `

    /*
        有效期至     */
    ExpireDate  *string `json:"expire_date,omitempty" `

    /*
        批次号     */
    BatchNo  *string `json:"batch_no,omitempty" `

    /*
        生产日期     */
    OriginalProduceDate  *string `json:"original_produce_date,omitempty" `

    /*
        有效期至     */
    OriginalExpireDate  *string `json:"original_expire_date,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto) SetProduceDateStr(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto {
    s.ProduceDateStr = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto) SetPkgAmount(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto {
    s.PkgAmount = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto) SetExpireDate(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto {
    s.ExpireDate = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto) SetBatchNo(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto {
    s.BatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto) SetOriginalProduceDate(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto {
    s.OriginalProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto) SetOriginalExpireDate(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto {
    s.OriginalExpireDate = &v
    return s
}
