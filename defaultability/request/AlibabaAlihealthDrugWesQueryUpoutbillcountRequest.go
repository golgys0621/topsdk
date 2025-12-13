package request


type AlibabaAlihealthDrugWesQueryUpoutbillcountRequest struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.license.token.get     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        单据类型：不填写查询全部。201:销售出库,202:退货出库,211:召回出库,212:赠品出库,214:盘亏出库,215:损坏出库,216:报废出库,217:其他出库,218:临床-退货出库,219:临床-销毁出库,220:报损出库,221:冻存出库至清洗,222:委托调拨出库指令,223:委托销售出库指令,237:直调退货     320:内部使用,321:使用出库,322:疫苗接种,323:使用出库,326:B2C出库,327:临床-领药出库     */
    BillType  *int64 `json:"bill_type,omitempty" required:"false" `
    /*
        单据上传开始日期     */
    BeginDate  *string `json:"begin_date" required:"true" `
    /*
        单据上传结束日期     */
    EndDate  *string `json:"end_date" required:"true" `
    /*
        委托单位     */
    AssRefEntId  *string `json:"ass_ref_ent_id,omitempty" required:"false" `
}

func (s *AlibabaAlihealthDrugWesQueryUpoutbillcountRequest) SetRefEntId(v string) *AlibabaAlihealthDrugWesQueryUpoutbillcountRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryUpoutbillcountRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugWesQueryUpoutbillcountRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryUpoutbillcountRequest) SetBillType(v int64) *AlibabaAlihealthDrugWesQueryUpoutbillcountRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryUpoutbillcountRequest) SetBeginDate(v string) *AlibabaAlihealthDrugWesQueryUpoutbillcountRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryUpoutbillcountRequest) SetEndDate(v string) *AlibabaAlihealthDrugWesQueryUpoutbillcountRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryUpoutbillcountRequest) SetAssRefEntId(v string) *AlibabaAlihealthDrugWesQueryUpoutbillcountRequest {
    s.AssRefEntId = &v
    return s
}

func (req *AlibabaAlihealthDrugWesQueryUpoutbillcountRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefEntId != nil) {
        paramMap["ref_ent_id"] = *req.RefEntId
    }
    if(req.LicenseToken != nil) {
        paramMap["license_token"] = *req.LicenseToken
    }
    if(req.BillType != nil) {
        paramMap["bill_type"] = *req.BillType
    }
    if(req.BeginDate != nil) {
        paramMap["begin_date"] = *req.BeginDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.AssRefEntId != nil) {
        paramMap["ass_ref_ent_id"] = *req.AssRefEntId
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugWesQueryUpoutbillcountRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}