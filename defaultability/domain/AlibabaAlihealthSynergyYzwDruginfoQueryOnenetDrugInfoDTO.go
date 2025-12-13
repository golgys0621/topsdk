package domain


type AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO struct {
    /*
        药品id     */
    DrugId  *string `json:"drug_id,omitempty" `

    /*
        批次     */
    BatchNo  *string `json:"batch_no,omitempty" `

    /*
        mah名字     */
    MahName  *string `json:"mah_name,omitempty" `

    /*
        生产企业名称     */
    ProduceEntName  *string `json:"produce_ent_name,omitempty" `

    /*
        mahRefEntId     */
    MahRefEntId  *string `json:"mah_ref_ent_id,omitempty" `

    /*
        生产企业refEntId     */
    ProduceRefEntId  *string `json:"produce_ref_ent_id,omitempty" `

    /*
        是否激活:0未激活，1激活     */
    ActiveStatus  *string `json:"active_status,omitempty" `

    /*
        药品名称     */
    PhysicName  *string `json:"physic_name,omitempty" `

    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        包装规格     */
    PkgSpec  *string `json:"pkg_spec,omitempty" `

    /*
        生产日期     */
    ProduceDate  *string `json:"produce_date,omitempty" `

    /*
        有效期至     */
    ExpireDate  *string `json:"expire_date,omitempty" `

}

func (s *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO) SetDrugId(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO {
    s.DrugId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO) SetBatchNo(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO {
    s.BatchNo = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO) SetMahName(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO {
    s.MahName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO) SetProduceEntName(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO) SetMahRefEntId(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO {
    s.MahRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO) SetProduceRefEntId(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO {
    s.ProduceRefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO) SetActiveStatus(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO {
    s.ActiveStatus = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO) SetPhysicName(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO) SetPrepnSpec(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO) SetPkgSpec(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO {
    s.PkgSpec = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO) SetProduceDate(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO) SetExpireDate(v string) *AlibabaAlihealthSynergyYzwDruginfoQueryOnenetDrugInfoDTO {
    s.ExpireDate = &v
    return s
}
