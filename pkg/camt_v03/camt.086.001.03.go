// Code generated by xsdgen. DO NOT EDIT.

package camt_v03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Othr"`
}

// May be one of INTM, SMRY
type AccountLevel1Code string

// May be one of INTM, SMRY, DETL
type AccountLevel2Code string

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Prtry"`
}

type AccountTax1 struct {
	ClctnMtd   BillingTaxCalculationMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 ClctnMtd"`
	Rgn        Max40Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Rgn,omitempty"`
	NonResCtry ResidenceLocation1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 NonResCtry,omitempty"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Prtry"`
}

type AmountAndDirection34 struct {
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Amt"`
	Sgn bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Sgn"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BalanceAdjustment1 struct {
	Tp                BalanceAdjustmentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Tp"`
	Desc              Max105Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Desc"`
	BalAmt            AmountAndDirection34       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BalAmt"`
	AvrgAmt           AmountAndDirection34       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 AvrgAmt,omitempty"`
	ErrDt             ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 ErrDt,omitempty"`
	PstngDt           ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PstngDt"`
	Days              float64                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Days,omitempty"`
	EarngsAdjstmntAmt AmountAndDirection34       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 EarngsAdjstmntAmt,omitempty"`
}

// May be one of LDGR, FLOT, CLLD
type BalanceAdjustmentType1Code string

type BankServicesBillingStatementV03 struct {
	RptHdr      ReportHeader6     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 RptHdr"`
	BllgStmtGrp []StatementGroup3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BllgStmtGrp"`
}

type BankTransactionCodeStructure4 struct {
	Domn  BankTransactionCodeStructure5            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Domn,omitempty"`
	Prtry ProprietaryBankTransactionCodeStructure1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Prtry,omitempty"`
}

type BankTransactionCodeStructure5 struct {
	Cd   ExternalBankTransactionDomain1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	Fmly BankTransactionCodeStructure6      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Fmly"`
}

type BankTransactionCodeStructure6 struct {
	Cd        ExternalBankTransactionFamily1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	SubFmlyCd ExternalBankTransactionSubFamily1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SubFmlyCd"`
}

type BillingBalance1 struct {
	Tp    BillingBalanceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Tp"`
	Val   AmountAndDirection34      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Val"`
	CcyTp BillingCurrencyType1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 CcyTp,omitempty"`
}

type BillingBalanceType1Choice struct {
	Cd    ExternalBillingBalanceType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	Prtry Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Prtry"`
}

// May be one of UPRC, STAM, BCHG, DPRC, FCHG, LPRC, MCHG, MXRD, TIR1, TIR2, TIR3, TIR4, TIR5, TIR6, TIR7, TIR8, TIR9, TPRC, ZPRC, BBSE
type BillingChargeMethod1Code string

type BillingCompensation1 struct {
	Tp    BillingCompensationType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Tp"`
	Val   AmountAndDirection34           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Val"`
	CcyTp BillingCurrencyType2Code       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 CcyTp,omitempty"`
}

type BillingCompensationType1Choice struct {
	Cd    ExternalBillingCompensationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	Prtry Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Prtry"`
}

// May be one of ACCT, STLM, PRCG
type BillingCurrencyType1Code string

// May be one of ACCT, STLM, PRCG, HOST
type BillingCurrencyType2Code string

type BillingMethod1 struct {
	SvcChrgHstAmt AmountAndDirection34   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SvcChrgHstAmt"`
	SvcTax        BillingServicesAmount1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SvcTax"`
	TtlChrg       BillingServicesAmount2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TtlChrg"`
	TaxId         []BillingServicesTax1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TaxId"`
}

type BillingMethod1Choice struct {
	MtdA BillingMethod1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 MtdA"`
	MtdB BillingMethod2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 MtdB"`
	MtdD BillingMethod3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 MtdD"`
}

type BillingMethod2 struct {
	SvcChrgHstAmt AmountAndDirection34   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SvcChrgHstAmt"`
	SvcTax        BillingServicesAmount1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SvcTax"`
	TaxId         []BillingServicesTax1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TaxId"`
}

type BillingMethod3 struct {
	SvcTaxPricAmt AmountAndDirection34  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SvcTaxPricAmt"`
	TaxId         []BillingServicesTax2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TaxId"`
}

