package domain


type AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist struct {
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
    CodeInfoList  *[]AlibabaAlihealthDrugLsydSearchbillDetailCodeInfoList `json:"code_info_list,omitempty" `

}

func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetExpiredDate(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ExpiredDate = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetProduceEntName(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProduceEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetProdCode(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProdCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetProductCode(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProductCode = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetProduceDate(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetProductBatchNo(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ProductBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetPhysicName(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PhysicName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetPreparationsUnit(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PreparationsUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetTempPkgSpec(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.TempPkgSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetMinPreparationsCount(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.MinPreparationsCount = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetMinPkgCount(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.MinPkgCount = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetPhysicTypeName(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PhysicTypeName = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetPhysicType(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.PhysicType = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetApproveNo(v string) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.ApproveNo = &v
    return s
}
func (s *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist) SetCodeInfoList(v []AlibabaAlihealthDrugLsydSearchbillDetailCodeInfoList) *AlibabaAlihealthDrugLsydSearchbillDetailBillchkinoutdetaillistdtolist {
    s.CodeInfoList = &v
    return s
}
