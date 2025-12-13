package domain


type AlibabaAlihealthDrugKytWesSaveentModel struct {
    /*
        新增失败的时候错误原因     */
    CheckMsg  *string `json:"check_msg,omitempty" `

    /*
        新增成功后分配的往来单位refEntId     */
    ParRefEntId  *string `json:"par_ref_ent_id,omitempty" `

    /*
        新增成功还是失败，true：新增成功     */
    AddSucess  *bool `json:"add_sucess,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesSaveentModel) SetCheckMsg(v string) *AlibabaAlihealthDrugKytWesSaveentModel {
    s.CheckMsg = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSaveentModel) SetParRefEntId(v string) *AlibabaAlihealthDrugKytWesSaveentModel {
    s.ParRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesSaveentModel) SetAddSucess(v bool) *AlibabaAlihealthDrugKytWesSaveentModel {
    s.AddSucess = &v
    return s
}