type BillingMethod4 struct {
	SvcDtl   []BillingServiceParameters2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SvcDtl"`
	TaxClctn TaxCalculation1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TaxClctn"`
}

type BillingPrice1 struct {
	Ccy      ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Ccy,omitempty"`
	UnitPric AmountAndDirection34         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 UnitPric,omitempty"`
	Mtd      BillingChargeMethod1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Mtd,omitempty"`
	Rule     Max20Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Rule,omitempty"`
}

type BillingRate1 struct {
	Id        BillingRateIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Id"`
	Val       float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Val"`
	DaysInPrd float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 DaysInPrd,omitempty"`
	DaysInYr  float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 DaysInYr,omitempty"`
}

type BillingRateIdentification1Choice struct {
	Cd    ExternalBillingRateIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	Prtry Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Prtry"`
}

type BillingService2 struct {
	SvcDtl            BillingServiceParameters3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SvcDtl"`
	Pric              BillingPrice1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Pric,omitempty"`
	PmtMtd            ServicePaymentMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PmtMtd"`
	OrgnlChrgPric     AmountAndDirection34      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 OrgnlChrgPric"`
	OrgnlChrgSttlmAmt AmountAndDirection34      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 OrgnlChrgSttlmAmt,omitempty"`
	BalReqrdAcctAmt   AmountAndDirection34      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BalReqrdAcctAmt,omitempty"`
	TaxDsgnt          ServiceTaxDesignation1    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TaxDsgnt"`
	TaxClctn          BillingMethod1Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TaxClctn,omitempty"`
}

type BillingServiceAdjustment1 struct {
	Tp           ServiceAdjustmentType1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Tp"`
	Desc         Max140Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Desc"`
	Amt          AmountAndDirection34             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Amt"`
	BalReqrdAmt  AmountAndDirection34             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BalReqrdAmt,omitempty"`
	ErrDt        ISODate                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 ErrDt,omitempty"`
	AdjstmntId   Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 AdjstmntId,omitempty"`
	SubSvc       BillingSubServiceIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SubSvc,omitempty"`
	PricChng     AmountAndDirection34             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PricChng,omitempty"`
	OrgnlPric    AmountAndDirection34             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 OrgnlPric,omitempty"`
	NewPric      AmountAndDirection34             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 NewPric,omitempty"`
	VolChng      float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 VolChng,omitempty"`
	OrgnlVol     float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 OrgnlVol,omitempty"`
	NewVol       float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 NewVol,omitempty"`
	OrgnlChrgAmt AmountAndDirection34             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 OrgnlChrgAmt,omitempty"`
	NewChrgAmt   AmountAndDirection34             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 NewChrgAmt,omitempty"`
}

type BillingServiceCommonIdentification1 struct {
	Issr Max6Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Issr"`
	Id   Max8Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Id"`
}

type BillingServiceIdentification2 struct {
	Id     Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Id"`
	SubSvc BillingSubServiceIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SubSvc,omitempty"`
	Desc   Max70Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Desc"`
}

type BillingServiceIdentification3 struct {
	Id     Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Id"`
	SubSvc BillingSubServiceIdentification1    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SubSvc,omitempty"`
	Desc   Max70Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Desc"`
	CmonCd BillingServiceCommonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 CmonCd,omitempty"`
	BkTxCd BankTransactionCodeStructure4       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BkTxCd,omitempty"`
	SvcTp  Max12Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SvcTp,omitempty"`
}

type BillingServiceParameters2 struct {
	BkSvc      BillingServiceIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BkSvc"`
	Vol        float64                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Vol,omitempty"`
	UnitPric   AmountAndDirection34          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 UnitPric,omitempty"`
	SvcChrgAmt AmountAndDirection34          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SvcChrgAmt"`
}

type BillingServiceParameters3 struct {
	BkSvc BillingServiceIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BkSvc"`
	Vol   float64                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Vol,omitempty"`
}

type BillingServicesAmount1 struct {
	HstAmt   AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 HstAmt"`
	PricgAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PricgAmt,omitempty"`
}

type BillingServicesAmount2 struct {
	HstAmt   AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 HstAmt"`
	SttlmAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SttlmAmt,omitempty"`
	PricgAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PricgAmt,omitempty"`
}

