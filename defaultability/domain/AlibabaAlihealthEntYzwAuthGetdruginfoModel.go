package domain


type AlibabaAlihealthEntYzwAuthGetdruginfoModel struct {
    /*
        总记录数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        药品集合     */
    Result  *[]AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO `json:"result,omitempty" `

}

func (s *AlibabaAlihealthEntYzwAuthGetdruginfoModel) SetTotalNum(v int64) *AlibabaAlihealthEntYzwAuthGetdruginfoModel {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthEntYzwAuthGetdruginfoModel) SetResult(v []AlibabaAlihealthEntYzwAuthGetdruginfoDrugEntBaseInfoDTO) *AlibabaAlihealthEntYzwAuthGetdruginfoModel {
    s.Result = &v
    return s
}
