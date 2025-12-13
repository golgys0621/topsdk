package response

import (
)

type AlibabaAlihealthDrugKytWesBillpartialrejectResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        是否响应成功
    */
    ResponseSuccess  bool `json:"response_success,omitempty" `
    /*
        返回值
    */
    MsgInfo  string `json:"msg_info,omitempty" `
    /*
        返回码
    */
    MsgCode  string `json:"msg_code,omitempty" `
}
