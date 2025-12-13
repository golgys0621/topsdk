package domain

import (
	"github.com/golgys0621/topsdk/util"
)

type AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto struct {
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
	   生产日期     */
	OriginalProduceDate *string `json:"original_produce_date,omitempty" `

	/*
	   有效期至     */
	OriginalExpireDate *string `json:"original_expire_date,omitempty" `
}

func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto) SetProduceDate(v util.LocalTime) *AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto {
	s.ProduceDate = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto) SetExpireDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto {
	s.ExpireDate = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto) SetBatchNo(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto {
	s.BatchNo = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto) SetCode(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto {
	s.Code = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto) SetPkgAmount(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto {
	s.PkgAmount = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto) SetOriginalProduceDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto {
	s.OriginalProduceDate = &v
	return s
}
func (s *AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto) SetOriginalExpireDate(v string) *AlibabaAlihealthDrugtraceTopYljgQueryRelationProduceInfoDto {
	s.OriginalExpireDate = &v
	return s
}
