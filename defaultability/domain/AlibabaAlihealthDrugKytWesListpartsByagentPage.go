package domain


type AlibabaAlihealthDrugKytWesListpartsByagentPage struct {
    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesListpartsByagentPage) SetTotalNum(v int64) *AlibabaAlihealthDrugKytWesListpartsByagentPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesListpartsByagentPage) SetResultList(v []AlibabaAlihealthDrugKytWesListpartsByagentPEntParDto) *AlibabaAlihealthDrugKytWesListpartsByagentPage {
    s.ResultList = &v
    return s
}
