package response

import (
    "github.com/golgys0621/topsdk/ability2940/domain"
)

type AlibabaAlihealthDrugYljgSaveentResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        往来单位新增接口返回
    */
    Result  domain.AlibabaAlihealthDrugYljgSaveentResultModel `json:"result,omitempty" `
}
