package domain


type AlibabaAlihealthDrugKytWesUpbillDetailwithcodeCodeInfoListDto struct {
    /*
        制剂规格     */
    PrepnSpec  *string `json:"prepn_spec,omitempty" `

    /*
        最小制剂数量     */
    PrepnAmount  *string `json:"prepn_amount,omitempty" `

    /*
        最小包装数量     */
    PkgAmount  *string `json:"pkg_amount,omitempty" `

    /*
        码等级【1代表最小码 如：申请的包装比例是1:5:10, 对应的码等级就是3、2、1, 代表大码、中码、小码】     */
    CodeLevel  *string `json:"code_level,omitempty" `

    /*
        监管码     */
    Code  *string `json:"code,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeCodeInfoListDto) SetPrepnSpec(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeCodeInfoListDto {
    s.PrepnSpec = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeCodeInfoListDto) SetPrepnAmount(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeCodeInfoListDto {
    s.PrepnAmount = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeCodeInfoListDto) SetPkgAmount(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeCodeInfoListDto {
    s.PkgAmount = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeCodeInfoListDto) SetCodeLevel(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeCodeInfoListDto {
    s.CodeLevel = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeCodeInfoListDto) SetCode(v string) *AlibabaAlihealthDrugKytWesUpbillDetailwithcodeCodeInfoListDto {
    s.Code = &v
    return s
}
