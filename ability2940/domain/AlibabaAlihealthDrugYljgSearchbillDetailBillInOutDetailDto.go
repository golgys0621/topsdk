package domain


type AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto struct {
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
    BillChkInOutDetailListDTOList  *[]AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist `json:"bill_chk_in_out_detail_list_d_t_o_list,omitempty" `

    /*
        单据中的码     */
    Codes  *[]string `json:"codes,omitempty" `

    /*
        出库单id     */
    BillOutId  *string `json:"bill_out_id,omitempty" `

    /*
        配送单位     */
    DisEntInfoList  *[]AlibabaAlihealthDrugYljgSearchbillDetailDisentinfolist `json:"dis_ent_info_list,omitempty" `

}

func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto) SetModDate(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto {
    s.ModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto) SetProcessDate(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto) SetBillTime(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto) SetToUserId(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto) SetToEntName(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto {
    s.ToEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto) SetFromUserId(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto) SetFromEntName(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto {
    s.FromEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto) SetBillTypeName(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto {
    s.BillTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto) SetBillType(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto) SetBillCode(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto) SetBillChkInOutDetailListDTOList(v []AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto {
    s.BillChkInOutDetailListDTOList = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto) SetCodes(v []string) *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto {
    s.Codes = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto) SetBillOutId(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto {
    s.BillOutId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto) SetDisEntInfoList(v []AlibabaAlihealthDrugYljgSearchbillDetailDisentinfolist) *AlibabaAlihealthDrugYljgSearchbillDetailBillInOutDetailDto {
    s.DisEntInfoList = &v
    return s
}
