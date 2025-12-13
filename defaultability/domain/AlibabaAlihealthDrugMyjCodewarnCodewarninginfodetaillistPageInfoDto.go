package domain


type AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistPageInfoDto struct {
    /*
        返回结果     */
    ResultList  *[]AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistCodeFlowWarningInfoDetailResultDTO `json:"result_list,omitempty" `

    /*
        总计     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistPageInfoDto) SetResultList(v []AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistCodeFlowWarningInfoDetailResultDTO) *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistPageInfoDto {
    s.ResultList = &v
    return s
}
func (s *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistPageInfoDto) SetTotalNum(v int64) *AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistPageInfoDto {
    s.TotalNum = &v
    return s
}
