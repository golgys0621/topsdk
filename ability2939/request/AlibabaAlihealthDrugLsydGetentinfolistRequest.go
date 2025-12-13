package request

import (
	"github.com/golgys0621/topsdk/ability2939/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugLsydGetentinfolistRequest struct {
	/*
	   refEntId     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   查询企业信息参数     */
	QueryParam *domain.AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoReqDto `json:"query_param,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugLsydGetentinfolistRequest) SetRefEntId(v string) *AlibabaAlihealthDrugLsydGetentinfolistRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugLsydGetentinfolistRequest) SetQueryParam(v domain.AlibabaAlihealthDrugLsydGetentinfolistTopEntInfoReqDto) *AlibabaAlihealthDrugLsydGetentinfolistRequest {
	s.QueryParam = &v
	return s
}

func (req *AlibabaAlihealthDrugLsydGetentinfolistRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.QueryParam != nil {
		paramMap["query_param"] = util.ConvertStruct(*req.QueryParam)
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugLsydGetentinfolistRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
