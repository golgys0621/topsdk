package request

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugKytWesQuerycodeactiveRequest struct {
	/*
	   企业唯一标识     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
	LicenseToken *string `json:"license_token" required:"true" `
	/*
	   码列表【多个码时用逗号拼接的字符串。要求数量在4000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】     */
	Codes *[]string `json:"codes" required:"true" `
}

func (s *AlibabaAlihealthDrugKytWesQuerycodeactiveRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesQuerycodeactiveRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesQuerycodeactiveRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesQuerycodeactiveRequest {
	s.LicenseToken = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesQuerycodeactiveRequest) SetCodes(v []string) *AlibabaAlihealthDrugKytWesQuerycodeactiveRequest {
	s.Codes = &v
	return s
}

func (req *AlibabaAlihealthDrugKytWesQuerycodeactiveRequest) ToMap() map[string]interface{} {
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

func (req *AlibabaAlihealthDrugKytWesQuerycodeactiveRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
