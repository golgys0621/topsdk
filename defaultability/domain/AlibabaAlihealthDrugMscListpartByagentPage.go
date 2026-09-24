package domain


type AlibabaAlihealthDrugMscListpartByagentPage struct {
    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

    /*
        返回列表     */
    ResultList  *[]AlibabaAlihealthDrugMscListpartByagentPEntParDto `json:"result_list,omitempty" `

}

func (s *AlibabaAlihealthDrugMscListpartByagentPage) SetTotalNum(v int64) *AlibabaAlihealthDrugMscListpartByagentPage {
    s.TotalNum = &v
    return s
}
func (s *AlibabaAlihealthDrugMscListpartByagentPage) SetResultList(v []AlibabaAlihealthDrugMscListpartByagentPEntParDto) *AlibabaAlihealthDrugMscListpartByagentPage {
    s.ResultList = &v
    return s
}
