package domain


type AlibabaAlihealthDrugMscGetentinfoEntExtend struct {
    /*
        可替换entid     */
    ReplaceEntId  *string `json:"replace_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugMscGetentinfoEntExtend) SetReplaceEntId(v string) *AlibabaAlihealthDrugMscGetentinfoEntExtend {
    s.ReplaceEntId = &v
    return s
}
