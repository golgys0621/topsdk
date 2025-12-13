package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugMscGetentinfolistRequest struct {
	/*
	   refEntId     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   查询企业信息参数     */
	QueryParam *domain.AlibabaAlihealthDrugMscGetentinfolistTopEntInfoReqDto `json:"query_param,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugMscGetentinfolistRequest) SetRefEntId(v string) *AlibabaAlihealthDrugMscGetentinfolistRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugMscGetentinfolistRequest) SetQueryParam(v domain.AlibabaAlihealthDrugMscGetentinfolistTopEntInfoReqDto) *AlibabaAlihealthDrugMscGetentinfolistRequest {
	s.QueryParam = &v
	return s
}

func (req *AlibabaAlihealthDrugMscGetentinfolistRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.QueryParam != nil {
		paramMap["query_param"] = util.ConvertStruct(*req.QueryParam)
	}
	return paramMap
}

func (req *AlibabaAlihealthDrugMscGetentinfolistRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
