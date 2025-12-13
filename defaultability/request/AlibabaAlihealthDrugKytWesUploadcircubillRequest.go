package request

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugKytWesUploadcircubillRequest struct {
	/*
	   货主企业（单据的所有者，上传人）【注意：该入参是ref_ent_id，不是ent_id】     */
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
	   药品类型【3普药2特药】不再以该写字为标准，任意填写2或3即可。     */
	PhysicType *int64 `json:"physic_type" required:"true" `
	/*
	   第三方物流代理企业【注意：该入参是ref_ent_id，不是ent_id】，该字段兼容之前接口逻辑，后期将不允许使用，不要填值。     */
	AgentRefUserId *string `json:"agent_ref_user_id,omitempty" required:"false" `
	/*
	   发货企业【注意：该入参是ent_id,并不是ref_ent_id】     */
	FromUserId *string `json:"from_user_id" required:"true" `
	/*
	   收货企业【注意：该入参是ent_id,并不是ref_ent_id】     */
	ToUserId *string `json:"to_user_id" required:"true" `
	/*
	   直调企业【注意：该入参是ent_id,并不是ref_ent_id】     */
	DestUserId *string `json:"dest_user_id,omitempty" required:"false" `
	/*
	   操作人标识（写appkey编号）     */
	OperIcCode *string `json:"oper_ic_code" required:"true" `
	/*
	   单据提交者姓名     */
	OperIcName *string `json:"oper_ic_name,omitempty" required:"false" `
	/*
	   单据文件体【bas64字符串】，看对接文档中的代码示例，示例中有相应说明。     */
	FileContent *string `json:"file_content" required:"true" `
	/*
	   单据文件名称     */
	UploadFileName *string `json:"upload_file_name" required:"true" `
	/*
	   请求端类型【暂定都写2】     */
	ClientType *string `json:"client_type" required:"true" `
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
	   药品配送企业refentid【出库单填写，与dis_ent_id入参选其一添加】     */
	DisRefEntId *string `json:"dis_ref_ent_id,omitempty" required:"false" `
	/*
	   药品配送企业entId【出库单填写】     */
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
	   【可不填】药品列表Json："codeCount":         药品数量 "commDrugId":     国家药品唯一标识 "exprieDate":         生产日期 "physicInfo":          药品信息 "pkgSpec":           包状规格 "prepnCount":       制剂数量 "produceBatchNo":生产批次 "produceDate":      生产日期     */
	DrugListJson *string `json:"drug_list_json,omitempty" required:"false" `
	/*
	   单据委托企业refEntId     */
	AssRefEntId *string `json:"ass_ref_ent_id,omitempty" required:"false" `
	/*
	   单据委托企业refEntId     */
	AssEntId *string `json:"ass_ent_id,omitempty" required:"false" `
	/*
	   码解析策略,1代表整单解析成功(任一码解析失败，上传时整单返回错误),传其他值或者不传代表部分解析成功(跳过无法解析的码，其余正常处理并上传)     */
	IgnorePartSuccessFlag *string `json:"ignore_part_success_flag,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.LicenseToken = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetBillCode(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.BillCode = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetBillTime(v util.LocalTime) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.BillTime = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetBillType(v int64) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.BillType = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetPhysicType(v int64) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.PhysicType = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetAgentRefUserId(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.AgentRefUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetFromUserId(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.FromUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetToUserId(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.ToUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetDestUserId(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.DestUserId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetOperIcCode(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.OperIcCode = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetOperIcName(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.OperIcName = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetFileContent(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.FileContent = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetUploadFileName(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.UploadFileName = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetClientType(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.ClientType = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetFromAddress(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.FromAddress = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetToAddress(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.ToAddress = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetFromBillCode(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.FromBillCode = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetOrderCode(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.OrderCode = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetFromPerson(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.FromPerson = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetToPerson(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.ToPerson = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetDisRefEntId(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.DisRefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetDisEntId(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.DisEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetQuReceivable(v int64) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.QuReceivable = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetXtIsCheck(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.XtIsCheck = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetXtCheckCode(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.XtCheckCode = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetXtCheckCodeDesc(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.XtCheckCodeDesc = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetDrugListJson(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.DrugListJson = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetAssRefEntId(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.AssRefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetAssEntId(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.AssEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesUploadcircubillRequest) SetIgnorePartSuccessFlag(v string) *AlibabaAlihealthDrugKytWesUploadcircubillRequest {
	s.IgnorePartSuccessFlag = &v
	return s
}

func (req *AlibabaAlihealthDrugKytWesUploadcircubillRequest) ToMap() map[string]interface{} {
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
	if req.FileContent != nil {
		paramMap["file_content"] = *req.FileContent
	}
	if req.UploadFileName != nil {
		paramMap["upload_file_name"] = *req.UploadFileName
	}
	if req.ClientType != nil {
		paramMap["client_type"] = *req.ClientType
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

func (req *AlibabaAlihealthDrugKytWesUploadcircubillRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
