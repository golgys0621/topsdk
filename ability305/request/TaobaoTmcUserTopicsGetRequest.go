package request


type TaobaoTmcUserTopicsGetRequest struct {
    /*
        卖家nick     */
    Nick  *string `json:"nick,omitempty" required:"false" `
}

func (s *TaobaoTmcUserTopicsGetRequest) SetNick(v string) *TaobaoTmcUserTopicsGetRequest {
    s.Nick = &v
    return s
}

func (req *TaobaoTmcUserTopicsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    return paramMap
}

func (req *TaobaoTmcUserTopicsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}