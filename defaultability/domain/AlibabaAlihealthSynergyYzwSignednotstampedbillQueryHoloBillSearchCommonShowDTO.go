package domain


type AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO struct {
    /*
        单据类型     */
    BillType  *int64 `json:"bill_type,omitempty" `

    /*
        单据id     */
    BillId  *string `json:"bill_id,omitempty" `

    /*
        单据明细ID     */
    BillDetailId  *string `json:"bill_detail_id,omitempty" `

    /*
        单据时间     */
    BillTime  *string `json:"bill_time,omitempty" `

    /*
        单据编码     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        批号     */
    ProduceBatchNo  *string `json:"produce_batch_no,omitempty" `

    /*
        药品id     */
    DrugId  *string `json:"drug_id,omitempty" `

    /*
        签章URL     */
    SealedReportUrl  *string `json:"sealed_report_url,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO) SetBillType(v int64) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO) SetBillId(v string) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO {
    s.BillId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO) SetBillDetailId(v string) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO {
    s.BillDetailId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO) SetBillTime(v string) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO) SetBillCode(v string) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO) SetProduceBatchNo(v string) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO) SetDrugId(v string) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO {
    s.DrugId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO) SetSealedReportUrl(v string) *AlibabaAlihealthSynergyYzwSignednotstampedbillQueryHoloBillSearchCommonShowDTO {
    s.SealedReportUrl = &v
    return s
}
