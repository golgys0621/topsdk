package request

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyAgentpersonSaveRequest struct {
	/*
	   refEntId     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   委托人姓名     */
	Username *string `json:"username" required:"true" `
	/*
	   委托人编号     */
	UserNo *string `json:"user_no" required:"true" `
	/*
	   身份证号     */
	IdentityNo *string `json:"identity_no" required:"true" `
	/*
	   身份证号有效期至     */
	IdentityExpireDate *string `json:"identity_expire_date" required:"true" `
	/*
	   授权区域     */
	AuthArea *string `json:"auth_area" required:"true" `
	/*
	   授权品种     */
	AuthProduct *string `json:"auth_product" required:"true" `
	/*
	   联系电话     */
	Tel *string `json:"tel,omitempty" required:"false" `
	/*
	   委托人分类,0销售、1采购、2其他     */
	Type *int64 `json:"type" required:"true" `
	/*
	   委托结束日期     */
	AuthExpireDate *util.LocalTime `json:"auth_expire_date" required:"true" `
	/*
	   身份证人像图片，仅支持png     */
	IdHeadPngFileContent *[]byte `json:"id_head_png_file_content,omitempty" required:"false" `
	/*
	   身份证国徽图片，仅支持png     */
	IdBackPngFileContent *[]byte `json:"id_back_png_file_content,omitempty" required:"false" `
	/*
	   委托开始日期     */
	AuthStartDate *util.LocalTime `json:"auth_start_date" required:"true" `
}

func (s *AlibabaAlihealthSynergySyAgentpersonSaveRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyAgentpersonSaveRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonSaveRequest) SetUsername(v string) *AlibabaAlihealthSynergySyAgentpersonSaveRequest {
	s.Username = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonSaveRequest) SetUserNo(v string) *AlibabaAlihealthSynergySyAgentpersonSaveRequest {
	s.UserNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonSaveRequest) SetIdentityNo(v string) *AlibabaAlihealthSynergySyAgentpersonSaveRequest {
	s.IdentityNo = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonSaveRequest) SetIdentityExpireDate(v string) *AlibabaAlihealthSynergySyAgentpersonSaveRequest {
	s.IdentityExpireDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonSaveRequest) SetAuthArea(v string) *AlibabaAlihealthSynergySyAgentpersonSaveRequest {
	s.AuthArea = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonSaveRequest) SetAuthProduct(v string) *AlibabaAlihealthSynergySyAgentpersonSaveRequest {
	s.AuthProduct = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonSaveRequest) SetTel(v string) *AlibabaAlihealthSynergySyAgentpersonSaveRequest {
	s.Tel = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonSaveRequest) SetType(v int64) *AlibabaAlihealthSynergySyAgentpersonSaveRequest {
	s.Type = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonSaveRequest) SetAuthExpireDate(v util.LocalTime) *AlibabaAlihealthSynergySyAgentpersonSaveRequest {
	s.AuthExpireDate = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonSaveRequest) SetIdHeadPngFileContent(v []byte) *AlibabaAlihealthSynergySyAgentpersonSaveRequest {
	s.IdHeadPngFileContent = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonSaveRequest) SetIdBackPngFileContent(v []byte) *AlibabaAlihealthSynergySyAgentpersonSaveRequest {
	s.IdBackPngFileContent = &v
	return s
}
func (s *AlibabaAlihealthSynergySyAgentpersonSaveRequest) SetAuthStartDate(v util.LocalTime) *AlibabaAlihealthSynergySyAgentpersonSaveRequest {
	s.AuthStartDate = &v
	return s
}

func (req *AlibabaAlihealthSynergySyAgentpersonSaveRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.Username != nil {
		paramMap["username"] = *req.Username
	}
	if req.UserNo != nil {
		paramMap["user_no"] = *req.UserNo
	}
	if req.IdentityNo != nil {
		paramMap["identity_no"] = *req.IdentityNo
	}
	if req.IdentityExpireDate != nil {
		paramMap["identity_expire_date"] = *req.IdentityExpireDate
	}
	if req.AuthArea != nil {
		paramMap["auth_area"] = *req.AuthArea
	}
	if req.AuthProduct != nil {
		paramMap["auth_product"] = *req.AuthProduct
	}
	if req.Tel != nil {
		paramMap["tel"] = *req.Tel
	}
	if req.Type != nil {
		paramMap["type"] = *req.Type
	}
	if req.AuthExpireDate != nil {
		paramMap["auth_expire_date"] = *req.AuthExpireDate
	}
	if req.AuthStartDate != nil {
		paramMap["auth_start_date"] = *req.AuthStartDate
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergySyAgentpersonSaveRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	if req.IdHeadPngFileContent != nil {
		fileMap["id_head_png_file_content"] = *req.IdHeadPngFileContent
	}
	if req.IdBackPngFileContent != nil {
		fileMap["id_back_png_file_content"] = *req.IdBackPngFileContent
	}
	return fileMap
}
