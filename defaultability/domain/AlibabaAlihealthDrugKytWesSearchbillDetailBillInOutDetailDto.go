package domain


type AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto struct {
    /*
        修改时间     */
    ModDate  *string `json:"mod_date,omitempty" `

    /*
        处理时间     */
    ProcessDate  *string `json:"process_date,omitempty" `

    /*
        单据日期     */
    BillTime  *string `json:"bill_time,omitempty" `

    /*
        收货企业id     */
    ToUserId  *string `json:"to_user_id,omitempty" `

    /*
        收货企业名称     */
    ToEntName  *string `json:"to_ent_name,omitempty" `

    /*
        发货企业id     */
    FromUserId  *string `json:"from_user_id,omitempty" `

    /*
        发货企业名称     */
    FromEntName  *string `json:"from_ent_name,omitempty" `

    /*
        单据类型名称     */
    BillTypeName  *string `json:"bill_type_name,omitempty" `

    /*
        单据类型     */
    BillType  *string `json:"bill_type,omitempty" `

    /*
        单据号码     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        单据详情     */
    BillChkInOutDetailListDTOList  *[]AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist `json:"bill_chk_in_out_detail_list_d_t_o_list,omitempty" `

    /*
        单据中的码     */
    Codes  *[]string `json:"codes,omitempty" `

    /*
        出库单id     */
    BillOutId  *string `json:"bill_out_id,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto) SetModDate(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto {
    s.ModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto) SetProcessDate(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto) SetBillTime(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto) SetToUserId(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto) SetToEntName(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto {
    s.ToEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto) SetFromUserId(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto) SetFromEntName(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto {
    s.FromEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto) SetBillTypeName(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto) SetBillType(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto) SetBillCode(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto) SetBillChkInOutDetailListDTOList(v []AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto {
    s.BillChkInOutDetailListDTOList = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto) SetCodes(v []string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto {
    s.Codes = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto) SetBillOutId(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillInOutDetailDto {
    s.BillOutId = &v
    return s
}
