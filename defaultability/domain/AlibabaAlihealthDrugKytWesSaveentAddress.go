package domain


type AlibabaAlihealthDrugKytWesSaveentAddress struct {
    /*
        境内填写区县名称/境外则填写境外国家中文名称     */
    AreaName  *string `json:"area_name,omitempty" `

    /*
        城市名称/境外不用填，境内必填     */
    CityName  *string `json:"city_name,omitempty" `

    /*
        省份名称/境外不用填，境内必填     */
    ProvName  *string `json:"prov_name,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSaveentAddress) SetAreaName(v string) *AlibabaAlihealthDrugKytWesSaveentAddress {
    s.AreaName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSaveentAddress) SetCityName(v string) *AlibabaAlihealthDrugKytWesSaveentAddress {
    s.CityName = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSaveentAddress) SetProvName(v string) *AlibabaAlihealthDrugKytWesSaveentAddress {
    s.ProvName = &v
    return s
}
