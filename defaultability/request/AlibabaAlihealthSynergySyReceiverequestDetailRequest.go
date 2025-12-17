package request


type AlibabaAlihealthSynergySyReceiverequestDetailRequest struct {
    /*
        本企业refEntid     */
    RefEntId  *string `json:"ref_ent_id,omitempty" required:"false" `
    /*
        收到的索取id     */
    ReceiveRequestId  *int64 `json:"receive_request_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergySyReceiverequestDetailRequest) SetRefEntId(v string) *AlibabaAlihealthSynergySyReceiverequestDetailRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergySyReceiverequestDetailRequest) SetReceiveRequestId(v int64) *AlibabaAlihealthSynergySyReceiverequestDetailRequest {
    s.ReceiveRequestId = &v
    return s
}

func (req *AlibabaAlihealthSynergySyReceiverequestDetailRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.ReceiveRequestId != nil) {
        paramMap["receive_request_id"] = *req.ReceiveRequestId
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergySyReceiverequestDetailRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}