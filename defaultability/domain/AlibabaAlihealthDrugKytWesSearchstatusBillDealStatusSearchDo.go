package domain


type AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo struct {
    /*
        单据号     */
    StoreInoutSeqNo  *string `json:"store_inout_seq_no,omitempty" `

    /*
        药品类型     */
    PhysicType  *string `json:"physic_type,omitempty" `

    /*
        上传文件名     */
    UploadFileName  *string `json:"upload_file_name,omitempty" `

    /*
        发货单位     */
    FromUserName  *string `json:"from_user_name,omitempty" `

    /*
        角色类型     */
    RoleType  *string `json:"role_type,omitempty" `

    /*
        上传日期     */
    CrtDate  *string `json:"crt_date,omitempty" `

    /*
        操作人标识     */
    IcCode  *string `json:"ic_code,omitempty" `

    /*
        文件名     */
    ShortFileName  *string `json:"short_file_name,omitempty" `

    /*
        企业名称     */
    RefUserName  *string `json:"ref_user_name,omitempty" `

    /*
        单据日期     */
    BillTime  *string `json:"bill_time,omitempty" `

    /*
        处理状态  0，处理中 1, 上传成功     3, 处理成功     4, 处理失败     */
    ResultType  *string `json:"result_type,omitempty" `

    /*
        上传标识     */
    UploadFlag  *string `json:"upload_flag,omitempty" `

    /*
        处理结果表状态(暂不用)     */
    ProcessFlag  *string `json:"process_flag,omitempty" `

    /*
        处理日期     */
    ProcessDate  *string `json:"process_date,omitempty" `

    /*
        单号号     */
    BillCode  *string `json:"bill_code,omitempty" `

    /*
        单据类型     */
    BillType  *string `json:"bill_type,omitempty" `

    /*
        收货单位     */
    ToUserName  *string `json:"to_user_name,omitempty" `

    /*
        发货单位主键     */
    FromUserId  *string `json:"from_user_id,omitempty" `

    /*
        发货单位唯一标识     */
    FromRefUserId  *string `json:"from_ref_user_id,omitempty" `

    /*
        收货单位主键     */
    ToUserId  *string `json:"to_user_id,omitempty" `

    /*
        用户唯一标识     */
    RefUserId  *string `json:"ref_user_id,omitempty" `

    /*
        收货单位唯一标识     */
    ToRefUserId  *string `json:"to_ref_user_id,omitempty" `

    /*
        用户主键     */
    UserId  *string `json:"user_id,omitempty" `

    /*
        处理信息     */
    ProcessInfo  *string `json:"process_info,omitempty" `

    /*
        51全部成功 52部分成功     */
    SubProcessFlag  *string `json:"sub_process_flag,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetStoreInoutSeqNo(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.StoreInoutSeqNo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetPhysicType(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetUploadFileName(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.UploadFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetFromUserName(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.FromUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetRoleType(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.RoleType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetCrtDate(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.CrtDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetIcCode(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.IcCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetShortFileName(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.ShortFileName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetRefUserName(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.RefUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetBillTime(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetResultType(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.ResultType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetUploadFlag(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.UploadFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetProcessFlag(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.ProcessFlag = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetProcessDate(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.ProcessDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetBillCode(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetBillType(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetToUserName(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.ToUserName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetFromUserId(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.FromUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetFromRefUserId(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetToUserId(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.ToUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetRefUserId(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.RefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetToRefUserId(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetUserId(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.UserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetProcessInfo(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.ProcessInfo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo) SetSubProcessFlag(v string) *AlibabaAlihealthDrugKytWesSearchstatusBillDealStatusSearchDo {
    s.SubProcessFlag = &v
    return s
}
