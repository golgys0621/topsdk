package domain


type AlibabaAlihealthDrugMscListpartsEntExtend struct {
    /*
        可替换的entId     */
    ReplaceEntId  *string `json:"replace_ent_id,omitempty" `

}

func (s *AlibabaAlihealthDrugMscListpartsEntExtend) SetReplaceEntId(v string) *AlibabaAlihealthDrugMscListpartsEntExtend {
    s.ReplaceEntId = &v
    return s
}
