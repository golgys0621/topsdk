package domain


type AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto struct {
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
    BillChkInOutDetailListDTOList  *[]AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist `json:"bill_chk_in_out_detail_list_d_t_o_list,omitempty" `

    /*
        单据中的码     */
    Codes  *[]string `json:"codes,omitempty" `

    /*
        出库单id     */
    BillOutId  *string `json:"bill_out_id,omitempty" `

    /*
        配送单位     */
    DisEntInfoList  *[]AlibabaAlihealthDrugLsydSearchbillDetailDisentinfolist `json:"dis_ent_info_list,omitempty" `

}

func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto) SetModDate(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto {
    s.ModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto) SetProcessDate(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto) SetBillTime(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto) SetToUserId(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto) SetToEntName(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto {
    s.ToEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto) SetFromUserId(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto) SetFromEntName(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto {
    s.FromEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto) SetBillTypeName(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto) SetBillType(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto) SetBillCode(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto) SetBillChkInOutDetailListDTOList(v []AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto {
    s.BillChkInOutDetailListDTOList = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto) SetCodes(v []string) *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto {
    s.Codes = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto) SetBillOutId(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto {
    s.BillOutId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto) SetDisEntInfoList(v []AlibabaAlihealthDrugLsydSearchbillDetailDisentinfolist) *AlibabaAlihealthDrugLsydSearchbillDetailBillInOutDetailDto {
    s.DisEntInfoList = &v
    return s
}
