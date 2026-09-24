package domain


type AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist struct {
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

    /*
        码信息     */
    CodeInfoList  *[]AlibabaAlihealthDrugYljgSearchbillDetailCodeInfoList `json:"code_info_list,omitempty" `

}

func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetExpiredDate(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ExpiredDate = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetProduceEntName(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetProdCode(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetProductCode(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProductCode = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetProduceDate(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetProductBatchNo(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProductBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetPhysicName(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetPreparationsUnit(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PreparationsUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetTempPkgSpec(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.TempPkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetMinPreparationsCount(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.MinPreparationsCount = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetMinPkgCount(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.MinPkgCount = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetPhysicTypeName(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PhysicTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetPhysicType(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetApproveNo(v string) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ApproveNo = &v
    return s
}
func (s *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist) SetCodeInfoList(v []AlibabaAlihealthDrugYljgSearchbillDetailCodeInfoList) *AlibabaAlihealthDrugYljgSearchbillDetailBillchkinoutdetaillistdtolist {
    s.CodeInfoList = &v
    return s
}
