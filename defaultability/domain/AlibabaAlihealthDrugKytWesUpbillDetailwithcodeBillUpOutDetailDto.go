package domain


type AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto struct {
    /*
        单据编码     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        单据类型描述     */
    BillTypeName  *string `json:"bill_type_name,omitempty" `

    /*
        单据类型     */
    BillType  *string `json:"bill_type,omitempty" `

    /*
        发货企业名称     */
    EntSendName  *string `json:"ent_send_name,omitempty" `

    /*
        发货企业的ref_ent_id     */
    EntSendId  *string `json:"ent_send_id,omitempty" `

    /*
        收货企业名称     */
    EntRecvName  *string `json:"ent_recv_name,omitempty" `

    /*
        收货企业ref_ent_id     */
    EntRecvId  *string `json:"ent_recv_id,omitempty" `

    /*
        单据日期     */
    StoreOutDate  *string `json:"store_out_date,omitempty" `

    /*
        最后更新时间     */
    UpdateDate  *string `json:"update_date,omitempty" `

    /*
        药品信息数据     */
    DrugInfosDtoList  *[]AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto `json:"drug_infos_dto_list,omitempty" `

    /*
        单据上传时间     */
    CrtDate  *string `json:"crt_date,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto) SetBillCode(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto) SetBillTypeName(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto) SetBillType(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto) SetEntSendName(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto {
    s.EntSendName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto) SetEntSendId(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto {
    s.EntSendId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto) SetEntRecvName(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto {
    s.EntRecvName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto) SetEntRecvId(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto {
    s.EntRecvId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto) SetStoreOutDate(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto {
    s.StoreOutDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto) SetUpdateDate(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto {
    s.UpdateDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto) SetDrugInfosDtoList(v []AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto {
    s.DrugInfosDtoList = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto) SetCrtDate(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeBillUpOutDetailDto {
    s.CrtDate = &v
    return s
}
