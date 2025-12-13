package domain


type AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto struct {
    /*
        生产日期     */
    ProduceDate  *string `json:"produce_date,omitempty" `

    /*
        生产企业名称     */
    ProductEntName  *string `json:"product_ent_name,omitempty" `

    /*
        产品包装规格     */
    PackageSpec  *string `json:"package_spec,omitempty" `

    /*
        药品商品名     */
    ProdName  *string `json:"prod_name,omitempty" `

    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        制剂单位编码     */
    PrepnUnit  *string `json:"prepn_unit,omitempty" `

    /*
        批次号     */
    ProduceBatchNo  *string `json:"produce_batch_no,omitempty" `

    /*
        药品标识     */
    ProdSeqNo  *string `json:"prod_seq_no,omitempty" `

    /*
        药品标识     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id,omitempty" `

    /*
        有效期至     */
    ValidEndDate  *string `json:"valid_end_date,omitempty" `

    /*
        按最小包装单位统计数量     */
    LeastPkgAmount  *string `json:"least_pkg_amount,omitempty" `

    /*
        按最小制剂单位统计数量     */
    LeastPrepnAmount  *string `json:"least_prepn_amount,omitempty" `

    /*
        码信息     */
    CodeInfoListDtoList  *[]AlibabaAlihealthDrugKytWesUpbillDetailwithcodeCodeInfoListDto `json:"code_info_list_dto_list,omitempty" `

    /*
        单据明细id     */
    DetailBillId  *int64 `json:"detail_bill_id,omitempty" `

    /*
        药品通用名     */
    PhysicName  *string `json:"physic_name,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) SetProduceDate(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) SetProductEntName(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto {
    s.ProductEntName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) SetPackageSpec(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto {
    s.PackageSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) SetProdName(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto {
    s.ProdName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) SetPrepnSpec(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) SetPrepnUnit(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto {
    s.PrepnUnit = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) SetProduceBatchNo(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) SetProdSeqNo(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto {
    s.ProdSeqNo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) SetValidEndDate(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto {
    s.ValidEndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) SetLeastPkgAmount(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto {
    s.LeastPkgAmount = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) SetLeastPrepnAmount(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto {
    s.LeastPrepnAmount = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) SetCodeInfoListDtoList(v []AlibabaAlihealthDrugKytWesUpbillDetailwithcodeCodeInfoListDto) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto {
    s.CodeInfoListDtoList = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) SetDetailBillId(v int64) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto {
    s.DetailBillId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto) SetPhysicName(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeDrugInfosDto {
    s.PhysicName = &v
    return s
}
