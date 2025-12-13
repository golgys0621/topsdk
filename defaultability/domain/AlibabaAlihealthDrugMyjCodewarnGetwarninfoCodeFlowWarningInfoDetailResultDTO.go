package domain


type AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO struct {
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

func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningCode(v string) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningCode = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetCodeLevel(v int64) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningDateDesc(v string) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningDateDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningLevel(v int64) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningEntOrgTypeDesc(v string) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningEntOrgTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningProvDesc(v string) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningProvDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningInfoDesc(v string) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningInfoDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningBillTypeDesc(v string) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningBillTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningLevelDesc(v string) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningLevelDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO) SetWarningInfoType(v int64) *AlibabaAlihealthDrugMyjCodewarnGetwarninfoCodeFlowWarningInfoDetailResultDTO {
    s.WarningInfoType = &v
    return s
}
