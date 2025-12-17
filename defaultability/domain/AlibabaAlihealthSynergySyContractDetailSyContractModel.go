package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyContractDetailSyContractModel struct {
	/*
	   创建时间     */
	GmtCreate *util.LocalTime `json:"gmt_create,omitempty" `

	/*
	   合同编号     */
	ContractNo *string `json:"contract_no,omitempty" `

	/*
	   接收方企业refEntId     */
	ReceiveRefEntId *string `json:"receive_ref_ent_id,omitempty" `

	/*
	   接受方企业entid     */
	ReceiveEntId *string `json:"receive_ent_id,omitempty" `

	/*
	   接受方企业名称     */
	ReceiveEntName *string `json:"receive_ent_name,omitempty" `

	/*
	   发起人     */
	CreatePerson *string `json:"create_person,omitempty" `

	/*
	   收发时间     */
	SendTime *util.LocalTime `json:"send_time,omitempty" `

	/*
	   合同来源：0己方发起 1 对方发起     */
	ContractSource *string `json:"contract_source,omitempty" `

	/*
	   合同来源描述     */
	ContractSourceDesc *string `json:"contract_source_desc,omitempty" `

	/*
	   页数     */
	TotalPage *int64 `json:"total_page,omitempty" `

	/*
	   合同名称     */
	ContractName *string `json:"contract_name,omitempty" `

	/*
	   合同类型：0, "质量保障协议", 1, "销售合同", 2, "采购合同"     */
	ContractType *int64 `json:"contract_type,omitempty" `

	/*
	   合同类型描述     */
	ContractTypeDesc *string `json:"contract_type_desc,omitempty" `

	/*
	   到期时间     */
	ExpireDate *string `json:"expire_date,omitempty" `

	/*
	   发送方企业名称     */
	SendEntName *string `json:"send_ent_name,omitempty" `

	/*
	   签署完成时间     */
	CompleteTime *util.LocalTime `json:"complete_time,omitempty" `

	/*
	   合同状态 ： 0, "暂存", 1, "发起方待选章提交", 2, "发起方审批中", 3, "发起方审批不通过", 4, "发起方待签章发送", 5, "发起方签章失败", 6, "发起方待发送", 7, "接收方待查收", 8, "接收方审批中", 9, "接收方拒签", 10, "接收方待签章", 11, "接收方签章失败", 12, "双方签署成功", 13, "发起方已撤回"     */
	ContractStatus *int64 `json:"contract_status,omitempty" `

	/*
	   合同状态描述     */
	ContractStatusDesc *string `json:"contract_status_desc,omitempty" `

	/*
	   合同id     */
	Id *int64 `json:"id,omitempty" `

	/*
	   发送方企业refEntId     */
	SendRefEntId *string `json:"send_ref_ent_id,omitempty" `

	/*
	   发送方企业entid     */
	SendEntId *string `json:"send_ent_id,omitempty" `
}

func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetGmtCreate(v util.LocalTime) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.GmtCreate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetContractNo(v string) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.ContractNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetReceiveRefEntId(v string) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.ReceiveRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetReceiveEntId(v string) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.ReceiveEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetReceiveEntName(v string) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.ReceiveEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetCreatePerson(v string) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.CreatePerson = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetSendTime(v util.LocalTime) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.SendTime = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetContractSource(v string) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.ContractSource = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetContractSourceDesc(v string) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.ContractSourceDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetTotalPage(v int64) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.TotalPage = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetContractName(v string) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.ContractName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetContractType(v int64) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.ContractType = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetContractTypeDesc(v string) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.ContractTypeDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetExpireDate(v string) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.ExpireDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetSendEntName(v string) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.SendEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetCompleteTime(v util.LocalTime) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.CompleteTime = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetContractStatus(v int64) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.ContractStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetContractStatusDesc(v string) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.ContractStatusDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetId(v int64) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.Id = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetSendRefEntId(v string) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.SendRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyContractDetailSyContractModel) SetSendEntId(v string) *AlibabaAlihealthSynergySyContractDetailSyContractModel {
	s.SendEntId = &v
	return s
}
