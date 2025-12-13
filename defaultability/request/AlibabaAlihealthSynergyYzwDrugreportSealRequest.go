package request


type AlibabaAlihealthSynergyYzwDrugreportSealRequest struct {
    /*
        调用企业id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        用户id     */
    UserId  *string `json:"user_id" required:"true" `
    /*
        药检报告id     */
    ReportId  *int64 `json:"report_id" required:"true" `
    /*
        印章名称     */
    SealName  *string `json:"seal_name" required:"true" `
    /*
        X坐标     */
    PositionX  *int64 `json:"position_x,omitempty" required:"false" `
    /*
        Y坐标     */
    PositionY  *int64 `json:"position_y,omitempty" required:"false" `
}

func (s *AlibabaAlihealthSynergyYzwDrugreportSealRequest) SetRefEntId(v string) *AlibabaAlihealthSynergyYzwDrugreportSealRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportSealRequest) SetUserId(v string) *AlibabaAlihealthSynergyYzwDrugreportSealRequest {
    s.UserId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportSealRequest) SetReportId(v int64) *AlibabaAlihealthSynergyYzwDrugreportSealRequest {
    s.ReportId = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportSealRequest) SetSealName(v string) *AlibabaAlihealthSynergyYzwDrugreportSealRequest {
    s.SealName = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportSealRequest) SetPositionX(v int64) *AlibabaAlihealthSynergyYzwDrugreportSealRequest {
    s.PositionX = &v
    return s
}
func (s *AlibabaAlihealthSynergyYzwDrugreportSealRequest) SetPositionY(v int64) *AlibabaAlihealthSynergyYzwDrugreportSealRequest {
    s.PositionY = &v
    return s
}

func (req *AlibabaAlihealthSynergyYzwDrugreportSealRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.UserId != nil) {
        paramMap["user_id"] = *req.UserId
    }
    if(req.ReportId != nil) {
        paramMap["report_id"] = *req.ReportId
    }
    if(req.SealName != nil) {
        paramMap["seal_name"] = *req.SealName
    }
    if(req.PositionX != nil) {
        paramMap["position_x"] = *req.PositionX
    }
    if(req.PositionY != nil) {
        paramMap["position_y"] = *req.PositionY
    }
    return paramMap
}

func (req *AlibabaAlihealthSynergyYzwDrugreportSealRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}