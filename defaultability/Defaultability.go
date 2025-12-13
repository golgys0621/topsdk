package defaultability

import (
	"errors"
	"log"

	"github.com/golgys0621/topsdk"
	"github.com/golgys0621/topsdk/defaultability/request"
	"github.com/golgys0621/topsdk/defaultability/response"
	"github.com/golgys0621/topsdk/util"
)

type Defaultability struct {
	Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability {
	return &Defaultability{client}
}

/*
查询上传报告信息接口
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwQuerysealdrugreport(req *request.AlibabaAlihealthSynergyYzwQuerysealdrugreportRequest) (*response.AlibabaAlihealthSynergyYzwQuerysealdrugreportResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.querysealdrugreport", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwQuerysealdrugreportResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwQuerysealdrugreport error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询批量上传药检报告ocr状态
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatus(req *request.AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusRequest) (*response.AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.getdrugreportwithocrstatus", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwGetdrugreportwithocrstatusResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwGetdrugreportwithocrstatus error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
通过文件保存一张网药检报告，支持委托企业代授权企业上传
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwDrugreportByfileSave(req *request.AlibabaAlihealthSynergyYzwDrugreportByfileSaveRequest) (*response.AlibabaAlihealthSynergyYzwDrugreportByfileSaveResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.drugreport.byfile.save", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwDrugreportByfileSaveResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwDrugreportByfileSave error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
药检报告的签章接口
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwDrugreportSeal(req *request.AlibabaAlihealthSynergyYzwDrugreportSealRequest) (*response.AlibabaAlihealthSynergyYzwDrugreportSealResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.drugreport.seal", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwDrugreportSealResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwDrugreportSeal error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
码上传权限包有效期查询
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscServiceinfo(req *request.AlibabaAlihealthDrugMscServiceinfoRequest) (*response.AlibabaAlihealthDrugMscServiceinfoResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.serviceinfo", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscServiceinfoResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscServiceinfo error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
对药检报告盖章，支持委托企业代授权企业盖章
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwDrugreportAssSeal(req *request.AlibabaAlihealthSynergyYzwDrugreportAssSealRequest) (*response.AlibabaAlihealthSynergyYzwDrugreportAssSealResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.drugreport.ass.seal", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwDrugreportAssSealResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwDrugreportAssSeal error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询企业是否上传过药检报告
*/
func (ability *Defaultability) AlibabaAlihealthEntYzwHasuploaddrugreport(req *request.AlibabaAlihealthEntYzwHasuploaddrugreportRequest) (*response.AlibabaAlihealthEntYzwHasuploaddrugreportResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.ent.yzw.hasuploaddrugreport", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthEntYzwHasuploaddrugreportResponse{}
	if err != nil {
		log.Println("alibabaAlihealthEntYzwHasuploaddrugreport error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查看通过http方式保存药检报告状态
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatus(req *request.AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusRequest) (*response.AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.getdrugreportwithhttpstatus", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwGetdrugreportwithhttpstatusResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwGetdrugreportwithhttpstatus error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
一张网报告操作日志
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwDrugreportOptHistoryAll(req *request.AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllRequest) (*response.AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.drugreport.opt.history.all", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwDrugreportOptHistoryAllResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwDrugreportOptHistoryAll error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
通过http方式上传药检报告
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwSavedrugreportbyhttp(req *request.AlibabaAlihealthSynergyYzwSavedrugreportbyhttpRequest) (*response.AlibabaAlihealthSynergyYzwSavedrugreportbyhttpResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.savedrugreportbyhttp", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwSavedrugreportbyhttpResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwSavedrugreportbyhttp error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询上游企业的待签收药检报告信息
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwQuerydrugreport(req *request.AlibabaAlihealthSynergyYzwQuerydrugreportRequest) (*response.AlibabaAlihealthSynergyYzwQuerydrugreportResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.querydrugreport", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwQuerydrugreportResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwQuerydrugreport error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询委托上传的报告
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwQueryassdrugreportinfo(req *request.AlibabaAlihealthSynergyYzwQueryassdrugreportinfoRequest) (*response.AlibabaAlihealthSynergyYzwQueryassdrugreportinfoResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.queryassdrugreportinfo", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwQueryassdrugreportinfoResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwQueryassdrugreportinfo error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
由于无法返回多条匹配信息，本接口已不再推荐使用，推荐使用新接口，新接口名称：alibaba.alihealth.drug.msc.getentinfolist
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscGetentinfonew(req *request.AlibabaAlihealthDrugMscGetentinfonewRequest) (*response.AlibabaAlihealthDrugMscGetentinfonewResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.getentinfonew", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscGetentinfonewResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscGetentinfonew error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询入库单据详情以及码
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscBillinDetailwithcode(req *request.AlibabaAlihealthDrugMscBillinDetailwithcodeRequest) (*response.AlibabaAlihealthDrugMscBillinDetailwithcodeResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.billin.detailwithcode", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscBillinDetailwithcodeResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscBillinDetailwithcode error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
批发企业签收并盖章和拒收药检报告
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwDrugreportOpt(req *request.AlibabaAlihealthSynergyYzwDrugreportOptRequest) (*response.AlibabaAlihealthSynergyYzwDrugreportOptResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.drugreport.opt", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwDrugreportOptResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwDrugreportOpt error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询药检报告签章信息
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwDrugreportGet(req *request.AlibabaAlihealthSynergyYzwDrugreportGetRequest) (*response.AlibabaAlihealthSynergyYzwDrugreportGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.drugreport.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwDrugreportGetResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwDrugreportGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
企业信息-根据企业名称或证号查询企业信息列表
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscGetentinfolist(req *request.AlibabaAlihealthDrugMscGetentinfolistRequest) (*response.AlibabaAlihealthDrugMscGetentinfolistResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.getentinfolist", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscGetentinfolistResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscGetentinfolist error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询一段时间内已经签收的单据
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwBillSignedQuery(req *request.AlibabaAlihealthSynergyYzwBillSignedQueryRequest) (*response.AlibabaAlihealthSynergyYzwBillSignedQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.bill.signed.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwBillSignedQueryResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwBillSignedQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询企业质检章
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwListseals(req *request.AlibabaAlihealthSynergyYzwListsealsRequest) (*response.AlibabaAlihealthSynergyYzwListsealsResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.listseals", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwListsealsResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwListseals error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
异常码详情列表
*/
func (ability *Defaultability) AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillist(req *request.AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistRequest) (*response.AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.myj.codewarn.codewarninginfodetaillist", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillistResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMyjCodewarnCodewarninginfodetaillist error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
码预警列表
*/
func (ability *Defaultability) AlibabaAlihealthDrugMyjCodewarnCodewarninglist(req *request.AlibabaAlihealthDrugMyjCodewarnCodewarninglistRequest) (*response.AlibabaAlihealthDrugMyjCodewarnCodewarninglistResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.myj.codewarn.codewarninglist", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMyjCodewarnCodewarninglistResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMyjCodewarnCodewarninglist error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询待签收企业
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwQueryunsignents(req *request.AlibabaAlihealthSynergyYzwQueryunsignentsRequest) (*response.AlibabaAlihealthSynergyYzwQueryunsignentsResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.queryunsignents", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwQueryunsignentsResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwQueryunsignents error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
生成码预警任务
*/
func (ability *Defaultability) AlibabaAlihealthDrugMyjCodewarnCreatecodewarntask(req *request.AlibabaAlihealthDrugMyjCodewarnCreatecodewarntaskRequest) (*response.AlibabaAlihealthDrugMyjCodewarnCreatecodewarntaskResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.myj.codewarn.createcodewarntask", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMyjCodewarnCreatecodewarntaskResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMyjCodewarnCreatecodewarntask error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
保存药检报告
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwSavedrugreport(req *request.AlibabaAlihealthSynergyYzwSavedrugreportRequest) (*response.AlibabaAlihealthSynergyYzwSavedrugreportResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.savedrugreport", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwSavedrugreportResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwSavedrugreport error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
企业搜索自己授权的物流企业
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscListauths(req *request.AlibabaAlihealthDrugMscListauthsRequest) (*response.AlibabaAlihealthDrugMscListauthsResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.listauths", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscListauthsResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscListauths error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据报告id或药品批次或药品和报告创建时间区间查询药检报告
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwQuerydrugbatchreport(req *request.AlibabaAlihealthSynergyYzwQuerydrugbatchreportRequest) (*response.AlibabaAlihealthSynergyYzwQuerydrugbatchreportResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.querydrugbatchreport", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwQuerydrugbatchreportResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwQuerydrugbatchreport error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
药检报告删除
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwDeletedrugreport(req *request.AlibabaAlihealthSynergyYzwDeletedrugreportRequest) (*response.AlibabaAlihealthSynergyYzwDeletedrugreportResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.deletedrugreport", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwDeletedrugreportResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwDeletedrugreport error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地分批上传文件
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwSavedatabucket(req *request.AlibabaAlihealthSynergyYzwSavedatabucketRequest) (*response.AlibabaAlihealthSynergyYzwSavedatabucketResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.savedatabucket", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwSavedatabucketResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwSavedatabucket error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询上游出库单数量
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscQueryUpoutbillcount(req *request.AlibabaAlihealthDrugMscQueryUpoutbillcountRequest) (*response.AlibabaAlihealthDrugMscQueryUpoutbillcountResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.query.upoutbillcount", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscQueryUpoutbillcountResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscQueryUpoutbillcount error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
同步获取码预警信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugMyjCodewarnGetwarninfo(req *request.AlibabaAlihealthDrugMyjCodewarnGetwarninfoRequest) (*response.AlibabaAlihealthDrugMyjCodewarnGetwarninfoResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.myj.codewarn.getwarninfo", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMyjCodewarnGetwarninfoResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMyjCodewarnGetwarninfo error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询本企业出入库单数量
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscQueryBillcount(req *request.AlibabaAlihealthDrugMscQueryBillcountRequest) (*response.AlibabaAlihealthDrugMscQueryBillcountResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.query.billcount", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscQueryBillcountResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscQueryBillcount error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询企业信息
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwGetentinfolist(req *request.AlibabaAlihealthSynergyYzwGetentinfolistRequest) (*response.AlibabaAlihealthSynergyYzwGetentinfolistResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.getentinfolist", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwGetentinfolistResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwGetentinfolist error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询授权企业的药品信息
*/
func (ability *Defaultability) AlibabaAlihealthEntYzwAuthGetdruginfo(req *request.AlibabaAlihealthEntYzwAuthGetdruginfoRequest) (*response.AlibabaAlihealthEntYzwAuthGetdruginfoResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.ent.yzw.auth.getdruginfo", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthEntYzwAuthGetdruginfoResponse{}
	if err != nil {
		log.Println("alibabaAlihealthEntYzwAuthGetdruginfo error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
企业创建我的药检报告，支持2个印章
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelf(req *request.AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfRequest) (*response.AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.drugreport.byfile.save.self", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwDrugreportByfileSaveSelfResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwDrugreportByfileSaveSelf error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
部分处理成功单据处理失败的码明细
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscListbillprocesspartsuccess(req *request.AlibabaAlihealthDrugMscListbillprocesspartsuccessRequest) (*response.AlibabaAlihealthDrugMscListbillprocesspartsuccessResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.listbillprocesspartsuccess", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscListbillprocesspartsuccessResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscListbillprocesspartsuccess error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
批发企业上传入出库单据接口
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscUploadcircubill(req *request.AlibabaAlihealthDrugMscUploadcircubillRequest) (*response.AlibabaAlihealthDrugMscUploadcircubillResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.uploadcircubill", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscUploadcircubillResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscUploadcircubill error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
通过时间段批量查询入出库单信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscSearchbill(req *request.AlibabaAlihealthDrugMscSearchbillRequest) (*response.AlibabaAlihealthDrugMscSearchbillResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.searchbill", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscSearchbillResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscSearchbill error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询单据详情
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscSearchbillDetail(req *request.AlibabaAlihealthDrugMscSearchbillDetailRequest) (*response.AlibabaAlihealthDrugMscSearchbillDetailResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.searchbill.detail", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscSearchbillDetailResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscSearchbillDetail error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
零头出入库单据上传
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscRemnantbillUpload(req *request.AlibabaAlihealthDrugMscRemnantbillUploadRequest) (*response.AlibabaAlihealthDrugMscRemnantbillUploadResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.remnantbill.upload", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscRemnantbillUploadResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscRemnantbillUpload error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
单据处理状态查询
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscBillSearchstatus(req *request.AlibabaAlihealthDrugMscBillSearchstatusRequest) (*response.AlibabaAlihealthDrugMscBillSearchstatusResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.bill.searchstatus", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscBillSearchstatusResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscBillSearchstatus error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
新增往来单位企业
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscSaveent(req *request.AlibabaAlihealthDrugMscSaveentRequest) (*response.AlibabaAlihealthDrugMscSaveentResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.saveent", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscSaveentResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscSaveent error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
由于无法返回多条匹配信息，本接口已不再推荐使用，推荐使用新接口，新接口名称：alibaba.alihealth.drug.msc.getentinfolist
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscGetentinfo(req *request.AlibabaAlihealthDrugMscGetentinfoRequest) (*response.AlibabaAlihealthDrugMscGetentinfoResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.getentinfo", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscGetentinfoResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscGetentinfo error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询上游出库单
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwQuerydrugreportass(req *request.AlibabaAlihealthSynergyYzwQuerydrugreportassRequest) (*response.AlibabaAlihealthSynergyYzwQuerydrugreportassResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.querydrugreportass", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwQuerydrugreportassResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwQuerydrugreportass error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询往来单位列表
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscListparts(req *request.AlibabaAlihealthDrugMscListpartsRequest) (*response.AlibabaAlihealthDrugMscListpartsResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.listparts", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscListpartsResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscListparts error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询企业质检章，支持委托企业
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwListsealass(req *request.AlibabaAlihealthSynergyYzwListsealassRequest) (*response.AlibabaAlihealthSynergyYzwListsealassResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.listsealass", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwListsealassResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwListsealass error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据企业唯一标识查看企业详细信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscGetbyrefentid(req *request.AlibabaAlihealthDrugMscGetbyrefentidRequest) (*response.AlibabaAlihealthDrugMscGetbyrefentidResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.getbyrefentid", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscGetbyrefentidResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscGetbyrefentid error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询已经签收但是还未盖章的单据
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwSignednotstampedbillQuery(req *request.AlibabaAlihealthSynergyYzwSignednotstampedbillQueryRequest) (*response.AlibabaAlihealthSynergyYzwSignednotstampedbillQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.signednotstampedbill.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwSignednotstampedbillQueryResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwSignednotstampedbillQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
物流企业查询货主企业信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscSynonymauths(req *request.AlibabaAlihealthDrugMscSynonymauthsRequest) (*response.AlibabaAlihealthDrugMscSynonymauthsResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.synonymauths", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscSynonymauthsResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscSynonymauths error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据企业主键查看企业详细信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscGetbyentid(req *request.AlibabaAlihealthDrugMscGetbyentidRequest) (*response.AlibabaAlihealthDrugMscGetbyentidResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.getbyentid", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscGetbyentidResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscGetbyentid error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
代授权单位操作药检报告
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwDrugreportOptByagent(req *request.AlibabaAlihealthSynergyYzwDrugreportOptByagentRequest) (*response.AlibabaAlihealthSynergyYzwDrugreportOptByagentResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.drugreport.opt.byagent", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwDrugreportOptByagentResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwDrugreportOptByagent error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
通过http方式保存药检报告,支持委托企业
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwSavedrugreportbyhttpass(req *request.AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassRequest) (*response.AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.savedrugreportbyhttpass", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwSavedrugreportbyhttpassResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwSavedrugreportbyhttpass error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
一张网查询药品信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugYzwDrugtable(req *request.AlibabaAlihealthDrugYzwDrugtableRequest) (*response.AlibabaAlihealthDrugYzwDrugtableResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.yzw.drugtable", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugYzwDrugtableResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugYzwDrugtable error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询药品id
*/
func (ability *Defaultability) AlibabaAlihealthEntYzwGetdrugidbyprodcode(req *request.AlibabaAlihealthEntYzwGetdrugidbyprodcodeRequest) (*response.AlibabaAlihealthEntYzwGetdrugidbyprodcodeResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.ent.yzw.getdrugidbyprodcode", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthEntYzwGetdrugidbyprodcodeResponse{}
	if err != nil {
		log.Println("alibabaAlihealthEntYzwGetdrugidbyprodcode error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
企业上传出入库信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscUploadinoutbill(req *request.AlibabaAlihealthDrugMscUploadinoutbillRequest) (*response.AlibabaAlihealthDrugMscUploadinoutbillResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.uploadinoutbill", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscUploadinoutbillResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscUploadinoutbill error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询出库单据详情以及码
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscBilloutDetailwithcodes(req *request.AlibabaAlihealthDrugMscBilloutDetailwithcodesRequest) (*response.AlibabaAlihealthDrugMscBilloutDetailwithcodesResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.billout.detailwithcodes", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscBilloutDetailwithcodesResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscBilloutDetailwithcodes error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据码查询上游药检报告
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwQueryreportbycode(req *request.AlibabaAlihealthSynergyYzwQueryreportbycodeRequest) (*response.AlibabaAlihealthSynergyYzwQueryreportbycodeResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.queryreportbycode", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwQueryreportbycodeResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwQueryreportbycode error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询药品目录信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugMscDrugtable(req *request.AlibabaAlihealthDrugMscDrugtableRequest) (*response.AlibabaAlihealthDrugMscDrugtableResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.msc.drugtable", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugMscDrugtableResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugMscDrugtable error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
批量上传药检报告OCR
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwSavedrugreportwithocr(req *request.AlibabaAlihealthSynergyYzwSavedrugreportwithocrRequest) (*response.AlibabaAlihealthSynergyYzwSavedrugreportwithocrResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.savedrugreportwithocr", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwSavedrugreportwithocrResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwSavedrugreportwithocr error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询一张网药品信息通过mah、生产企业和批次号
*/
func (ability *Defaultability) AlibabaAlihealthSynergyYzwDruginfoQuery(req *request.AlibabaAlihealthSynergyYzwDruginfoQueryRequest) (*response.AlibabaAlihealthSynergyYzwDruginfoQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.synergy.yzw.druginfo.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthSynergyYzwDruginfoQueryResponse{}
	if err != nil {
		log.Println("alibabaAlihealthSynergyYzwDruginfoQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询上游出库单数量
*/
func (ability *Defaultability) AlibabaAlihealthDrugWesQueryUpoutbillcount(req *request.AlibabaAlihealthDrugWesQueryUpoutbillcountRequest) (*response.AlibabaAlihealthDrugWesQueryUpoutbillcountResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.wes.query.upoutbillcount", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugWesQueryUpoutbillcountResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugWesQueryUpoutbillcount error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询本企业出入库单数量
*/
func (ability *Defaultability) AlibabaAlihealthDrugWesQueryBillcount(req *request.AlibabaAlihealthDrugWesQueryBillcountRequest) (*response.AlibabaAlihealthDrugWesQueryBillcountResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.wes.query.billcount", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugWesQueryBillcountResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugWesQueryBillcount error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
wes码关联关系查询
*/
func (ability *Defaultability) AlibabaAlihealthDrugCodeKytWesQuerycoderelation(req *request.AlibabaAlihealthDrugCodeKytWesQuerycoderelationRequest) (*response.AlibabaAlihealthDrugCodeKytWesQuerycoderelationResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.code.kyt.wes.querycoderelation", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugCodeKytWesQuerycoderelationResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugCodeKytWesQuerycoderelation error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
生产批发单据上传
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesUploadcircubill(req *request.AlibabaAlihealthDrugKytWesUploadcircubillRequest) (*response.AlibabaAlihealthDrugKytWesUploadcircubillResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.uploadcircubill", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesUploadcircubillResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesUploadcircubill error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
WES上游出库单追溯单据拒收
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesBillpartialreject(req *request.AlibabaAlihealthDrugKytWesBillpartialrejectRequest) (*response.AlibabaAlihealthDrugKytWesBillpartialrejectResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.billpartialreject", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesBillpartialrejectResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesBillpartialreject error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询往来单位列表
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesListparts(req *request.AlibabaAlihealthDrugKytWesListpartsRequest) (*response.AlibabaAlihealthDrugKytWesListpartsResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.listparts", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesListpartsResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesListparts error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
单据删除
*/
func (ability *Defaultability) AlibabaAlihealthDrugCodeKytWesDelbillinfo(req *request.AlibabaAlihealthDrugCodeKytWesDelbillinfoRequest) (*response.AlibabaAlihealthDrugCodeKytWesDelbillinfoResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.code.kyt.wes.delbillinfo", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugCodeKytWesDelbillinfoResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugCodeKytWesDelbillinfo error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
物流企业查询货主企业信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesSynonymauths(req *request.AlibabaAlihealthDrugKytWesSynonymauthsRequest) (*response.AlibabaAlihealthDrugKytWesSynonymauthsResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.synonymauths", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesSynonymauthsResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesSynonymauths error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询上游出库单明细（带追溯码信息）
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesUpbillDetailwithcode(req *request.AlibabaAlihealthDrugKytWesUpbillDetailwithcodeRequest) (*response.AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.upbill.detailwithcode", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesUpbillDetailwithcodeResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesUpbillDetailwithcode error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
企业上传出入库信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesUploadinoutbill(req *request.AlibabaAlihealthDrugKytWesUploadinoutbillRequest) (*response.AlibabaAlihealthDrugKytWesUploadinoutbillResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.uploadinoutbill", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesUploadinoutbillResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesUploadinoutbill error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
单据处理状态查询
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesSearchstatus(req *request.AlibabaAlihealthDrugKytWesSearchstatusRequest) (*response.AlibabaAlihealthDrugKytWesSearchstatusResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.searchstatus", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesSearchstatusResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesSearchstatus error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询企业上传过的单据信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugCodeKytWesQuerycodebillinfo(req *request.AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoRequest) (*response.AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.code.kyt.wes.querycodebillinfo", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugCodeKytWesQuerycodebillinfoResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugCodeKytWesQuerycodebillinfo error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
部分处理成功单据处理失败的码明细
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesListbillprocesspartsuccess(req *request.AlibabaAlihealthDrugKytWesListbillprocesspartsuccessRequest) (*response.AlibabaAlihealthDrugKytWesListbillprocesspartsuccessResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.listbillprocesspartsuccess", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesListbillprocesspartsuccessResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesListbillprocesspartsuccess error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
wes查询追溯码信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugCodeKytWesQuerycode(req *request.AlibabaAlihealthDrugCodeKytWesQuerycodeRequest) (*response.AlibabaAlihealthDrugCodeKytWesQuerycodeResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.code.kyt.wes.querycode", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugCodeKytWesQuerycodeResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugCodeKytWesQuerycode error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取一个wes接口可用的token，每个企业每秒5次
*/
func (ability *Defaultability) AlibabaAlihealthDrugCodeKytWesLicenseTokenGet(req *request.AlibabaAlihealthDrugCodeKytWesLicenseTokenGetRequest) (*response.AlibabaAlihealthDrugCodeKytWesLicenseTokenGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.code.kyt.wes.license.token.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugCodeKytWesLicenseTokenGetResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugCodeKytWesLicenseTokenGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
由于无法返回多条匹配信息，本接口已不再推荐使用，推荐使用新接口，新接口名称：alibaba.alihealth.drug.wes.getentinfolist
*/
func (ability *Defaultability) AlibabaAlihealthDrugWesGetentinfonew(req *request.AlibabaAlihealthDrugWesGetentinfonewRequest) (*response.AlibabaAlihealthDrugWesGetentinfonewResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.wes.getentinfonew", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugWesGetentinfonewResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugWesGetentinfonew error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
企业信息-根据企业名称或证号查询企业信息列表
*/
func (ability *Defaultability) AlibabaAlihealthDrugWesGetentinfolist(req *request.AlibabaAlihealthDrugWesGetentinfolistRequest) (*response.AlibabaAlihealthDrugWesGetentinfolistResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.wes.getentinfolist", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugWesGetentinfolistResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugWesGetentinfolist error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
药品全量数据下载
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesGetdruginfoDownloadurl(req *request.AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlRequest) (*response.AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.getdruginfo.downloadurl", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesGetdruginfoDownloadurl error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
通过一个码查询上游出库单
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesQueryUpbillcode(req *request.AlibabaAlihealthDrugKytWesQueryUpbillcodeRequest) (*response.AlibabaAlihealthDrugKytWesQueryUpbillcodeResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.query.upbillcode", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesQueryUpbillcodeResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesQueryUpbillcode error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
物流代货主查找往来单位接口
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesListpartsByagent(req *request.AlibabaAlihealthDrugKytWesListpartsByagentRequest) (*response.AlibabaAlihealthDrugKytWesListpartsByagentResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.listparts.byagent", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesListpartsByagentResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesListpartsByagent error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
即将下线，使用alibaba.alihealth.drug.code.kyt.wes.license.token.get
*/
func (ability *Defaultability) AlibabaAlihealthDrugCodeKytWesGetlicense(req *request.AlibabaAlihealthDrugCodeKytWesGetlicenseRequest) (*response.AlibabaAlihealthDrugCodeKytWesGetlicenseResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.code.kyt.wes.getlicense", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugCodeKytWesGetlicenseResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugCodeKytWesGetlicense error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询码是否激活
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesQuerycodeactive(req *request.AlibabaAlihealthDrugKytWesQuerycodeactiveRequest) (*response.AlibabaAlihealthDrugKytWesQuerycodeactiveResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.querycodeactive", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesQuerycodeactiveResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesQuerycodeactive error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
新增往来单位企业记录
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesSaveent(req *request.AlibabaAlihealthDrugKytWesSaveentRequest) (*response.AlibabaAlihealthDrugKytWesSaveentResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.saveent", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesSaveentResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesSaveent error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
wes零头出入库单据上传
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesRemnantbillUpload(req *request.AlibabaAlihealthDrugKytWesRemnantbillUploadRequest) (*response.AlibabaAlihealthDrugKytWesRemnantbillUploadResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.remnantbill.upload", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesRemnantbillUploadResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesRemnantbillUpload error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
检查输入的码之间是否有上下级关系
*/
func (ability *Defaultability) AlibabaAlihealthDrugCodeKytWesCheckcoderelation(req *request.AlibabaAlihealthDrugCodeKytWesCheckcoderelationRequest) (*response.AlibabaAlihealthDrugCodeKytWesCheckcoderelationResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.code.kyt.wes.checkcoderelation", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugCodeKytWesCheckcoderelationResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugCodeKytWesCheckcoderelation error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
批号信息等修改时影响的单据号
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesQueryupoutbilllog(req *request.AlibabaAlihealthDrugKytWesQueryupoutbilllogRequest) (*response.AlibabaAlihealthDrugKytWesQueryupoutbilllogResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.queryupoutbilllog", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesQueryupoutbilllogResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesQueryupoutbilllog error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询单据详情
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesSearchbillDetail(req *request.AlibabaAlihealthDrugKytWesSearchbillDetailRequest) (*response.AlibabaAlihealthDrugKytWesSearchbillDetailResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.searchbill.detail", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesSearchbillDetailResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesSearchbillDetail error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询药品码段信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesDrugrescode(req *request.AlibabaAlihealthDrugKytWesDrugrescodeRequest) (*response.AlibabaAlihealthDrugKytWesDrugrescodeResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.drugrescode", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesDrugrescodeResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesDrugrescode error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据企业主键查看企业详细信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesGetbyentid(req *request.AlibabaAlihealthDrugKytWesGetbyentidRequest) (*response.AlibabaAlihealthDrugKytWesGetbyentidResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.getbyentid", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesGetbyentidResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesGetbyentid error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
通过时间段批量查询入出库单信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesSearchbill(req *request.AlibabaAlihealthDrugKytWesSearchbillRequest) (*response.AlibabaAlihealthDrugKytWesSearchbillResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.searchbill", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesSearchbillResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesSearchbill error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
由于无法返回多条匹配信息，本接口已不再推荐使用，推荐使用新接口，新接口名称：alibaba.alihealth.drug.wes.getentinfolist
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesGetentinfo(req *request.AlibabaAlihealthDrugKytWesGetentinfoRequest) (*response.AlibabaAlihealthDrugKytWesGetentinfoResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.getentinfo", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesGetentinfoResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesGetentinfo error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据企业唯一标识查看企业详细信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesGetbyrefentid(req *request.AlibabaAlihealthDrugKytWesGetbyrefentidRequest) (*response.AlibabaAlihealthDrugKytWesGetbyrefentidResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.getbyrefentid", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesGetbyrefentidResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesGetbyrefentid error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据统一信用代码查询企业信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesGetbyorgcode(req *request.AlibabaAlihealthDrugKytWesGetbyorgcodeRequest) (*response.AlibabaAlihealthDrugKytWesGetbyorgcodeResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.getbyorgcode", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesGetbyorgcodeResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesGetbyorgcode error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
企业搜索自己授权的物流企业
*/
func (ability *Defaultability) AlibabaAlihealthDrugWesListauths(req *request.AlibabaAlihealthDrugWesListauthsRequest) (*response.AlibabaAlihealthDrugWesListauthsResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.wes.listauths", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugWesListauthsResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugWesListauths error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询wes服务期时间
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesServiceInfo(req *request.AlibabaAlihealthDrugKytWesServiceInfoRequest) (*response.AlibabaAlihealthDrugKytWesServiceInfoResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.service.info", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesServiceInfoResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesServiceInfo error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询货主/本企业上游企业出库单据信息
*/
func (ability *Defaultability) AlibabaAlihealthDrugKytWesListupout(req *request.AlibabaAlihealthDrugKytWesListupoutRequest) (*response.AlibabaAlihealthDrugKytWesListupoutResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alihealth.drug.kyt.wes.listupout", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlihealthDrugKytWesListupoutResponse{}
	if err != nil {
		log.Println("alibabaAlihealthDrugKytWesListupout error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
