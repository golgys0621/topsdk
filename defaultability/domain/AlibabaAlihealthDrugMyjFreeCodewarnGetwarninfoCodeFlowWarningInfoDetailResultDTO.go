package domain


type AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO struct {
    /*
        异常的追溯码     */
    WarningCode  *string `json:"warning_code,omitempty" `

    /*
        码级别     */
    CodeLevel  *int64 `json:"code_level,omitempty" `

    /*
        异常单据上传时间     */
    WarningDateDesc  *string `json:"warning_date_desc,omitempty" `

    /*
        异常级别，1高风险，2低风险     */
    WarningLevel  *int64 `json:"warning_level,omitempty" `

    /*
        异常企业企业类型     */
    WarningEntOrgTypeDesc  *string `json:"warning_ent_org_type_desc,omitempty" `

    /*
        异常企业省份     */
    WarningProvDesc  *string `json:"warning_prov_desc,omitempty" `

    /*
        异常类型描述     */
    WarningInfoDesc  *string `json:"warning_info_desc,omitempty" `

    /*
        异常单据类型     */
    WarningBillTypeDesc  *string `json:"warning_bill_type_desc,omitempty" `

    /*
        异常级别     */
    WarningLevelDesc  *string `json:"warning_level_desc,omitempty" `

    /*
        异常类型（1,高风险：该码在其他终端企业有禁止再流通的记录，如销毁、抽检、报废出库等；2,高风险：该码在其他终端企业存在使用出库的记录；3,低风险：该码在其他终端企业存在流通记录，并未查询到退货单据；4,低风险：该码所属的生产企业尚未进行激活；5,高风险：该码校验不合格，非平台码段；6,高风险：该码已经注销）     */
    WarningInfoType  *int64 `json:"warning_info_type,omitempty" `

}

func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningCode(v string) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetCodeLevel(v int64) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningDateDesc(v string) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningDateDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningLevel(v int64) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningEntOrgTypeDesc(v string) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningEntOrgTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningProvDesc(v string) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningProvDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningInfoDesc(v string) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningInfoDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningBillTypeDesc(v string) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningBillTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningLevelDesc(v string) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningLevelDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningInfoType(v int64) *AlibabaAlihealthDrugMyjFreeCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningInfoType = &v
    return s
}
