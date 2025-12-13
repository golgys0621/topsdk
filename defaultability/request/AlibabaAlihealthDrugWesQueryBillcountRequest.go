package request


type AlibabaAlihealthDrugWesQueryBillcountRequest struct {
    /*
        企业ID     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.license.token.get     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        单据类型：不填写查询全部。101:生产入库,102:采购入库,103:退货入库,104:调拨入库,107:供应入库,108:召回入库,109:打标入库,110:赠品入库,111:盘盈入库,112:报废入库,113:其他入库,114:临床-采购入库,115:临床-退货入库,116:消费者退货入库,117:临床-替换入库,118:报溢入库,119:新鲜制剂入库,120:委托退货入库指令     201:销售出库,202:退货出库,211:召回出库,212:赠品出库,214:盘亏出库,215:损坏出库,216:报废出库,217:其他出库,218:临床-退货出库,219:临床-销毁出库,220:报损出库,221:冻存出库至清洗,222:委托调拨出库指令,223:委托销售出库指令,237:直调退货     320:内部使用,321:使用出库,322:疫苗接种,323:使用出库,326:B2C出库,327:临床-领药出库     */
    BillType  *int64 `json:"bill_type,omitempty" required:"false" `
    /*
        单据上传开始日期     */
    BeginDate  *string `json:"begin_date" required:"true" `
    /*
        单据上传结束日期     */
    EndDate  *string `json:"end_date" required:"true" `
}

func (s *AlibabaAlihealthDrugWesQueryBillcountRequest) SetRefEntId(v string) *AlibabaAlihealthDrugWesQueryBillcountRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryBillcountRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugWesQueryBillcountRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryBillcountRequest) SetBillType(v int64) *AlibabaAlihealthDrugWesQueryBillcountRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryBillcountRequest) SetBeginDate(v string) *AlibabaAlihealthDrugWesQueryBillcountRequest {
    s.BeginDate = &v
    return s
}
func (s *AlibabaAlihealthDrugWesQueryBillcountRequest) SetEndDate(v string) *AlibabaAlihealthDrugWesQueryBillcountRequest {
    s.EndDate = &v
    return s
}

func (req *AlibabaAlihealthDrugWesQueryBillcountRequest) ToMap() map[string]interface{} {
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
    return paramMap
}

func (req *AlibabaAlihealthDrugWesQueryBillcountRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}