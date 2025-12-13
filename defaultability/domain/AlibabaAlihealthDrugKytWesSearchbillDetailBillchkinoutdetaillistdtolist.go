package domain


type AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist struct {
    /*
        有效期至     */
    ExpiredDate  *string `json:"expired_date,omitempty" `

    /*
        生产企业名称     */
    ProduceEntName  *string `json:"produce_ent_name,omitempty" `

    /*
        子类编码     */
    ProdCode  *string `json:"prod_code,omitempty" `

    /*
        子类编码前7位     */
    ProductCode  *string `json:"product_code,omitempty" `

    /*
        生产日期     */
    ProduceDate  *string `json:"produce_date,omitempty" `

    /*
        批次号     */
    ProductBatchNo  *string `json:"product_batch_no,omitempty" `

    /*
        药品id     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" `

    /*
        药品名称     */
    PhysicName  *string `json:"physic_name,omitempty" `

    /*
        制剂单位     */
    PreparationsUnit  *string `json:"preparations_unit,omitempty" `

    /*
        包装规格     */
    TempPkgSpec  *string `json:"temp_pkg_spec,omitempty" `

    /*
        最小制剂数量     */
    MinPreparationsCount  *string `json:"min_preparations_count,omitempty" `

    /*
        最小包装数量     */
    MinPkgCount  *string `json:"min_pkg_count,omitempty" `

    /*
        药品类型名称     */
    PhysicTypeName  *string `json:"physic_type_name,omitempty" `

    /*
        药品类型编码     */
    PhysicType  *string `json:"physic_type,omitempty" `

    /*
        国药准字     */
    ApproveNo  *string `json:"approve_no,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) SetExpiredDate(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ExpiredDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) SetProduceEntName(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) SetProdCode(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) SetProductCode(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProductCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) SetProduceDate(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) SetProductBatchNo(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProductBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) SetPhysicName(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) SetPreparationsUnit(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PreparationsUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) SetTempPkgSpec(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist {
    s.TempPkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) SetMinPreparationsCount(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist {
    s.MinPreparationsCount = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) SetMinPkgCount(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist {
    s.MinPkgCount = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) SetPhysicTypeName(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PhysicTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) SetPhysicType(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist) SetApproveNo(v string) *AlibabaAlihealthDrugKytWesSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ApproveNo = &v
    return s
}
