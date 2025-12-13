package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto struct {
	/*
	   处理标志     */
	ProcessFlag *string `json:"process_flag,omitempty" `

	/*
	   状态     */
	Status *string `json:"status,omitempty" `

	/*
	   关联关系类型     */
	RelationType *string `json:"relation_type,omitempty" `

	/*
	   紧急人     */
	UserCert *string `json:"user_cert,omitempty" `

	/*
	   生产编号     */
	ProdCode *string `json:"prod_code,omitempty" `

	/*
	   操作人姓名     */
	OperIcName *string `json:"oper_ic_name,omitempty" `

	/*
	   上传文件名     */
	UploadFileName *string `json:"upload_file_name,omitempty" `

	/*
	   上传文件路径     */
	UploadFilePath *string `json:"upload_file_path,omitempty" `

	/*
	   激活时间     */
	ActiveDate *util.LocalTime `json:"active_date,omitempty" `

	/*
	   企业ID     */
	RefEntId *string `json:"ref_ent_id,omitempty" `

	/*
	   旧企业ID     */
	EntId *string `json:"ent_id,omitempty" `

	/*
	   处理日期     */
	ProcessDate *util.LocalTime `json:"process_date,omitempty" `

	/*
	   处理结束时间     */
	ProcessEndDate *util.LocalTime `json:"process_end_date,omitempty" `

	/*
	   操作人编码     */
	OperIcCode *string `json:"oper_ic_code,omitempty" `

	/*
	   上传标识     */
	UploadFlag *string `json:"upload_flag,omitempty" `

	/*
	   激活时间     */
	CrtDate *util.LocalTime `json:"crt_date,omitempty" `

	/*
	   处理数量     */
	ProcessCount *string `json:"process_count,omitempty" `

	/*
	   总激活数量     */
	ActiveCount *int64 `json:"active_count,omitempty" `

	/*
	   最大包装数量     */
	OtherNum *int64 `json:"other_num,omitempty" `

	/*
	   小码数量     */
	SmallNum *int64 `json:"small_num,omitempty" `

	/*
	   关联关系文件上传日期     */
	CrtDateString *string `json:"crt_date_string,omitempty" `

	/*
	   单据id     */
	BillInId *string `json:"bill_in_id,omitempty" `

	/*
	   激活信息id     */
	CodeActiveInfoId *string `json:"code_active_info_id,omitempty" `
}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetProcessFlag(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.ProcessFlag = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetStatus(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.Status = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetRelationType(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.RelationType = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetUserCert(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.UserCert = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetProdCode(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.ProdCode = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetOperIcName(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.OperIcName = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetUploadFileName(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.UploadFileName = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetUploadFilePath(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.UploadFilePath = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetActiveDate(v util.LocalTime) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.ActiveDate = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetRefEntId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetEntId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.EntId = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetProcessDate(v util.LocalTime) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.ProcessDate = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetProcessEndDate(v util.LocalTime) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.ProcessEndDate = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetOperIcCode(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.OperIcCode = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetUploadFlag(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.UploadFlag = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetCrtDate(v util.LocalTime) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.CrtDate = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetProcessCount(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.ProcessCount = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetActiveCount(v int64) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.ActiveCount = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetOtherNum(v int64) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.OtherNum = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetSmallNum(v int64) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.SmallNum = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetCrtDateString(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.CrtDateString = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetBillInId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.BillInId = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto) SetCodeActiveInfoId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationCodeActiveInfoDto {
	s.CodeActiveInfoId = &v
	return s
}
