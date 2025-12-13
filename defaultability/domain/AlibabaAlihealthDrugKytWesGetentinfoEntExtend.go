package domain


type AlibabaAlihealthDrugKytWesGetentinfoEntExtend struct {
    /*
        可替换entid     */
    ReplaceEntId  *string `json:"replace_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugKytWesGetentinfoEntExtend) SetReplaceEntId(v string) *AlibabaAlihealthDrugKytWesGetentinfoEntExtend {
    s.ReplaceEntId = &v
    return s
}
