package domain


type AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto struct {
    /*
        码对象     */
    CodeStatusTypeDTO  *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeStatusTypeDto `json:"code_status_type_d_t_o,omitempty" `

    /*
        追溯码     */
    Code  *string `json:"code,omitempty" `

    /*
        企业信息对象     */
    PUserEntDTO  *AlibabaAlihealthDrugCodeKytWesQuerycodePUserEntDto `json:"p_user_ent_d_t_o,omitempty" `

    /*
        码包装层级，1代表最小包装。如：申请的包装比例是1:5:10, 对应的包装等级就是 3、2、1，代表大包装、中包装、小包装     */
    PackageLevel  *string `json:"package_level,omitempty" `

    /*
        药品基本信息对象     */
    DrugEntBaseDTO  *AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto `json:"drug_ent_base_d_t_o,omitempty" `

    /*
        码生产信息对象     */
    CodeProduceInfoDTO  *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeProduceInfoDto `json:"code_produce_info_d_t_o,omitempty" `

    /*
        包装比例     */
    PkgRatio  *string `json:"pkg_ratio,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto) SetCodeStatusTypeDTO(v AlibabaAlihealthDrugCodeKytWesQuerycodeCodeStatusTypeDto) *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto {
    s.CodeStatusTypeDTO = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto) SetCode(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto {
    s.Code = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto) SetPUserEntDTO(v AlibabaAlihealthDrugCodeKytWesQuerycodePUserEntDto) *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto {
    s.PUserEntDTO = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto) SetPackageLevel(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto {
    s.PackageLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto) SetDrugEntBaseDTO(v AlibabaAlihealthDrugCodeKytWesQuerycodeDrugEntBaseDto) *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto {
    s.DrugEntBaseDTO = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto) SetCodeProduceInfoDTO(v AlibabaAlihealthDrugCodeKytWesQuerycodeCodeProduceInfoDto) *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto {
    s.CodeProduceInfoDTO = &v
    return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto) SetPkgRatio(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeFullInfoDto {
    s.PkgRatio = &v
    return s
}
