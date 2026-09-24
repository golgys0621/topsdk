package domain


type AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto struct {
    /*
        生产日期，格式化日期，到月的，默认到月底。     */
    ProduceDateStr  *string `json:"produce_date_str,omitempty" `

    /*
        最小包装数量     */
    PkgAmount  *string `json:"pkg_amount,omitempty" `

    /*
        有效期至，格式化日期，到月的，默认到月底。     */
    ExpireDate  *string `json:"expire_date,omitempty" `

    /*
        批次号     */
    BatchNo  *string `json:"batch_no,omitempty" `

    /*
        生产日期,生产企业原始上传日期，有可能只到月     */
    OriginalProduceDate  *string `json:"original_produce_date,omitempty" `

    /*
        有效期至，生产企业原始上传日期，有可能只到月     */
    OriginalExpireDate  *string `json:"original_expire_date,omitempty" `

    /*
        上市许可持有人企业名字     */
    MahName  *string `json:"mah_name,omitempty" `

    /*
        上市许可持有人企业refEntId     */
    MahRefEntId  *string `json:"mah_ref_ent_id,omitempty" `

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
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto) SetMahName(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto {
    s.MahName = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto) SetMahRefEntId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeProduceInfoDto {
    s.MahRefEntId = &v
    return s
}
