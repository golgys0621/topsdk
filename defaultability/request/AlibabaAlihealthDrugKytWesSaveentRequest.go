package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugKytWesSaveentRequest struct {
	/*
	   添加企业唯一标识     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
	LicenseToken *string `json:"license_token" required:"true" `
	/*
	   新增企业信息     */
	AddEntReq *domain.AlibabaAlihealthDrugKytWesSaveentAddEntReqDTO `json:"add_ent_req" required:"true" `
	/*
	   图片数据流。图片大小务必控制在2M以内     */
	LicPictureByte *[]byte `json:"lic_picture_byte" required:"true" `
}

func (s *AlibabaAlihealthDrugKytWesSaveentRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesSaveentRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesSaveentRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesSaveentRequest {
	s.LicenseToken = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesSaveentRequest) SetAddEntReq(v domain.AlibabaAlihealthDrugKytWesSaveentAddEntReqDTO) *AlibabaAlihealthDrugKytWesSaveentRequest {
	s.AddEntReq = &v
	return s
}
func (s *AlibabaAlihealthDrugKytWesSaveentRequest) SetLicPictureByte(v []byte) *AlibabaAlihealthDrugKytWesSaveentRequest {
	s.LicPictureByte = &v
	return s
}

func (req *AlibabaAlihealthDrugKytWesSaveentRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.LicenseToken != nil {
		paramMap["license_token"] = *req.LicenseToken
	}
	if req.AddEntReq != nil {
		paramMap["add_ent_req"] = util.ConvertStruct(*req.AddEntReq)
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugKytWesSaveentRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	if req.LicPictureByte != nil {
		fileMap["lic_picture_byte"] = *req.LicPictureByte
	}
	return fileMap
}