type BillingServicesAmount3 struct {
	SrcAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SrcAmt"`
	HstAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 HstAmt"`
}

type BillingServicesTax1 struct {
	Nb       Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Nb"`
	Desc     Max40Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Desc,omitempty"`
	Rate     float64              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Rate"`
	HstAmt   AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 HstAmt"`
	PricgAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PricgAmt,omitempty"`
}

type BillingServicesTax2 struct {
	Nb       Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Nb"`
	Desc     Max40Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Desc,omitempty"`
	Rate     float64              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Rate"`
	PricgAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PricgAmt"`
}

type BillingServicesTax3 struct {
	Nb        Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Nb"`
	Desc      Max40Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Desc,omitempty"`
	Rate      float64              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Rate"`
	TtlTaxAmt AmountAndDirection34 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TtlTaxAmt"`
}

type BillingStatement3 struct {
	StmtId      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 StmtId"`
	FrToDt      DatePeriod1                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 FrToDt"`
	CreDtTm     ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 CreDtTm"`
	Sts         BillingStatementStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Sts"`
	AcctChrtcs  CashAccountCharacteristics3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 AcctChrtcs"`
	RateData    []BillingRate1              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 RateData,omitempty"`
	CcyXchg     []CurrencyExchange6         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 CcyXchg,omitempty"`
	Bal         []BillingBalance1           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Bal,omitempty"`
	Compstn     []BillingCompensation1      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Compstn,omitempty"`
	Svc         []BillingService2           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Svc,omitempty"`
	TaxRgn      []BillingTaxRegion2         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TaxRgn,omitempty"`
	BalAdjstmnt []BalanceAdjustment1        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BalAdjstmnt,omitempty"`
	SvcAdjstmnt []BillingServiceAdjustment1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SvcAdjstmnt,omitempty"`
}

// May be one of ORGN, RPLC, TEST
type BillingStatementStatus1Code string

type BillingSubServiceIdentification1 struct {
	Issr BillingSubServiceQualifier1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Issr"`
	Id   Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Id"`
}

type BillingSubServiceQualifier1Choice struct {
	Cd    BillingSubServiceQualifier1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	Prtry Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Prtry"`
}

// May be one of LBOX, STOR, BILA, SEQN, MACT
type BillingSubServiceQualifier1Code string

// May be one of NTAX, MTDA, MTDB, MTDC, MTDD, UDFD
type BillingTaxCalculationMethod1Code string

type BillingTaxIdentification2 struct {
	VATRegnNb Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 VATRegnNb,omitempty"`
	TaxRegnNb Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TaxRegnNb,omitempty"`
	TaxCtct   Contact4  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TaxCtct,omitempty"`
}

type BillingTaxRegion2 struct {
	RgnNb       Max40Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 RgnNb"`
	RgnNm       Max40Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 RgnNm"`
	CstmrTaxId  Max40Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 CstmrTaxId"`
	PtDt        ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PtDt,omitempty"`
	SndgFI      BillingTaxIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SndgFI,omitempty"`
	InvcNb      Max40Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 InvcNb,omitempty"`
	MtdC        BillingMethod4            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 MtdC,omitempty"`
	SttlmAmt    AmountAndDirection34      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SttlmAmt"`
	TaxDueToRgn AmountAndDirection34      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TaxDueToRgn"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PstlAdr,omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Id"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Prxy,omitempty"`
}

type CashAccountCharacteristics3 struct {
	AcctLvl      AccountLevel2Code                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 AcctLvl"`
	CshAcct      CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 CshAcct"`
	AcctSvcr     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 AcctSvcr,omitempty"`
	PrntAcct     ParentCashAccount3                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PrntAcct,omitempty"`
	CompstnMtd   CompensationMethod1Code                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 CompstnMtd"`
	DbtAcct      AccountIdentification4Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 DbtAcct,omitempty"`
	DelydDbtDt   ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 DelydDbtDt,omitempty"`
	SttlmAdvc    Max105Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SttlmAdvc,omitempty"`
	AcctBalCcyCd ActiveOrHistoricCurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 AcctBalCcyCd"`
	SttlmCcyCd   ActiveOrHistoricCurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SttlmCcyCd,omitempty"`
	HstCcyCd     ActiveOrHistoricCurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 HstCcyCd,omitempty"`
	Tax          AccountTax1                                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Tax,omitempty"`
	AcctSvcrCtct Contact4                                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 AcctSvcrCtct"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 MmbId"`
}

// May be one of NOCP, DBTD, INVD, DDBT
type CompensationMethod1Code string

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PrefrdMtd,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyExchange6 struct {
	SrcCcy   ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SrcCcy"`
	TrgtCcy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TrgtCcy"`
	XchgRate float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 XchgRate"`
	Desc     Max40Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Desc,omitempty"`
	UnitCcy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 UnitCcy,omitempty"`
	Cmnts    Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cmnts,omitempty"`
	QtnDt    ISODateTime                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 QtnDt,omitempty"`
}

