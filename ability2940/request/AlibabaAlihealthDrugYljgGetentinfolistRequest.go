package request

import (
	"github.com/golgys0621/topsdk/ability2940/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugYljgGetentinfolistRequest struct {
	/*
	   refEntId     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   查询企业信息参数     */
	QueryParam *domain.AlibabaAlihealthDrugYljgGetentinfolistTopEntInfoReqDto `json:"query_param,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugYljgGetentinfolistRequest) SetRefEntId(v string) *AlibabaAlihealthDrugYljgGetentinfolistRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugYljgGetentinfolistRequest) SetQueryParam(v domain.AlibabaAlihealthDrugYljgGetentinfolistTopEntInfoReqDto) *AlibabaAlihealthDrugYljgGetentinfolistRequest {
	s.QueryParam = &v
	return s
}

func (req *AlibabaAlihealthDrugYljgGetentinfolistRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.QueryParam != nil {
		paramMap["query_param"] = util.ConvertStruct(*req.QueryParam)
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugYljgGetentinfolistRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
