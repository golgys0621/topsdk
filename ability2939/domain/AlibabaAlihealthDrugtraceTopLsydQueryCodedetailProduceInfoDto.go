package domain


type AlibabaAlihealthDrugtraceTopLsydQueryCodedetailProduceInfoDto struct {
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
        有效期     */
    OriginalExpireDate  *string `json:"original_expire_date,omitempty" `

}

func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailProduceInfoDto) SetProduceDateStr(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailProduceInfoDto {
    s.ProduceDateStr = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailProduceInfoDto) SetPkgAmount(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailProduceInfoDto {
    s.PkgAmount = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailProduceInfoDto) SetExpireDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailProduceInfoDto {
    s.ExpireDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailProduceInfoDto) SetBatchNo(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailProduceInfoDto {
    s.BatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailProduceInfoDto) SetOriginalProduceDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailProduceInfoDto {
    s.OriginalProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailProduceInfoDto) SetOriginalExpireDate(v string) *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailProduceInfoDto {
    s.OriginalExpireDate = &v
    return s
}