type DatePeriod1 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 FrDt,omitempty"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 ToDt"`
}

type Document struct {
	BkSvcsBllgStmt BankServicesBillingStatementV03 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BkSvcsBllgStmt"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalBankTransactionDomain1Code string

// Must be at least 1 items long
type ExternalBankTransactionFamily1Code string

// Must be at least 1 items long
type ExternalBankTransactionSubFamily1Code string

// Must be at least 1 items long
type ExternalBillingBalanceType1Code string

// Must be at least 1 items long
type ExternalBillingCompensationType1Code string

// Must be at least 1 items long
type ExternalBillingRateIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Othr,omitempty"`
}

type FinancialInstitutionIdentification19 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 LEI,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Issr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max10Text string

// Must be at least 1 items long
type Max128Text string

// Must be at least 1 items long
type Max12Text string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max20Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max40Text string

// Must be at least 1 items long
type Max4Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

// Must be at least 1 items long
type Max8Text string

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Prtry"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Id,omitempty"`
}

type Pagination1 struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 LastPgInd"`
}

type ParentCashAccount3 struct {
	Lvl  AccountLevel1Code                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Lvl,omitempty"`
	Id   CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Id"`
	Svcr BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Svcr,omitempty"`
}

type Party43Choice struct {
	OrgId OrganisationIdentification29         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 OrgId"`
	FIId  FinancialInstitutionIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 FIId"`
}

type PartyIdentification138 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Nm"`
	LglNm     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 LglNm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PstlAdr,omitempty"`
	Id        Party43Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Id"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 CtctDtls,omitempty"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

type ProprietaryBankTransactionCodeStructure1 struct {
	Cd   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Issr,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Prtry"`
}

type ReportHeader6 struct {
	RptId    Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 RptId"`
	MsgPgntn Pagination1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 MsgPgntn,omitempty"`
}

type ResidenceLocation1Choice struct {
	Ctry CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Ctry"`
	Area Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Area"`
}

// May be one of COMP, NCMP
type ServiceAdjustmentType1Code string

// May be one of BCMP, FLAT, PVCH, INVS, WVED, FREE
type ServicePaymentMethod1Code string

type ServiceTaxDesignation1 struct {
	Cd     ServiceTaxDesignation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	Rgn    Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Rgn,omitempty"`
	TaxRsn []TaxReason1               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TaxRsn,omitempty"`
}

// May be one of XMPT, ZERO, TAXE
type ServiceTaxDesignation1Code string

type StatementGroup3 struct {
	GrpId        Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 GrpId"`
	Sndr         PartyIdentification138 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Sndr"`
	SndrIndvCtct []Contact4             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 SndrIndvCtct,omitempty"`
	Rcvr         PartyIdentification138 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Rcvr"`
	RcvrIndvCtct []Contact4             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 RcvrIndvCtct,omitempty"`
	BllgStmt     []BillingStatement3    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BllgStmt"`
}

type TaxCalculation1 struct {
	HstCcy                ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 HstCcy"`
	TaxblSvcChrgConvs     []BillingServicesAmount3     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TaxblSvcChrgConvs"`
	TtlTaxblSvcChrgHstAmt AmountAndDirection34         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TtlTaxblSvcChrgHstAmt"`
	TaxId                 []BillingServicesTax3        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TaxId"`
	TtlTax                AmountAndDirection34         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 TtlTax"`
}

type TaxReason1 struct {
	Cd     Max10Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Cd"`
	Expltn Max105Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 Expltn"`
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02")), nil
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
