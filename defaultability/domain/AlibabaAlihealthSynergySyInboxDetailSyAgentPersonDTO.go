package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO struct {
	/*
	   委托人姓名     */
	Username *string `json:"username,omitempty" `

	/*
	   委托人编号     */
	UserNo *string `json:"user_no,omitempty" `

	/*
	   电话     */
	Tel *string `json:"tel,omitempty" `

	/*
	   身份证号     */
	IdentityNo *string `json:"identity_no,omitempty" `

	/*
	   身份证号有效期至     */
	IdentityExpireDate *util.LocalTime `json:"identity_expire_date,omitempty" `

	/*
	   授权地区     */
	AuthArea *string `json:"auth_area,omitempty" `

	/*
	   授权品种     */
	AuthProduct *string `json:"auth_product,omitempty" `
}

func (s *AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO) SetUsername(v string) *AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO {
	s.Username = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO) SetUserNo(v string) *AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO {
	s.UserNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO) SetTel(v string) *AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO {
	s.Tel = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO) SetIdentityNo(v string) *AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO {
	s.IdentityNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO) SetIdentityExpireDate(v util.LocalTime) *AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO {
	s.IdentityExpireDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO) SetAuthArea(v string) *AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO {
	s.AuthArea = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO) SetAuthProduct(v string) *AlibabaAlihealthSynergySyInboxDetailSyAgentPersonDTO {
	s.AuthProduct = &v
	return s
}
