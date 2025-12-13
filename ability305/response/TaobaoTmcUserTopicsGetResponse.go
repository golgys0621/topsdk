package response

import (
)

type TaobaoTmcUserTopicsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        错误信息
    */
    ResultMessage  string `json:"result_message,omitempty" `
    /*
        topic列表
    */
    Topics  []string `json:"topics,omitempty" `
    /*
        错误码
    */
    ResultCode  string `json:"result_code,omitempty" `
}
