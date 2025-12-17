package request

import (
	"github.com/golgys0621/topsdk/defaultability/domain"
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySyInboxCheckRequest struct {
	/*
	   企业资料查收列表     */
	EntResourceCheckInfos *[]domain.AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoDetailTopRequest `json:"ent_resource_check_infos,omitempty" required:"false" `
	/*
	   0拒收，1全部查收，2部分查收     */
	CheckStatus *int64 `json:"check_status" required:"true" `
	/*
	   拒绝原因，拒收和部分查收时必传     */
	RejectReason *string `json:"reject_reason,omitempty" required:"false" `
	/*
	   委托人资料查收列表     */
	AgentCheckInfos *[]domain.AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoTopRequest `json:"agent_check_infos,omitempty" required:"false" `
	/*
	   产品资料查收列表     */
	ProductCheckInfos *[]domain.AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoTopRequest `json:"product_check_infos,omitempty" required:"false" `
	/*
	   本企业refEntId     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   收件箱id     */
	InboxId *int64 `json:"inbox_id" required:"true" `
}

func (s *AlibabaAlihealthSynergySyInboxCheckRequest) SetEntResourceCheckInfos(v []domain.AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoDetailTopRequest) *AlibabaAlihealthSynergySyInboxCheckRequest {
	s.EntResourceCheckInfos = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxCheckRequest) SetCheckStatus(v int64) *AlibabaAlihealthSynergySyInboxCheckRequest {
	s.CheckStatus = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxCheckRequest) SetRejectReason(v string) *AlibabaAlihealthSynergySyInboxCheckRequest {
	s.RejectReason = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxCheckRequest) SetAgentCheckInfos(v []domain.AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoTopRequest) *AlibabaAlihealthSynergySyInboxCheckRequest {
	s.AgentCheckInfos = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxCheckRequest) SetProductCheckInfos(v []domain.AlibabaAlihealthSynergySyInboxCheckSyInboxCheckInfoTopRequest) *AlibabaAlihealthSynergySyInboxCheckRequest {
	s.ProductCheckInfos = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxCheckRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyInboxCheckRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySyInboxCheckRequest) SetInboxId(v int64) *AlibabaAlihealthSynergySyInboxCheckRequest {
	s.InboxId = &v
	return s
}

func (req *AlibabaAlihealthSynergySyInboxCheckRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.EntResourceCheckInfos != nil {
		paramMap["ent_resource_check_infos"] = util.ConvertStructList(*req.EntResourceCheckInfos)
	}
	if req.CheckStatus != nil {
		paramMap["check_status"] = *req.CheckStatus
	}
	if req.RejectReason != nil {
		paramMap["reject_reason"] = *req.RejectReason
	}
	if req.AgentCheckInfos != nil {
		paramMap["agent_check_infos"] = util.ConvertStructList(*req.AgentCheckInfos)
	}
	if req.ProductCheckInfos != nil {
		paramMap["product_check_infos"] = util.ConvertStructList(*req.ProductCheckInfos)
	}
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.InboxId != nil {
		paramMap["inbox_id"] = *req.InboxId
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergySyInboxCheckRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
