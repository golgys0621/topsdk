package domain


type AlibabaAlihealthDrugCodeKytWesQuerycoderelationPkgInfoDto struct {
    /*
        码信息     */
    CodeList  *[]string `json:"code_list,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationPkgInfoDto) SetCodeList(v []string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationPkgInfoDto {
    s.CodeList = &v
    return s
}
