package domain


type AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistPageInfoDto struct {
    /*
        返回结果     */
    ResultList  *[]AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistCodeFlowWarningInfoDetailResultDTO `json:"result_list,omitempty" `

    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistPageInfoDto) SetResultList(v []AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistCodeFlowWarningInfoDetailResultDTO) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistPageInfoDto {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugMyjCodewarnfreeCodewarningfreeinfodetaillistPageInfoDto {
    s.TotalNum = &v
    return s
}
