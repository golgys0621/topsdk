package request

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthSynergySySignatureCreateRequest struct {
	/*
	   资料接收企业refEntId串，多个用英文逗号拼接     */
	ReceiveRefEntIds *string `json:"receive_ref_ent_ids" required:"true" `
	/*
	   给收件企业的留言     */
	Note *string `json:"note,omitempty" required:"false" `
	/*
	   产品资料id集合，英文逗号拼接,当签章类型为产品资料时该字段生效     */
	ProductResources *[]string `json:"product_resources,omitempty" required:"false" `
	/*
	   委托人资料id集合，英文逗号拼接,当签章类型为企业资料时该字段生效     */
	AgentPersonResourcesIds *[]string `json:"agent_person_resources_ids,omitempty" required:"false" `
	/*
	   签章创建人     */
	SubmitPerson *string `json:"submit_person,omitempty" required:"false" `
	/*
	   发送企业id     */
	RefEntId *string `json:"ref_ent_id" required:"true" `
	/*
	   企业资料id集合，英文逗号拼接,当签章类型为企业资料时该字段生效     */
	EntResources *[]string `json:"ent_resources,omitempty" required:"false" `
	/*
	   签章类型0:企业资料(支持委托人资料)1:产品资料(不支持委托人资料)     */
	Type *string `json:"type" required:"true" `
	/*
	   资料接收企业对应的索取id，多个逗号拼接     */
	SyRequestDetailIds *string `json:"sy_request_detail_ids,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergySySignatureCreateRequest) SetReceiveRefEntIds(v string) *AlibabaAlihealthSynergySySignatureCreateRequest {
	s.ReceiveRefEntIds = &v
	return s
}
func (s *AlibabaAlihealthSynergySySignatureCreateRequest) SetNote(v string) *AlibabaAlihealthSynergySySignatureCreateRequest {
	s.Note = &v
	return s
}
func (s *AlibabaAlihealthSynergySySignatureCreateRequest) SetProductResources(v []string) *AlibabaAlihealthSynergySySignatureCreateRequest {
	s.ProductResources = &v
	return s
}
func (s *AlibabaAlihealthSynergySySignatureCreateRequest) SetAgentPersonResourcesIds(v []string) *AlibabaAlihealthSynergySySignatureCreateRequest {
	s.AgentPersonResourcesIds = &v
	return s
}
func (s *AlibabaAlihealthSynergySySignatureCreateRequest) SetSubmitPerson(v string) *AlibabaAlihealthSynergySySignatureCreateRequest {
	s.SubmitPerson = &v
	return s
}
func (s *AlibabaAlihealthSynergySySignatureCreateRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySySignatureCreateRequest {
	s.RefEntId = &v
	return s
}
func (s *AlibabaAlihealthSynergySySignatureCreateRequest) SetEntResources(v []string) *AlibabaAlihealthSynergySySignatureCreateRequest {
	s.EntResources = &v
	return s
}
func (s *AlibabaAlihealthSynergySySignatureCreateRequest) SetType(v string) *AlibabaAlihealthSynergySySignatureCreateRequest {
	s.Type = &v
	return s
}
func (s *AlibabaAlihealthSynergySySignatureCreateRequest) SetSyRequestDetailIds(v string) *AlibabaAlihealthSynergySySignatureCreateRequest {
	s.SyRequestDetailIds = &v
	return s
}

func (req *AlibabaAlihealthSynergySySignatureCreateRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.ReceiveRefEntIds != nil {
		paramMap["receive_ref_ent_ids"] = *req.ReceiveRefEntIds
	}
	if req.Note != nil {
		paramMap["note"] = *req.Note
	}
	if req.ProductResources != nil {
		paramMap["product_resources"] = util.ConvertBasicList(*req.ProductResources)
	}
	if req.AgentPersonResourcesIds != nil {
		paramMap["agent_person_resources_ids"] = util.ConvertBasicList(*req.AgentPersonResourcesIds)
	}
	if req.SubmitPerson != nil {
		paramMap["submit_person"] = *req.SubmitPerson
	}
	if req.RefEntId != nil {
		paramMap["ref_ent_id"] = *req.RefEntId
	}
	if req.EntResources != nil {
		paramMap["ent_resources"] = util.ConvertBasicList(*req.EntResources)
	}
	if req.Type != nil {
		paramMap["type"] = *req.Type
	}
	if req.SyRequestDetailIds != nil {
		paramMap["sy_request_detail_ids"] = *req.SyRequestDetailIds
	}
	return paramMap
}

func (req *AlibabaAlihealthSynergySySignatureCreateRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
