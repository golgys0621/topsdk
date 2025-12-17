package request

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyResourceResultRequest struct {
	/*
	   任务id,多个任务id使用,拼接     */
	Ids *[]int64 `json:"ids" required:"true" `
	/*
	   企业refEntId     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
}

func (s *AlibabaAlihealthSynergySyResourceResultRequest) SetIds(v []int64) *AlibabaAlihealthSynergySyResourceResultRequest {
	s.Ids = &v
	return s
}
func (s *AlibabaAlihealthSynergySyResourceResultRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyResourceResultRequest {
	s.RefEntId = &v
	return s
}

func (req *AlibabaAlihealthSynergySyResourceResultRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Ids != nil {
		paramMap["ids"] = util.ConvertBasicList(*req.Ids)
	}
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergySyResourceResultRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
