package request

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugCodeKytWesQuerycodeRequest struct {
	/*
	   企业唯一标识     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
	LicenseToken *string `json:"license_token" required:"true" `
	/*
	   码列表【多个码用逗号分隔的字符串。建议100个以下，数量过多容易超时】     */
	Codes *[]string `json:"codes" required:"true" `
}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeRequest) SetRefEntId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugCodeKytWesQuerycodeRequest {
	s.LicenseToken = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycodeRequest) SetCodes(v []string) *AlibabaAlihealthDrugCodeKytWesQuerycodeRequest {
	s.Codes = &v
	return s
}

func (req *AlibabaAlihealthDrugCodeKytWesQuerycodeRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.LicenseToken != nil {
		paramMap["license_token"] = *req.LicenseToken
	}
	if req.Codes != nil {
		paramMap["codes"] = util.ConvertBasicList(*req.Codes)
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugCodeKytWesQuerycodeRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
