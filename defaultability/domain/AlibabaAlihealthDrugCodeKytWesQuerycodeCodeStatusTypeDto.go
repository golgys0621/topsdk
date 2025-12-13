package domain


type AlibabaAlihealthDrugCodeKytWesQuerycodeCodeStatusTypeDto struct {
    /*
        码状态（I核注O核销A激活C注销E错误码B申请未激活）     */
    CodeStatus  *string `json:"code_status,omitempty" `

}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeStatusTypeDto) SetCodeStatus(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeCodeStatusTypeDto {
    s.CodeStatus = &v
    return s
}
