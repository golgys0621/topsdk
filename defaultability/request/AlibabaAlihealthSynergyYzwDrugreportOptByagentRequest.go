package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest struct {
	/*
	   代理操作企业refEntId(调用企业refEntId)     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   用户id     */
	UserId *string `json:"user_id" required:"true" `
	/*
	   操作类型：签收(onlySign),拒收(rejectReceive),签收并盖章(signAndSeal),盖章签收后的报告(sealAfterSign)     */
	OperType *string `json:"oper_type" required:"true" `
	/*
	   单据id     */
	BillId *int64 `json:"bill_id" required:"true" `
	/*
	   单据明细id     */
	DetailBillId *int64 `json:"detail_bill_id" required:"true" `
	/*
	   单据类型     */
	BillType *int64 `json:"bill_type" required:"true" `
	/*
	   拒绝原因(当操作类型为拒收(rejectReceive)需要该字段)     */
	OperNote *string `json:"oper_note,omitempty" required:"false" `
	/*
	   授权单位企业id     */
	AssRefEntId *string `json:"ass_ref_ent_id" required:"true" `
	/*
	   代理操作企业(调用企业)的章信息     */
	SelfSealInfo *domain.AlibabaAlihealthSynergyYzwDrugreportOptByagentDrugReportSealInfo `json:"self_seal_info,omitempty" required:"false" `
	/*
	   授权单位的章信息     */
	AssSealInfo *domain.AlibabaAlihealthSynergyYzwDrugreportOptByagentDrugReportSealInfo `json:"ass_seal_info,omitempty" required:"false" `
	/*
	   onlySelf(仅盖自己的章),onlyAuthorized(仅盖授权单位的章),both(自己和授权单位的章);当操作为签收并盖章(signAndSeal),盖章签收后的报告(sealAfterSign)需要该字段     */
	SealMode *string `json:"seal_mode,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest) SetUserId(v string) *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest {
	s.UserId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest) SetOperType(v string) *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest {
	s.OperType = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest) SetBillId(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest {
	s.BillId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest) SetDetailBillId(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest {
	s.DetailBillId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest) SetBillType(v int64) *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest {
	s.BillType = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest) SetOperNote(v string) *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest {
	s.OperNote = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest) SetAssRefEntId(v string) *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest {
	s.AssRefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest) SetSelfSealInfo(v domain.AlibabaAlihealthSynergyYzwDrugreportOptByagentDrugReportSealInfo) *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest {
	s.SelfSealInfo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest) SetAssSealInfo(v domain.AlibabaAlihealthSynergyYzwDrugreportOptByagentDrugReportSealInfo) *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest {
	s.AssSealInfo = &v
	return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest) SetSealMode(v string) *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest {
	s.SealMode = &v
	return s
}

func (req *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.UserId != nil {
		paramMap["user_id"] = *req.UserId
	}
	if req.OperType != nil {
		paramMap["oper_type"] = *req.OperType
	}
	if req.BillId != nil {
		paramMap["bill_id"] = *req.BillId
	}
	if req.DetailBillId != nil {
		paramMap["detail_bill_id"] = *req.DetailBillId
	}
	if req.BillType != nil {
		paramMap["bill_type"] = *req.BillType
	}
	if req.OperNote != nil {
		paramMap["oper_note"] = *req.OperNote
	}
	if req.AssRefEntId != nil {
		paramMap["ass_ref_ent_id"] = *req.AssRefEntId
	}
	if req.SelfSealInfo != nil {
		paramMap["self_seal_info"] = util.ConvertStruct(*req.SelfSealInfo)
	}
	if req.AssSealInfo != nil {
		paramMap["ass_seal_info"] = util.ConvertStruct(*req.AssSealInfo)
	}
	if req.SealMode != nil {
		paramMap["seal_mode"] = *req.SealMode
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
