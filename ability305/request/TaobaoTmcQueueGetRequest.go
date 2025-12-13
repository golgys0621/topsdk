package request


type TaobaoTmcQueueGetRequest struct {
    /*
        TMC组名     */
    GroupName  *string `json:"group_name" required:"true" `
}

func (s *TaobaoTmcQueueGetRequest) SetGroupName(v string) *TaobaoTmcQueueGetRequest {
    s.GroupName = &v
    return s
}

func (req *TaobaoTmcQueueGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GroupName != nil) {
        paramMap["group_name"] = *req.GroupName
    }
    return paramMap
}

func (req *TaobaoTmcQueueGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}