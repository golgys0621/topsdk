package domain


type AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistPageInfoDto struct {
    /*
        返回结果     */
    ResultList  *[]AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO `json:"result_list,omitempty" `

    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistPageInfoDto) SetResultList(v []AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistCodeFlowWarningPageResultDTO) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistPageInfoDto {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreelistPageInfoDto {
    s.TotalNum = &v
    return s
}
