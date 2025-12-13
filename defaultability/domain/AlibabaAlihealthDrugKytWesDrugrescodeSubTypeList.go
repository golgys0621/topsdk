package domain


type AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList struct {
    /*
        制剂单位     */
    PrepnUnit  *string `json:"prepn_unit,omitempty" `

    /*
        包装规格     */
    PackageSpec  *string `json:"package_spec,omitempty" `

    /*
        码列表     */
    CodeResList  *[]AlibabaAlihealthDrugKytWesDrugrescodeCodeResList `json:"code_res_list,omitempty" `

    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        企业药品ID     */
    ProdSeqNo  *string `json:"prod_seq_no,omitempty" `

    /*
        批准文号     */
    ApproveNo  *string `json:"approve_no,omitempty" `

    /*
        药品详情类型     */
    PhysicDetailType  *string `json:"physic_detail_type,omitempty" `

    /*
        包装单位     */
    PackUnit  *string `json:"pack_unit,omitempty" `

    /*
        药品ID     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" `

    /*
        包装单位     */
    PackUnitName  *string `json:"pack_unit_name,omitempty" `

    /*
        制剂描述     */
    PrepnDesc  *string `json:"prepn_desc,omitempty" `

    /*
        制剂单位描述     */
    PrepnUnitName  *string `json:"prepn_unit_name,omitempty" `

    /*
        子类型     */
    SubTypeNo  *string `json:"sub_type_no,omitempty" `

    /*
        旧批准文号     */
    ApproveNoOld  *string `json:"approve_no_old,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList) SetPrepnUnit(v string) *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList {
    s.PrepnUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList) SetPackageSpec(v string) *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList {
    s.PackageSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList) SetCodeResList(v []AlibabaAlihealthDrugKytWesDrugrescodeCodeResList) *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList {
    s.CodeResList = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList) SetPrepnSpec(v string) *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList) SetProdSeqNo(v string) *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList {
    s.ProdSeqNo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList) SetApproveNo(v string) *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList {
    s.ApproveNo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList) SetPhysicDetailType(v string) *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList {
    s.PhysicDetailType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList) SetPackUnit(v string) *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList {
    s.PackUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList) SetPackUnitName(v string) *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList {
    s.PackUnitName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList) SetPrepnDesc(v string) *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList {
    s.PrepnDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList) SetPrepnUnitName(v string) *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList {
    s.PrepnUnitName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList) SetSubTypeNo(v string) *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList {
    s.SubTypeNo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList) SetApproveNoOld(v string) *AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList {
    s.ApproveNoOld = &v
    return s
}
