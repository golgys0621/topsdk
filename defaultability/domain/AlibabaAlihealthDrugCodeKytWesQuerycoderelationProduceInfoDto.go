package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto struct {
	/*
	   生产日期     */
	ProduceDate *util.LocalTime `json:"produce_date,omitempty" `

	/*
	   有效期     */
	ExpireDate *string `json:"expire_date,omitempty" `

	/*
	   批次号     */
	BatchNo *string `json:"batch_no,omitempty" `

	/*
	   码     */
	Code *string `json:"code,omitempty" `

	/*
	   最小包装数量     */
	PkgAmount *string `json:"pkg_amount,omitempty" `

	/*
	   授权企业id     */
	AuthrizerEntId *string `json:"authrizer_ent_id,omitempty" `

	/*
	   生产企业id     */
	ProduceEntId *string `json:"produce_ent_id,omitempty" `

	/*
	   生产日期     */
	OriginalProduceDate *string `json:"original_produce_date,omitempty" `

	/*
	   有效期至     */
	OriginalExpireDate *string `json:"original_expire_date,omitempty" `
}

func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto) SetProduceDate(v util.LocalTime) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto {
	s.ProduceDate = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto) SetExpireDate(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto {
	s.ExpireDate = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto) SetBatchNo(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto {
	s.BatchNo = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto) SetCode(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto {
	s.Code = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto) SetPkgAmount(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto {
	s.PkgAmount = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto) SetAuthrizerEntId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto {
	s.AuthrizerEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto) SetProduceEntId(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto {
	s.ProduceEntId = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto) SetOriginalProduceDate(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto {
	s.OriginalProduceDate = &v
	return s
}
func (s *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto) SetOriginalExpireDate(v string) *AlibabaAlihealthDrugCodeKytWesQuerycoderelationProduceInfoDto {
	s.OriginalExpireDate = &v
	return s
}
