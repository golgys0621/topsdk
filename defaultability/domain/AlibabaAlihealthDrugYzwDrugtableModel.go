package domain


type AlibabaAlihealthDrugYzwDrugtableModel struct {
    /*
        总记录数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        药品集合     */
    Result  *[]AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO `json:"result,omitempty" `

}

func (s *AlibabaAlihealthDrugYzwDrugtableModel) SetTotalNum(v int64) *AlibabaAlihealthDrugYzwDrugtableModel {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugYzwDrugtableModel) SetResult(v []AlibabaAlihealthDrugYzwDrugtableDrugEntBaseInfoDTO) *AlibabaAlihealthDrugYzwDrugtableModel {
    s.Result = &v
    return s
}
