package response

import (
)

type AlibabaAlihealthDrugMyjCodewarnCreatecodewarntaskResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回任务批次编码
    */
    Model  string `json:"model,omitempty" `
    /*
        返回结果
    */
    MsgCode  string `json:"msg_code,omitempty" `
    /*
        返回结果
    */
    MsgInfo  string `json:"msg_info,omitempty" `
    /*
        是否成功(true 成功 ,false失败)
    */
    ResponseSuccess  bool `json:"response_success,omitempty" `
}
