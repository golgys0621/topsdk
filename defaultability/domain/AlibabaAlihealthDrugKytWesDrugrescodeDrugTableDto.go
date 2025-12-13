package domain


type AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto struct {
    /*
        制剂类型描述     */
    PrepnTypeDesc  *string `json:"prepn_type_desc,omitempty" `

    /*
        药品类型描述     */
    PhysicTypeDesc  *string `json:"physic_type_desc,omitempty" `

    /*
        药品类型(详见码表) 1：特殊药品原料药，2：特殊药品制剂，3：普通药品，9：未分类     */
    PhysicType  *int64 `json:"physic_type,omitempty" `

    /*
        药品名称     */
    PhysicName  *string `json:"physic_name,omitempty" `

    /*
        药品自类编码     */
    ProdCode  *string `json:"prod_code,omitempty" `

    /*
        药品详细类型     */
    PhysicDetailType  *int64 `json:"physic_detail_type,omitempty" `

    /*
        企业主键     */
    RefEntId  *string `json:"ref_ent_id,omitempty" `

    /*
        商品名称     */
    ProdName  *string `json:"prod_name,omitempty" `

    /*
        修改日期     */
    ModDate  *string `json:"mod_date,omitempty" `

    /*
        企业名称（添加药品的企业）     */
    EntName  *string `json:"ent_name,omitempty" `

    /*
        包装单位描述     */
    PkgUnitDesc  *string `json:"pkg_unit_desc,omitempty" `

    /*
        药品类型详情描述     */
    PhysicDetailTypeDesc  *string `json:"physic_detail_type_desc,omitempty" `

    /*
        制剂单位描述     */
    PrepnUnitDesc  *string `json:"prepn_unit_desc,omitempty" `

    /*
        子列表     */
    SubTypeList  *[]AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList `json:"sub_type_list,omitempty" `

    /*
        生产企业名称     */
    ProduceRefName  *string `json:"produce_ref_name,omitempty" `

    /*
        mah企业名称     */
    AuthorizedRefName  *string `json:"authorized_ref_name,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetPrepnTypeDesc(v string) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.PrepnTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetPhysicTypeDesc(v string) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.PhysicTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetPhysicType(v int64) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetPhysicName(v string) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetProdCode(v string) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetPhysicDetailType(v int64) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.PhysicDetailType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetProdName(v string) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.ProdName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetModDate(v string) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.ModDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetEntName(v string) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.EntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetPkgUnitDesc(v string) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.PkgUnitDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetPhysicDetailTypeDesc(v string) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.PhysicDetailTypeDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetPrepnUnitDesc(v string) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.PrepnUnitDesc = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetSubTypeList(v []AlibabaAlihealthDrugKytWesDrugrescodeSubTypeList) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.SubTypeList = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetProduceRefName(v string) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.ProduceRefName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto) SetAuthorizedRefName(v string) *AlibabaAlihealthDrugKytWesDrugrescodeDrugTableDto {
    s.AuthorizedRefName = &v
    return s
}
