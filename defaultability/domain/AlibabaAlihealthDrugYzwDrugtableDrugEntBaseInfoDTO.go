package domain


type AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO struct {
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

func (s *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO) SetDrugId(v string) *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO {
    s.DrugId = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO) SetPhysicName(v string) *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO) SetApprovalLicenceNo(v string) *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO {
    s.ApprovalLicenceNo = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO) SetMahEntName(v string) *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO {
    s.MahEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO) SetProduceEntName(v string) *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO) SetPrepnSpec(v string) *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO) SetPkgSpec(v string) *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO) SetPrepnTypeDesc(v string) *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO {
    s.PrepnTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO) SetPackEntName(v string) *AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO {
    s.PackEntName = &v
    return s
}
