package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugWesGetentinfolistRequest struct {
	/*
	   refEntId     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   licenseToken     */
	LicenseToken *string `json:"license_token" required:"true" `
	/*
	   查询企业信息参数     */
	QueryParam *domain.AlibabaAlihealthDrugWesGetentinfolistTopEntInfoReqDto `json:"query_param,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugWesGetentinfolistRequest) SetRefEntId(v string) *AlibabaAlihealthDrugWesGetentinfolistRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugWesGetentinfolistRequest {
	s.LicenseToken = &v
	return s
}
func (s *AlibabaAlihealthDrugWesGetentinfolistRequest) SetQueryParam(v domain.AlibabaAlihealthDrugWesGetentinfolistTopEntInfoReqDto) *AlibabaAlihealthDrugWesGetentinfolistRequest {
	s.QueryParam = &v
	return s
}

func (req *AlibabaAlihealthDrugWesGetentinfolistRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.LicenseToken != nil {
		paramMap["license_token"] = *req.LicenseToken
	}
	if req.QueryParam != nil {
		paramMap["query_param"] = util.ConvertStruct(*req.QueryParam)
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugWesGetentinfolistRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
