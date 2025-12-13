package request

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugKytWesUploadinoutbillRequest struct {
	/*
	   货主（单据的所有者，上传人），上传企业的单位维一编码【注意：该入参是ref_ent_id，不是ent_id】     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
	LicenseToken *string `json:"license_token" required:"true" `
	/*
	   单据编号【同一个企业不能上传相同单据号】     */
	BillCode *string `json:"bill_code" required:"true" `
	/*
	   单据时间（扫码时间）     */
	BillTime *util.LocalTime `json:"bill_time" required:"true" `
	/*
	   单据类型【102代表采购入库,201代表销售出库，其它单据类型详见文档】     */
	BillType *int64 `json:"bill_type" required:"true" `
	/*
	   药品类型【3普药2特药】【可以随便填写，单据上传后会以实际为准】     */
	PhysicType *int64 `json:"physic_type" required:"true" `
	/*
	   第三方物流代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。     */
	AgentRefUserId *string `json:"agent_ref_user_id,omitempty" required:"false" `
	/*
	   发货企业entId【注意：该入参是ent_id,并不是ref_ent_id】     */
	FromUserId *string `json:"from_user_id" required:"true" `
	/*
	   收货企业entId【注意：该入参是ent_id,并不是ref_ent_id】     */
	ToUserId *string `json:"to_user_id" required:"true" `
	/*
	   直调企业标识【注意：该入参是ent_id,并不是ref_ent_id】     */
	DestUserId *string `json:"dest_user_id,omitempty" required:"false" `
	/*
	   单据提交者（调用接口时的appkey编号）     */
	OperIcCode *string `json:"oper_ic_code" required:"true" `
	/*
	   单据提交者姓名（出入库单上传人的名子）     */
	OperIcName *string `json:"oper_ic_name" required:"true" `
	/*
	   仓号     */
	WarehouseId *string `json:"warehouse_id,omitempty" required:"false" `
	/*
	   已经废弃     */
	DrugId *string `json:"drug_id,omitempty" required:"false" `
	/*
	   追溯码【多个码时用逗号拼接的字符串。要求数量在1万个码以下，允许多个药品多个码混合上传】     */
	TraceCodes *[]string `json:"trace_codes" required:"true" `
	/*
	   调用方类型[必须填2]     */
	ClientType *string `json:"client_type" required:"true" `
	/*
	   已废弃，无需填写     */
	ReturnReasonCode *string `json:"return_reason_code,omitempty" required:"false" `
	/*
	   已废弃，无需填写     */
	ReturnReasonDes *string `json:"return_reason_des,omitempty" required:"false" `
	/*
	   已废弃，无需填写     */
	CancelReasonCode *string `json:"cancel_reason_code,omitempty" required:"false" `
	/*
	   已废弃，无需填写     */
	CancelReasonDes *string `json:"cancel_reason_des,omitempty" required:"false" `
	/*
	   已废弃，无需填写     */
	ExecuterName *string `json:"executer_name,omitempty" required:"false" `
	/*
	   已废弃，无需填写     */
	ExecuterCode *string `json:"executer_code,omitempty" required:"false" `
	/*
	   已废弃，无需填写     */
	SuperviserName *string `json:"superviser_name,omitempty" required:"false" `
	/*
	   已废弃，无需填写     */
	SuperviserCode *string `json:"superviser_code,omitempty" required:"false" `
	/*
	   发货地址     */
	FromAddress *string `json:"from_address,omitempty" required:"false" `
	/*
	   收货地址     */
	ToAddress *string `json:"to_address,omitempty" required:"false" `
	/*
	   发货单编号     */
	FromBillCode *string `json:"from_bill_code,omitempty" required:"false" `
	/*
	   订货单编号     */
	OrderCode *string `json:"order_code,omitempty" required:"false" `
	/*
	   发货人     */
	FromPerson *string `json:"from_person,omitempty" required:"false" `
	/*
	   收货人     */
	ToPerson *string `json:"to_person,omitempty" required:"false" `
	/*
	   药品配送ref_enti_d企业     */
	DisRefEntId *string `json:"dis_ref_ent_id,omitempty" required:"false" `
	/*
	   药品配送企业ent_Id(出库单)     */
	DisEntId *string `json:"dis_ent_id,omitempty" required:"false" `
	/*
	   应收货总数量     */
	QuReceivable *int64 `json:"qu_receivable,omitempty" required:"false" `
	/*
	   是否验证，0：未通过验证，1：已验证     */
	XtIsCheck *string `json:"xt_is_check,omitempty" required:"false" `
	/*
	   未验证通过原因【验证未通过时填写】     */
	XtCheckCode *string `json:"xt_check_code,omitempty" required:"false" `
	/*
	   未验证通过原因描述【验证未通过时填写】     */
	XtCheckCodeDesc *string `json:"xt_check_code_desc,omitempty" required:"false" `
	/*
	   药品列表Json："codeCount": 药品数量 "commDrugId": 国家药品唯一标识 "exprieDate": 生产日期 "physicInfo": 药品信息 "pkgSpec": 包状规格 "prepnCount": 制剂数量 "produceBatchNo":生产批次 "produceDate": 生产日期     */
	DrugListJson *string `json:"drug_list_json,omitempty" required:"false" `
	/*
	   单据委托企业refEntId     */
	AssRefEntId *string `json:"ass_ref_ent_id,omitempty" required:"false" `
	/*
	   单据委托企业entId     */
	AssEntId *string `json:"ass_ent_id,omitempty" required:"false" `
	/*
	   码解析策略,1代表整单解析成功(任一码解析失败，上传时整单返回错误),传其他值或者不传代表部分解析成功(跳过无法解析的码，其余正常处理并上传)     */
	IgnorePartSuccessFlag *string `json:"ignore_part_success_flag,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.LicenseToken = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetBillCode(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.BillCode = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetBillTime(v util.LocalTime) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.BillTime = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetBillType(v int64) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.BillType = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetPhysicType(v int64) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.PhysicType = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetAgentRefUserId(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.AgentRefUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetFromUserId(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.FromUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetToUserId(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.ToUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetDestUserId(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.DestUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetOperIcCode(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.OperIcCode = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetOperIcName(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.OperIcName = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetWarehouseId(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.WarehouseId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetDrugId(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.DrugId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetTraceCodes(v []string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.TraceCodes = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetClientType(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.ClientType = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetReturnReasonCode(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.ReturnReasonCode = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetReturnReasonDes(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.ReturnReasonDes = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetCancelReasonCode(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.CancelReasonCode = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetCancelReasonDes(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.CancelReasonDes = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetExecuterName(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.ExecuterName = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetExecuterCode(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.ExecuterCode = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetSuperviserName(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.SuperviserName = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetSuperviserCode(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.SuperviserCode = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetFromAddress(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.FromAddress = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetToAddress(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.ToAddress = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetFromBillCode(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.FromBillCode = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetOrderCode(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.OrderCode = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetFromPerson(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.FromPerson = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetToPerson(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.ToPerson = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetDisRefEntId(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.DisRefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetDisEntId(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.DisEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetQuReceivable(v int64) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.QuReceivable = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetXtIsCheck(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.XtIsCheck = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetXtCheckCode(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.XtCheckCode = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetXtCheckCodeDesc(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.XtCheckCodeDesc = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetDrugListJson(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.DrugListJson = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetAssRefEntId(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.AssRefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetAssEntId(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.AssEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) SetIgnorePartSuccessFlag(v string) *AlibabaAlihealthDrugKytWesUploadinoutbillRequest {
	s.IgnorePartSuccessFlag = &v
	return s
}

func (req *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.LicenseToken != nil {
		paramMap["license_token"] = *req.LicenseToken
	}
	if req.BillCode != nil {
		paramMap["bill_code"] = *req.BillCode
	}
	if req.BillTime != nil {
		paramMap["bill_time"] = *req.BillTime
	}
	if req.BillType != nil {
		paramMap["bill_type"] = *req.BillType
	}
	if req.PhysicType != nil {
		paramMap["physic_type"] = *req.PhysicType
	}
	if req.AgentRefUserId != nil {
		paramMap["agent_ref_user_id"] = *req.AgentRefUserId
	}
	if req.FromUserId != nil {
		paramMap["from_user_id"] = *req.FromUserId
	}
	if req.ToUserId != nil {
		paramMap["to_user_id"] = *req.ToUserId
	}
	if req.DestUserId != nil {
		paramMap["dest_user_id"] = *req.DestUserId
	}
	if req.OperIcCode != nil {
		paramMap["oper_ic_code"] = *req.OperIcCode
	}
	if req.OperIcName != nil {
		paramMap["oper_ic_name"] = *req.OperIcName
	}
	if req.WarehouseId != nil {
		paramMap["warehouse_id"] = *req.WarehouseId
	}
	if req.DrugId != nil {
		paramMap["drug_id"] = *req.DrugId
	}
	if req.TraceCodes != nil {
		paramMap["trace_codes"] = util.ConvertBasicList(*req.TraceCodes)
	}
	if req.ClientType != nil {
		paramMap["client_type"] = *req.ClientType
	}
	if req.ReturnReasonCode != nil {
		paramMap["return_reason_code"] = *req.ReturnReasonCode
	}
	if req.ReturnReasonDes != nil {
		paramMap["return_reason_des"] = *req.ReturnReasonDes
	}
	if req.CancelReasonCode != nil {
		paramMap["cancel_reason_code"] = *req.CancelReasonCode
	}
	if req.CancelReasonDes != nil {
		paramMap["cancel_reason_des"] = *req.CancelReasonDes
	}
	if req.ExecuterName != nil {
		paramMap["executer_name"] = *req.ExecuterName
	}
	if req.ExecuterCode != nil {
		paramMap["executer_code"] = *req.ExecuterCode
	}
	if req.SuperviserName != nil {
		paramMap["superviser_name"] = *req.SuperviserName
	}
	if req.SuperviserCode != nil {
		paramMap["superviser_code"] = *req.SuperviserCode
	}
	if req.FromAddress != nil {
		paramMap["from_address"] = *req.FromAddress
	}
	if req.ToAddress != nil {
		paramMap["to_address"] = *req.ToAddress
	}
	if req.FromBillCode != nil {
		paramMap["from_bill_code"] = *req.FromBillCode
	}
	if req.OrderCode != nil {
		paramMap["order_code"] = *req.OrderCode
	}
	if req.FromPerson != nil {
		paramMap["from_person"] = *req.FromPerson
	}
	if req.ToPerson != nil {
		paramMap["to_person"] = *req.ToPerson
	}
	if req.DisRefEntId != nil {
		paramMap["dis_ref_ent_id"] = *req.DisRefEntId
	}
	if req.DisEntId != nil {
		paramMap["dis_ent_id"] = *req.DisEntId
	}
	if req.QuReceivable != nil {
		paramMap["qu_receivable"] = *req.QuReceivable
	}
	if req.XtIsCheck != nil {
		paramMap["xt_is_check"] = *req.XtIsCheck
	}
	if req.XtCheckCode != nil {
		paramMap["xt_check_code"] = *req.XtCheckCode
	}
	if req.XtCheckCodeDesc != nil {
		paramMap["xt_check_code_desc"] = *req.XtCheckCodeDesc
	}
	if req.DrugListJson != nil {
		paramMap["drug_list_json"] = *req.DrugListJson
	}
	if req.AssRefEntId != nil {
		paramMap["ass_ref_ent_id"] = *req.AssRefEntId
	}
	if req.AssEntId != nil {
		paramMap["ass_ent_id"] = *req.AssEntId
	}
	if req.IgnorePartSuccessFlag != nil {
		paramMap["ignore_part_success_flag"] = *req.IgnorePartSuccessFlag
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugKytWesUploadinoutbillRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
