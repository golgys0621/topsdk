package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO struct {
	/*
	   单据类型     */
	BillType *int64 `json:"bill_type,omitempty" `

	/*
	   单据ID     */
	BillId *string `json:"bill_id,omitempty" `

	/*
	   单据明细ID     */
	BillDetailId *string `json:"bill_detail_id,omitempty" `

	/*
	   单据时间     */
	BillTime *string `json:"bill_time,omitempty" `

	/*
	   包装     */
	PkgSpec *string `json:"pkg_spec,omitempty" `

	/*
	   制剂     */
	PrepnSpec *string `json:"prepn_spec,omitempty" `

	/*
	   药品名     */
	PhysicName *string `json:"physic_name,omitempty" `

	/*
	   剂型     */
	PrepnTypeDesc *string `json:"prepn_type_desc,omitempty" `

	/*
	   报告状态(0 ：待发送  2：待签收 3：已签收 4：已拒绝 7:对方已签收(更正待处理)  13:对方已拒绝(更正待签收))     */
	DrugReportSignStatus *string `json:"drug_report_sign_status,omitempty" `

	/*
	   单据编码     */
	BillCode *string `json:"bill_code,omitempty" `

	/*
	   批号     */
	ProduceBatchNo *string `json:"produce_batch_no,omitempty" `

	/*
	   药品ID     */
	DrugId *string `json:"drug_id,omitempty" `

	/*
	   生产日期     */
	ProduceDate *string `json:"produce_date,omitempty" `

	/*
	   生产企业ID     */
	ProduceEntId *string `json:"produce_ent_id,omitempty" `

	/*
	   委托企业     */
	AssRefEntId *string `json:"ass_ref_ent_id,omitempty" `

	/*
	   报告名称     */
	FileName *string `json:"file_name,omitempty" `

	/*
	   签章位置     */
	SealSignatures *[]AlibabaAlihealthSynergyYzwQuerydrugreportassSealSignatureDTO `json:"seal_signatures,omitempty" `

	/*
	   签章URL     */
	SealedReportUrl *string `json:"sealed_report_url,omitempty" `

	/*
	   发货企业     */
	FromRefEntId *string `json:"from_ref_ent_id,omitempty" `

	/*
	   发货企业     */
	FromEntName *string `json:"from_ent_name,omitempty" `

	/*
	   单据上传日期     */
	CrtDate *util.LocalTime `json:"crt_date,omitempty" `

	/*
	   签收时间     */
	SignedTime *string `json:"signed_time,omitempty" `

	/*
	   盖章状态 0:待盖章 5:已盖章 6:同批号已盖章      */
	SealStatus *string `json:"seal_status,omitempty" `

	/*
	   报告ID     */
	DrugReportId *int64 `json:"drug_report_id,omitempty" `

	/*
	   生产企业     */
	ProduceEntName *string `json:"produce_ent_name,omitempty" `
}

func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetBillType(v int64) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.BillType = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetBillId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.BillId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetBillDetailId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.BillDetailId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetBillTime(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.BillTime = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetPkgSpec(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.PkgSpec = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetPrepnSpec(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.PrepnSpec = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetPhysicName(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.PhysicName = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetPrepnTypeDesc(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.PrepnTypeDesc = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetDrugReportSignStatus(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.DrugReportSignStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetBillCode(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.BillCode = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetProduceBatchNo(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.ProduceBatchNo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetDrugId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.DrugId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetProduceDate(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.ProduceDate = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetProduceEntId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.ProduceEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetAssRefEntId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.AssRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetFileName(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.FileName = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetSealSignatures(v []AlibabaAlihealthSynergyYzwQuerydrugreportassSealSignatureDTO) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.SealSignatures = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetSealedReportUrl(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.SealedReportUrl = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetFromRefEntId(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.FromRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetFromEntName(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.FromEntName = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetCrtDate(v util.LocalTime) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.CrtDate = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetSignedTime(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.SignedTime = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetSealStatus(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.SealStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetDrugReportId(v int64) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.DrugReportId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO) SetProduceEntName(v string) *AlibabaAlihealthSynergyYzwQuerydrugreportassOnenetDrugReportTopDTO {
	s.ProduceEntName = &v
	return s
}
