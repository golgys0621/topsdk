package request


type AlibabaAlihealthDrugKytWesRemnantbillUploadRequest struct {
    /*
        企业ref_ent_id     */
    RefEntId  *string `json:"ref_ent_id" required:"true" `
    /*
        获取licenseToken，通过alibaba.alihealth.drug.code.kyt.wes.getlicense     */
    LicenseToken  *string `json:"license_token" required:"true" `
    /*
        零头入库：106；零头出库：210     */
    BillType  *string `json:"bill_type" required:"true" `
    /*
        单据编号     */
    BillCode  *string `json:"bill_code" required:"true" `
    /*
        单据时间:yyyy-MM-dd HH:mm:ss     */
    BillTime  *string `json:"bill_time" required:"true" `
    /*
        发货企业【注意：该入参是ref_ent_id,并不是ent_id】     */
    FromRefUserId  *string `json:"from_ref_user_id" required:"true" `
    /*
        收货企业【注意：该入参是ref_ent_id,并不是ent_id】     */
    ToRefUserId  *string `json:"to_ref_user_id" required:"true" `
    /*
        委托企业【注意：该入参是ref_ent_id,并不是ent_id】     */
    AssRefEntId  *string `json:"ass_ref_ent_id,omitempty" required:"false" `
    /*
        配送企业【注意：该入参是ref_ent_id,并不是ent_id】     */
    DisRefEntId  *string `json:"dis_ref_ent_id,omitempty" required:"false" `
    /*
        药品ID     */
    DrugEntBaseInfoId  *string `json:"drug_ent_base_info_id" required:"true" `
    /*
        生产日期：yyyy-MM-dd      */
    ProduceDate  *string `json:"produce_date" required:"true" `
    /*
        有效期:yyyyMMdd      */
    ExpireDate  *string `json:"expire_date" required:"true" `
    /*
        生产批次     */
    ProduceBatchNo  *string `json:"produce_batch_no" required:"true" `
    /*
        药品制剂数量     */
    InputAmount  *string `json:"input_amount" required:"true" `
}

func (s *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) SetRefEntId(v string) *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest {
    s.RefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) SetLicenseToken(v string) *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest {
    s.LicenseToken = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) SetBillType(v string) *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest {
    s.BillType = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) SetBillCode(v string) *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest {
    s.BillCode = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) SetBillTime(v string) *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest {
    s.BillTime = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) SetFromRefUserId(v string) *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest {
    s.FromRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) SetToRefUserId(v string) *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest {
    s.ToRefUserId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) SetAssRefEntId(v string) *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest {
    s.AssRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) SetDisRefEntId(v string) *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest {
    s.DisRefEntId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) SetDrugEntBaseInfoId(v string) *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest {
    s.DrugEntBaseInfoId = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) SetProduceDate(v string) *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest {
    s.ProduceDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) SetExpireDate(v string) *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest {
    s.ExpireDate = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) SetProduceBatchNo(v string) *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest {
    s.ProduceBatchNo = &v
    return s
}
func (s *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) SetInputAmount(v string) *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest {
    s.InputAmount = &v
    return s
}

func (req *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) ToMap() map[string]interface{} {
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
    if(req.BillCode != nil) {
        paramMap["bill_code"] = *req.BillCode
    }
    if(req.BillTime != nil) {
        paramMap["bill_time"] = *req.BillTime
    }
    if(req.FromRefUserId != nil) {
        paramMap["from_ref_user_id"] = *req.FromRefUserId
    }
    if(req.ToRefUserId != nil) {
        paramMap["to_ref_user_id"] = *req.ToRefUserId
    }
    if(req.AssRefEntId != nil) {
        paramMap["ass_ref_ent_id"] = *req.AssRefEntId
    }
    if(req.DisRefEntId != nil) {
        paramMap["dis_ref_ent_id"] = *req.DisRefEntId
    }
    if(req.DrugEntBaseInfoId != nil) {
        paramMap["drug_ent_base_info_id"] = *req.DrugEntBaseInfoId
    }
    if(req.ProduceDate != nil) {
        paramMap["produce_date"] = *req.ProduceDate
    }
    if(req.ExpireDate != nil) {
        paramMap["expire_date"] = *req.ExpireDate
    }
    if(req.ProduceBatchNo != nil) {
        paramMap["produce_batch_no"] = *req.ProduceBatchNo
    }
    if(req.InputAmount != nil) {
        paramMap["input_amount"] = *req.InputAmount
    }
    return paramMap
}

func (req *AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}