package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugSyGetentinfolistRequest struct {
	/*
	   本企业refEntId     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   查询企业信息参数     */
	TopEntInfoReqDto *domain.AlibabaAlihealthDrugSyGetentinfolistTopEntInfoReqDto `json:"top_ent_info_req_dto" required:"true" `
}

func (s *AlibabaAlihealthDrugSyGetentinfolistRequest) SetRefEntId(v string) *AlibabaAlihealthDrugSyGetentinfolistRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugSyGetentinfolistRequest) SetTopEntInfoReqDto(v domain.AlibabaAlihealthDrugSyGetentinfolistTopEntInfoReqDto) *AlibabaAlihealthDrugSyGetentinfolistRequest {
	s.TopEntInfoReqDto = &v
	return s
}

func (req *AlibabaAlihealthDrugSyGetentinfolistRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.TopEntInfoReqDto != nil {
		paramMap["top_ent_info_req_dto"] = util.ConvertStruct(*req.TopEntInfoReqDto)
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugSyGetentinfolistRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
