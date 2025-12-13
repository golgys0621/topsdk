package domain


type AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO struct {
    /*
        药品id     */
    DrugId  *string `json:"drug_id,omitempty" `

    /*
        药品通用名     */
    PhysicName  *string `json:"physic_name,omitempty" `

    /*
        批准文号     */
    ApprovalLicenceNo  *string `json:"approval_licence_no,omitempty" `

    /*
        mah名称     */
    MahEntName  *string `json:"mah_ent_name,omitempty" `

    /*
        生产企业名称     */
    ProduceEntName  *string `json:"produce_ent_name,omitempty" `

    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        包装规格     */
    PkgSpec  *string `json:"pkg_spec,omitempty" `

    /*
        剂型     */
    PrepnTypeDesc  *string `json:"prepn_type_desc,omitempty" `

    /*
        包装厂名称     */
    PackEntName  *string `json:"pack_ent_name,omitempty" `

}

func (s *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO) SetDrugId(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO {
    s.DrugId = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO) SetPhysicName(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO) SetApprovalLicenceNo(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO {
    s.ApprovalLicenceNo = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO) SetMahEntName(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO {
    s.MahEntName = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO) SetProduceEntName(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO) SetPrepnSpec(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO) SetPkgSpec(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO) SetPrepnTypeDesc(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO {
    s.PrepnTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO) SetPackEntName(v string) *AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO {
    s.PackEntName = &v
    return s
}
