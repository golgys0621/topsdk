package domain


type AlibabaAlihealthDrugMyjCodewarnCodewarninglistPageInfoDto struct {
    /*
        返回结果     */
    ResultList  *[]AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO `json:"result_list,omitempty" `

    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistPageInfoDto) SetResultList(v []AlibabaAlihealthDrugMyjCodewarnCodewarninglistCodeFlowWarningPageResultDTO) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistPageInfoDto {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninglistPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugMyjCodewarnCodewarninglistPageInfoDto {
    s.TotalNum = &v
    return s
}
