// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v07

import "github.com/moov-io/iso20022/pkg/common"

type AccountCriteria3Choice struct {
	QryNm   common.Max35Text `xml:"QryNm"`
	NewCrit AccountCriteria7 `xml:"NewCrit"`
}

type AccountCriteria7 struct {
	NewQryNm *common.Max35Text            `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  []CashAccountSearchCriteria7 `xml:"SchCrit,omitempty" json:",omitempty"`
	RtrCrit  *CashAccountReturnCriteria5  `xml:"RtrCrit,omitempty" json:",omitempty"`
}

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"IBAN"`
	Othr GenericAccountIdentification1 `xml:"Othr"`
}

type AccountIdentificationSearchCriteria2Choice struct {
	EQ     AccountIdentification4Choice `xml:"EQ"`
	CTTxt  common.Max35Text             `xml:"CTTxt"`
	NCTTxt common.Max35Text             `xml:"NCTTxt"`
}

type AccountQuery3 struct {
	QryTp    *QueryType2Code         `xml:"QryTp,omitempty" json:",omitempty"`
	AcctCrit *AccountCriteria3Choice `xml:"AcctCrit,omitempty" json:",omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                   `xml:"Prtry"`
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"Cd"`
	Prtry GenericIdentification30 `xml:"Prtry"`
}

type BalanceType11Choice struct {
	Cd    ExternalSystemBalanceType1Code `xml:"Cd"`
	Prtry common.Max35Text               `xml:"Prtry"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"FinInstnId"`
	BrnchId    *BranchData3                         `xml:"BrnchId,omitempty" json:",omitempty"`
}

type BranchData3 struct {
	Id      *common.Max35Text     `xml:"Id,omitempty" json:",omitempty"`
	LEI     *common.LEIIdentifier `xml:"LEI,omitempty" json:",omitempty"`
	Nm      *common.Max140Text    `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr *PostalAddress24      `xml:"PstlAdr,omitempty" json:",omitempty"`
}

type CashAccountReturnCriteria5 struct {
	NmInd         bool                        `xml:"NmInd,omitempty" json:",omitempty"`
	CcyInd        bool                        `xml:"CcyInd,omitempty" json:",omitempty"`
	TpInd         bool                        `xml:"TpInd,omitempty" json:",omitempty"`
	MulLmtInd     bool                        `xml:"MulLmtInd,omitempty" json:",omitempty"`
	MulBalRtrCrit *CashBalanceReturnCriteria2 `xml:"MulBalRtrCrit,omitempty" json:",omitempty"`
	BilLmtInd     bool                        `xml:"BilLmtInd,omitempty" json:",omitempty"`
	BilBalRtrCrit *CashBalanceReturnCriteria2 `xml:"BilBalRtrCrit,omitempty" json:",omitempty"`
	StgOrdrInd    bool                        `xml:"StgOrdrInd,omitempty" json:",omitempty"`
	AcctOwnrInd   bool                        `xml:"AcctOwnrInd,omitempty" json:",omitempty"`
	AcctSvcrInd   bool                        `xml:"AcctSvcrInd,omitempty" json:",omitempty"`
}

type CashAccountSearchCriteria7 struct {
	AcctId   []AccountIdentificationSearchCriteria2Choice  `xml:"AcctId,omitempty" json:",omitempty"`
	Tp       []CashAccountType2Choice                      `xml:"Tp,omitempty" json:",omitempty"`
	Ccy      []common.ActiveOrHistoricCurrencyCode         `xml:"Ccy,omitempty" json:",omitempty"`
	Bal      []CashBalance12                               `xml:"Bal,omitempty" json:",omitempty"`
	AcctOwnr *PartyIdentification135                       `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctSvcr *BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type CashBalance12 struct {
	Tp       []BalanceType11Choice                          `xml:"Tp,omitempty" json:",omitempty"`
	CtrPtyTp *BalanceCounterparty1Code                      `xml:"CtrPtyTp"`
	CtrPtyId []BranchAndFinancialInstitutionIdentification6 `xml:"CtrPtyId,omitempty" json:",omitempty"`
	ValDt    []DateAndDateTimeSearch4Choice                 `xml:"ValDt,omitempty" json:",omitempty"`
	PrcgDt   *DateAndDateTimeSearch4Choice                  `xml:"PrcgDt,omitempty" json:",omitempty"`
}

type CashBalanceReturnCriteria2 struct {
	TpInd       bool `xml:"TpInd"`
	StsInd      bool `xml:"StsInd"`
	ValDtInd    bool `xml:"ValDtInd"`
	PrcgDtInd   bool `xml:"PrcgDtInd"`
	NbOfPmtsInd bool `xml:"NbOfPmtsInd"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                     `xml:"MmbId"`
}

type Contact4 struct {
	NmPrfx    *common.NamePrefix2Code      `xml:"NmPrfx,omitempty" json:",omitempty"`
	Nm        *common.Max140Text           `xml:"Nm,omitempty" json:",omitempty"`
	PhneNb    *common.PhoneNumber          `xml:"PhneNb,omitempty" json:",omitempty"`
	MobNb     *common.PhoneNumber          `xml:"MobNb,omitempty" json:",omitempty"`
	FaxNb     *common.PhoneNumber          `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr  *common.Max2048Text          `xml:"EmailAdr,omitempty" json:",omitempty"`
	EmailPurp *common.Max35Text            `xml:"EmailPurp,omitempty" json:",omitempty"`
	JobTitl   *common.Max35Text            `xml:"JobTitl,omitempty" json:",omitempty"`
	Rspnsblty *common.Max35Text            `xml:"Rspnsblty,omitempty" json:",omitempty"`
	Dept      *common.Max70Text            `xml:"Dept,omitempty" json:",omitempty"`
	Othr      []OtherContact1              `xml:"Othr,omitempty" json:",omitempty"`
	PrefrdMtd *PreferredContactMethod1Code `xml:"PrefrdMtd,omitempty" json:",omitempty"`
}

type DateAndDateTimeSearch4Choice struct {
	DtTm DateTimeSearch2Choice   `xml:"DtTm"`
	Dt   DatePeriodSearch1Choice `xml:"Dt"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth *common.Max35Text  `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

type DatePeriodSearch1Choice struct {
	FrDt   common.ISODate `xml:"FrDt"`
	ToDt   common.ISODate `xml:"ToDt"`
	FrToDt DatePeriod2    `xml:"FrToDt"`
	EQDt   common.ISODate `xml:"EQDt"`
	NEQDt  common.ISODate `xml:"NEQDt"`
}

type DateTimePeriod1 struct {
	FrDtTm common.ISODateTime `xml:"FrDtTm"`
	ToDtTm common.ISODateTime `xml:"ToDtTm"`
}

type DateTimeSearch2Choice struct {
	FrDtTm   common.ISODateTime `xml:"FrDtTm"`
	ToDtTm   common.ISODateTime `xml:"ToDtTm"`
	FrToDtTm DateTimePeriod1    `xml:"FrToDtTm"`
	EQDtTm   common.ISODateTime `xml:"EQDtTm"`
	NEQDtTm  common.ISODateTime `xml:"NEQDtTm"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                                `xml:"Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       *common.BICFIDec2014Identifier       `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	LEI         *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Nm          *common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     *PostalAddress24                     `xml:"PstlAdr,omitempty" json:",omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text          `xml:"Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text         `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                          `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                         `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericIdentification1 struct {
	Id      common.Max35Text  `xml:"Id"`
	SchmeNm *common.Max35Text `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"Id"`
	Issr    common.Max35Text              `xml:"Issr"`
	SchmeNm *common.Max35Text             `xml:"SchmeNm,omitempty" json:",omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                             `xml:"Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                            `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                       `xml:"Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                      `xml:"Issr,omitempty" json:",omitempty"`
}

type GetAccountV07 struct {
	MsgHdr      MessageHeader9       `xml:"MsgHdr"`
	AcctQryDef  *AccountQuery3       `xml:"AcctQryDef,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type MessageHeader9 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
	ReqTp   *RequestType4Choice `xml:"ReqTp,omitempty" json:",omitempty"`
}

type OrganisationIdentification29 struct {
	AnyBIC *common.AnyBICDec2014Identifier      `xml:"AnyBIC,omitempty" json:",omitempty"`
	LEI    *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                        `xml:"Prtry"`
}

type OtherContact1 struct {
	ChanlTp common.Max4Text   `xml:"ChanlTp"`
	Id      common.Max128Text `xml:"Id,omitempty" json:",omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"OrgId"`
	PrvtId PersonIdentification13       `xml:"PrvtId"`
}

type PartyIdentification135 struct {
	Nm        *common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   *PostalAddress24    `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        *Party38Choice      `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes *common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  *Contact4           `xml:"CtctDtls,omitempty" json:",omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1          `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

type PostalAddress24 struct {
	AdrTp       *AddressType3Choice `xml:"AdrTp,omitempty" json:",omitempty"`
	Dept        *common.Max70Text   `xml:"Dept,omitempty" json:",omitempty"`
	SubDept     *common.Max70Text   `xml:"SubDept,omitempty" json:",omitempty"`
	StrtNm      *common.Max70Text   `xml:"StrtNm,omitempty" json:",omitempty"`
	BldgNb      *common.Max16Text   `xml:"BldgNb,omitempty" json:",omitempty"`
	BldgNm      *common.Max35Text   `xml:"BldgNm,omitempty" json:",omitempty"`
	Flr         *common.Max70Text   `xml:"Flr,omitempty" json:",omitempty"`
	PstBx       *common.Max16Text   `xml:"PstBx,omitempty" json:",omitempty"`
	Room        *common.Max70Text   `xml:"Room,omitempty" json:",omitempty"`
	PstCd       *common.Max16Text   `xml:"PstCd,omitempty" json:",omitempty"`
	TwnNm       *common.Max35Text   `xml:"TwnNm,omitempty" json:",omitempty"`
	TwnLctnNm   *common.Max35Text   `xml:"TwnLctnNm,omitempty" json:",omitempty"`
	DstrctNm    *common.Max35Text   `xml:"DstrctNm,omitempty" json:",omitempty"`
	CtrySubDvsn *common.Max35Text   `xml:"CtrySubDvsn,omitempty" json:",omitempty"`
	Ctry        *common.CountryCode `xml:"Ctry,omitempty" json:",omitempty"`
	AdrLine     []common.Max70Text  `xml:"AdrLine,omitempty" json:",omitempty"`
}

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"PmtCtrl"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"Enqry"`
	Prtry   GenericIdentification1                 `xml:"Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm *common.Max350Text         `xml:"PlcAndNm,omitempty" json:",omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type ActiveAmountRange3Choice struct {
	ImpldCcyAndAmtRg ImpliedCurrencyAndAmountRange1 `xml:"ImpldCcyAndAmtRg"`
	CcyAndAmtRg      ActiveCurrencyAndAmountRange3  `xml:"CcyAndAmtRg"`
}

type ActiveCurrencyAndAmountRange3 struct {
	Amt       ImpliedCurrencyAmountRange1Choice `xml:"Amt"`
	CdtDbtInd *common.CreditDebitCode           `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Ccy       common.ActiveCurrencyCode         `xml:"Ccy"`
}

type AmountRangeBoundary1 struct {
	BdryAmt float64 `xml:"BdryAmt"`
	Incl    bool    `xml:"Incl"`
}

type DateAndPeriod2Choice struct {
	Dt   common.ISODate `xml:"Dt"`
	Prd  Period2        `xml:"Prd"`
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

type FromToAmountRange1 struct {
	FrAmt AmountRangeBoundary1 `xml:"FrAmt"`
	ToAmt AmountRangeBoundary1 `xml:"ToAmt"`
}

type FromToPercentageRange1 struct {
	Fr PercentageRangeBoundary1 `xml:"Fr"`
	To PercentageRangeBoundary1 `xml:"To"`
}

type GetLimitV07 struct {
	MsgHdr      MessageHeader9       `xml:"MsgHdr"`
	LmtQryDef   *LimitQuery4         `xml:"LmtQryDef,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type ImpliedCurrencyAmountRange1Choice struct {
	FrAmt   AmountRangeBoundary1 `xml:"FrAmt"`
	ToAmt   AmountRangeBoundary1 `xml:"ToAmt"`
	FrToAmt FromToAmountRange1   `xml:"FrToAmt"`
	EQAmt   float64              `xml:"EQAmt"`
	NEQAmt  float64              `xml:"NEQAmt"`
}

type ImpliedCurrencyAndAmountRange1 struct {
	Amt       ImpliedCurrencyAmountRange1Choice `xml:"Amt"`
	CdtDbtInd *common.CreditDebitCode           `xml:"CdtDbtInd,omitempty" json:",omitempty"`
}

type LimitCriteria6 struct {
	NewQryNm *common.Max35Text      `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  []LimitSearchCriteria6 `xml:"SchCrit,omitempty" json:",omitempty"`
	RtrCrit  *LimitReturnCriteria2  `xml:"RtrCrit,omitempty" json:",omitempty"`
}

type LimitCriteria6Choice struct {
	QryNm   common.Max35Text `xml:"QryNm"`
	NewCrit LimitCriteria6   `xml:"NewCrit"`
}

type LimitQuery4 struct {
	QryTp   *QueryType2Code       `xml:"QryTp,omitempty" json:",omitempty"`
	LmtCrit *LimitCriteria6Choice `xml:"LmtCrit,omitempty" json:",omitempty"`
}

type LimitReturnCriteria2 struct {
	StartDtTmInd bool `xml:"StartDtTmInd,omitempty" json:",omitempty"`
	StsInd       bool `xml:"StsInd,omitempty" json:",omitempty"`
	UsdAmtInd    bool `xml:"UsdAmtInd,omitempty" json:",omitempty"`
	UsdPctgInd   bool `xml:"UsdPctgInd,omitempty" json:",omitempty"`
}

type LimitSearchCriteria6 struct {
	SysId          *SystemIdentification2Choice                   `xml:"SysId,omitempty" json:",omitempty"`
	BilLmtCtrPtyId []BranchAndFinancialInstitutionIdentification6 `xml:"BilLmtCtrPtyId,omitempty" json:",omitempty"`
	DfltLmtTp      []LimitType1Choice                             `xml:"DfltLmtTp,omitempty" json:",omitempty"`
	CurLmtTp       []LimitType1Choice                             `xml:"CurLmtTp,omitempty" json:",omitempty"`
	AcctOwnr       *BranchAndFinancialInstitutionIdentification6  `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctId         *AccountIdentification4Choice                  `xml:"AcctId,omitempty" json:",omitempty"`
	UsdAmt         *ActiveAmountRange3Choice                      `xml:"UsdAmt,omitempty" json:",omitempty"`
	UsdPctg        *PercentageRange1Choice                        `xml:"UsdPctg,omitempty" json:",omitempty"`
	LmtCcy         *common.ActiveCurrencyCode                     `xml:"LmtCcy,omitempty" json:",omitempty"`
	LmtAmt         *ActiveAmountRange3Choice                      `xml:"LmtAmt,omitempty" json:",omitempty"`
	LmtVldAsOfDt   *DateAndPeriod2Choice                          `xml:"LmtVldAsOfDt,omitempty" json:",omitempty"`
}

type LimitType1Choice struct {
	Cd    LimitType3Code   `xml:"Cd"`
	Prtry common.Max35Text `xml:"Prtry"`
}

type MarketInfrastructureIdentification1Choice struct {
	Cd    ExternalMarketInfrastructure1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

type PercentageRange1Choice struct {
	Fr   PercentageRangeBoundary1 `xml:"Fr"`
	To   PercentageRangeBoundary1 `xml:"To"`
	FrTo FromToPercentageRange1   `xml:"FrTo"`
	EQ   float64                  `xml:"EQ"`
	NEQ  float64                  `xml:"NEQ"`
}

type PercentageRangeBoundary1 struct {
	BdryRate float64 `xml:"BdryRate"`
	Incl     bool    `xml:"Incl"`
}

type Period2 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

type SystemIdentification2Choice struct {
	MktInfrstrctrId MarketInfrastructureIdentification1Choice `xml:"MktInfrstrctrId"`
	Ctry            common.CountryCode                        `xml:"Ctry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64                   `xml:",chardata"`
	Ccy   common.ActiveCurrencyCode `xml:"Ccy,attr"`
}

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"AmtWthtCcy"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"AmtWthCcy"`
}

type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"Dt"`
	DtTm common.ISODateTime `xml:"DtTm"`
}

type Limit8 struct {
	StartDtTm *DateAndDateTime2Choice `xml:"StartDtTm,omitempty" json:",omitempty"`
	Amt       Amount2Choice           `xml:"Amt"`
	CdtDbtInd *common.CreditDebitCode `xml:"CdtDbtInd,omitempty" json:",omitempty"`
}

type LimitIdentification2Choice struct {
	Cur     LimitIdentification5 `xml:"Cur"`
	Dflt    LimitIdentification5 `xml:"Dflt"`
	AllCur  LimitIdentification6 `xml:"AllCur"`
	AllDflt LimitIdentification6 `xml:"AllDflt"`
}

type LimitIdentification5 struct {
	SysId          *SystemIdentification2Choice                  `xml:"SysId,omitempty" json:",omitempty"`
	BilLmtCtrPtyId *BranchAndFinancialInstitutionIdentification6 `xml:"BilLmtCtrPtyId,omitempty" json:",omitempty"`
	Tp             LimitType1Choice                              `xml:"Tp"`
	AcctOwnr       *BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctId         *AccountIdentification4Choice                 `xml:"AcctId,omitempty" json:",omitempty"`
}

type LimitIdentification6 struct {
	SysId    *SystemIdentification2Choice                  `xml:"SysId,omitempty" json:",omitempty"`
	Tp       LimitType1Choice                              `xml:"Tp"`
	AcctOwnr *BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctId   *AccountIdentification4Choice                 `xml:"AcctId,omitempty" json:",omitempty"`
}

type LimitStructure3 struct {
	LmtId        LimitIdentification2Choice `xml:"LmtId"`
	NewLmtValSet Limit8                     `xml:"NewLmtValSet"`
}

type MessageHeader1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

type ModifyLimitV07 struct {
	MsgHdr      MessageHeader1       `xml:"MsgHdr"`
	LmtDtls     []LimitStructure3    `xml:"LmtDtls,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type DeleteLimitV07 struct {
	MsgHdr      MessageHeader1        `xml:"MsgHdr"`
	LmtDtls     LimitStructure2Choice `xml:"LmtDtls"`
	SplmtryData []SupplementaryData1  `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type LimitStructure2Choice struct {
	CurLmtId   LimitIdentification5 `xml:"CurLmtId"`
	AllCurLmts LimitIdentification6 `xml:"AllCurLmts"`
}

type BusinessDay8 struct {
	SysId       []SystemIdentification2Choice    `xml:"SysId"`
	BizDayOrErr BusinessDayReportOrError10Choice `xml:"BizDayOrErr"`
}

type BusinessDay9 struct {
	SysDt        *DateAndDateTime2Choice        `xml:"SysDt,omitempty" json:",omitempty"`
	SysSts       *SystemStatus3                 `xml:"SysSts,omitempty" json:",omitempty"`
	SysInfPerCcy []SystemAvailabilityAndEvents3 `xml:"SysInfPerCcy,omitempty" json:",omitempty"`
}

type BusinessDayReportOrError10Choice struct {
	BizDayInf BusinessDay9     `xml:"BizDayInf"`
	BizErr    []ErrorHandling5 `xml:"BizErr"`
}

type BusinessDayReportOrError9Choice struct {
	BizRpt  []BusinessDay8   `xml:"BizRpt"`
	OprlErr []ErrorHandling5 `xml:"OprlErr"`
}

type ClosureReason2Choice struct {
	Cd    SystemClosureReason1Code `xml:"Cd"`
	Prtry common.Max35Text         `xml:"Prtry"`
}

type DateTimePeriod1Choice struct {
	FrDtTm common.ISODateTime `xml:"FrDtTm"`
	ToDtTm common.ISODateTime `xml:"ToDtTm"`
	DtTmRg DateTimePeriod1    `xml:"DtTmRg"`
}

type ErrorHandling3Choice struct {
	Cd    ExternalSystemErrorHandling1Code `xml:"Cd"`
	Prtry common.Max35Text                 `xml:"Prtry"`
}

type ErrorHandling5 struct {
	Err  ErrorHandling3Choice `xml:"Err"`
	Desc *common.Max140Text   `xml:"Desc,omitempty" json:",omitempty"`
}

type MessageHeader7 struct {
	MsgId       common.Max35Text        `xml:"MsgId"`
	CreDtTm     *common.ISODateTime     `xml:"CreDtTm,omitempty" json:",omitempty"`
	ReqTp       *RequestType4Choice     `xml:"ReqTp,omitempty" json:",omitempty"`
	OrgnlBizQry *OriginalBusinessQuery1 `xml:"OrgnlBizQry,omitempty" json:",omitempty"`
	QryNm       *common.Max35Text       `xml:"QryNm,omitempty" json:",omitempty"`
}

type OriginalBusinessQuery1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	MsgNmId *common.Max35Text   `xml:"MsgNmId,omitempty" json:",omitempty"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

type ReturnBusinessDayInformationV07 struct {
	MsgHdr      MessageHeader7                  `xml:"MsgHdr"`
	RptOrErr    BusinessDayReportOrError9Choice `xml:"RptOrErr"`
	SplmtryData []SupplementaryData1            `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type SystemAvailabilityAndEvents3 struct {
	SysCcy  *common.ActiveCurrencyCode `xml:"SysCcy,omitempty" json:",omitempty"`
	SsnPrd  *TimePeriod1               `xml:"SsnPrd,omitempty" json:",omitempty"`
	Evt     []SystemEvent3             `xml:"Evt,omitempty" json:",omitempty"`
	ClsrInf []SystemClosure2           `xml:"ClsrInf,omitempty" json:",omitempty"`
}

type SystemClosure2 struct {
	Prd *DateTimePeriod1Choice `xml:"Prd,omitempty" json:",omitempty"`
	Rsn ClosureReason2Choice   `xml:"Rsn"`
}

type SystemEvent3 struct {
	Tp       SystemEventType4Choice `xml:"Tp"`
	SchdldTm common.ISODateTime     `xml:"SchdldTm"`
	FctvTm   *common.ISODateTime    `xml:"FctvTm,omitempty" json:",omitempty"`
	StartTm  *common.ISODateTime    `xml:"StartTm,omitempty" json:",omitempty"`
	EndTm    *common.ISODateTime    `xml:"EndTm,omitempty" json:",omitempty"`
}

type SystemEventType4Choice struct {
	Cd    ExternalSystemEventType1Code `xml:"Cd"`
	Prtry GenericIdentification1       `xml:"Prtry"`
}

type SystemStatus2Choice struct {
	Cd    SystemStatus2Code      `xml:"Cd"`
	Prtry GenericIdentification1 `xml:"Prtry"`
}

type SystemStatus3 struct {
	Sts     SystemStatus2Choice    `xml:"Sts"`
	VldtyTm *DateTimePeriod1Choice `xml:"VldtyTm,omitempty" json:",omitempty"`
}

type TimePeriod1 struct {
	FrTm common.ISOTime `xml:"FrTm"`
	ToTm common.ISOTime `xml:"ToTm"`
}

type BackupPaymentV07 struct {
	MsgHdr      MessageHeader1        `xml:"MsgHdr"`
	OrgnlMsgId  *MessageHeader1       `xml:"OrgnlMsgId,omitempty" json:",omitempty"`
	InstrInf    *PaymentInstruction13 `xml:"InstrInf,omitempty" json:",omitempty"`
	TrfdAmt     Amount2Choice         `xml:"TrfdAmt"`
	Cdtr        SystemMember3         `xml:"Cdtr"`
	CdtrAgt     *SystemMember3        `xml:"CdtrAgt,omitempty" json:",omitempty"`
	DbtrAgt     *SystemMember3        `xml:"DbtrAgt,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1  `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type MemberIdentification3Choice struct {
	BICFI       common.BICFIDec2014Identifier       `xml:"BICFI"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId"`
	Othr        GenericFinancialIdentification1     `xml:"Othr"`
}

type PaymentInstruction13 struct {
	ReqdExctnDtTm *common.ISODateTime `xml:"ReqdExctnDtTm,omitempty" json:",omitempty"`
	PmtTp         *PaymentType4Choice `xml:"PmtTp,omitempty" json:",omitempty"`
}

type PaymentType4Choice struct {
	Cd    PaymentType3Code `xml:"Cd"`
	Prtry common.Max35Text `xml:"Prtry"`
}

type SystemMember3 struct {
	SysId *SystemIdentification2Choice `xml:"SysId,omitempty" json:",omitempty"`
	MmbId MemberIdentification3Choice  `xml:"MmbId"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type AmendmentInformationDetails13 struct {
	OrgnlMndtId      *common.Max35Text                             `xml:"OrgnlMndtId,omitempty" json:",omitempty"`
	OrgnlCdtrSchmeId *PartyIdentification135                       `xml:"OrgnlCdtrSchmeId,omitempty" json:",omitempty"`
	OrgnlCdtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlCdtrAgt,omitempty" json:",omitempty"`
	OrgnlCdtrAgtAcct *CashAccount38                                `xml:"OrgnlCdtrAgtAcct,omitempty" json:",omitempty"`
	OrgnlDbtr        *PartyIdentification135                       `xml:"OrgnlDbtr,omitempty" json:",omitempty"`
	OrgnlDbtrAcct    *CashAccount38                                `xml:"OrgnlDbtrAcct,omitempty" json:",omitempty"`
	OrgnlDbtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlDbtrAgt,omitempty" json:",omitempty"`
	OrgnlDbtrAgtAcct *CashAccount38                                `xml:"OrgnlDbtrAgtAcct,omitempty" json:",omitempty"`
	OrgnlFnlColltnDt *common.ISODate                               `xml:"OrgnlFnlColltnDt,omitempty" json:",omitempty"`
	OrgnlFrqcy       *Frequency36Choice                            `xml:"OrgnlFrqcy,omitempty" json:",omitempty"`
	OrgnlRsn         *MandateSetupReason1Choice                    `xml:"OrgnlRsn,omitempty" json:",omitempty"`
	OrgnlTrckgDays   *Exact2NumericText                            `xml:"OrgnlTrckgDays,omitempty" json:",omitempty"`
}

type AmountType4Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"EqvtAmt"`
}

type Case5 struct {
	Id             common.Max35Text `xml:"Id"`
	Cretr          Party40Choice    `xml:"Cretr"`
	ReopCaseIndctn bool             `xml:"ReopCaseIndctn,omitempty" json:",omitempty"`
}

type CaseAssignment5 struct {
	Id      common.Max35Text   `xml:"Id"`
	Assgnr  Party40Choice      `xml:"Assgnr"`
	Assgne  Party40Choice      `xml:"Assgne"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice         `xml:"Id"`
	Tp   *CashAccountType2Choice              `xml:"Tp,omitempty" json:",omitempty"`
	Ccy  *common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
	Nm   *common.Max70Text                    `xml:"Nm,omitempty" json:",omitempty"`
	Prxy *ProxyAccountIdentification1         `xml:"Prxy,omitempty" json:",omitempty"`
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

type CreditTransferMandateData1 struct {
	MndtId       *common.Max35Text          `xml:"MndtId,omitempty" json:",omitempty"`
	Tp           *MandateTypeInformation2   `xml:"Tp,omitempty" json:",omitempty"`
	DtOfSgntr    *common.ISODate            `xml:"DtOfSgntr,omitempty" json:",omitempty"`
	DtOfVrfctn   *common.ISODateTime        `xml:"DtOfVrfctn,omitempty" json:",omitempty"`
	ElctrncSgntr *common.Max10KBinary       `xml:"ElctrncSgntr,omitempty" json:",omitempty"`
	FrstPmtDt    *common.ISODate            `xml:"FrstPmtDt,omitempty" json:",omitempty"`
	FnlPmtDt     *common.ISODate            `xml:"FnlPmtDt,omitempty" json:",omitempty"`
	Frqcy        *Frequency36Choice         `xml:"Frqcy,omitempty" json:",omitempty"`
	Rsn          *MandateSetupReason1Choice `xml:"Rsn,omitempty" json:",omitempty"`
}

type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `xml:"Tp,omitempty" json:",omitempty"`
	Ref *common.Max35Text       `xml:"Ref,omitempty" json:",omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"Cd"`
	Prtry common.Max35Text  `xml:"Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text            `xml:"Issr,omitempty" json:",omitempty"`
}

type DiscountAmountAndType1 struct {
	Tp  *DiscountAmountType1Choice        `xml:"Tp,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd *common.CreditDebitCode           `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Rsn       *common.Max4Text                  `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf  *common.Max140Text                `xml:"AddtlInf,omitempty" json:",omitempty"`
}

type DocumentLineIdentification1 struct {
	Tp     *DocumentLineType1 `xml:"Tp,omitempty" json:",omitempty"`
	Nb     *common.Max35Text  `xml:"Nb,omitempty" json:",omitempty"`
	RltdDt *common.ISODate    `xml:"RltdDt,omitempty" json:",omitempty"`
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"Id"`
	Desc *common.Max2048Text           `xml:"Desc,omitempty" json:",omitempty"`
	Amt  *RemittanceAmount3            `xml:"Amt,omitempty" json:",omitempty"`
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text       `xml:"Issr,omitempty" json:",omitempty"`
}

type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount    `xml:"Amt"`
	CcyOfTrf *common.ActiveOrHistoricCurrencyCode `xml:"CcyOfTrf"`
}

type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"Tp"`
	Prd    FrequencyPeriod1    `xml:"Prd"`
	PtInTm FrequencyAndMoment1 `xml:"PtInTm"`
}

type FrequencyAndMoment1 struct {
	Tp     Frequency6Code    `xml:"Tp"`
	PtInTm Exact2NumericText `xml:"PtInTm"`
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"Tp"`
	CntPerPrd float64        `xml:"CntPerPrd"`
}

type Garnishment3 struct {
	Tp                GarnishmentType1                   `xml:"Tp"`
	Grnshee           *PartyIdentification135            `xml:"Grnshee,omitempty" json:",omitempty"`
	GrnshmtAdmstr     *PartyIdentification135            `xml:"GrnshmtAdmstr,omitempty" json:",omitempty"`
	RefNb             *common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Dt                *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
	FmlyMdclInsrncInd bool                               `xml:"FmlyMdclInsrncInd,omitempty" json:",omitempty"`
	MplyeeTermntnInd  bool                               `xml:"MplyeeTermntnInd,omitempty" json:",omitempty"`
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text      `xml:"Issr,omitempty" json:",omitempty"`
}

type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type InstructionForAssignee1 struct {
	Cd       *ExternalAgentInstruction1Code `xml:"Cd,omitempty" json:",omitempty"`
	InstrInf *common.Max140Text             `xml:"InstrInf,omitempty" json:",omitempty"`
}

type InstructionForCreditorAgent3 struct {
	Cd       *ExternalCreditorAgentInstruction1Code `xml:"Cd,omitempty" json:",omitempty"`
	InstrInf *common.Max140Text                     `xml:"InstrInf,omitempty" json:",omitempty"`
}

type InstructionForNextAgent1 struct {
	Cd       *Instruction4Code  `xml:"Cd,omitempty" json:",omitempty"`
	InstrInf *common.Max140Text `xml:"InstrInf,omitempty" json:",omitempty"`
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type MandateClassification1Choice struct {
	Cd    common.MandateClassification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

type MandateRelatedData1Choice struct {
	DrctDbtMndt *MandateRelatedInformation14 `xml:"DrctDbtMndt,omitempty" json:",omitempty"`
	CdtTrfMndt  *CreditTransferMandateData1  `xml:"CdtTrfMndt,omitempty" json:",omitempty"`
}

type MandateRelatedInformation14 struct {
	MndtId        *common.Max35Text              `xml:"MndtId,omitempty" json:",omitempty"`
	DtOfSgntr     *common.ISODate                `xml:"DtOfSgntr,omitempty" json:",omitempty"`
	AmdmntInd     bool                           `xml:"AmdmntInd,omitempty" json:",omitempty"`
	AmdmntInfDtls *AmendmentInformationDetails13 `xml:"AmdmntInfDtls,omitempty" json:",omitempty"`
	ElctrncSgntr  *common.Max1025Text            `xml:"ElctrncSgntr,omitempty" json:",omitempty"`
	FrstColltnDt  *common.ISODate                `xml:"FrstColltnDt,omitempty" json:",omitempty"`
	FnlColltnDt   *common.ISODate                `xml:"FnlColltnDt,omitempty" json:",omitempty"`
	Frqcy         *Frequency36Choice             `xml:"Frqcy,omitempty" json:",omitempty"`
	Rsn           *MandateSetupReason1Choice     `xml:"Rsn,omitempty" json:",omitempty"`
	TrckgDays     *Exact2NumericText             `xml:"TrckgDays,omitempty" json:",omitempty"`
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"Cd"`
	Prtry common.Max70Text                `xml:"Prtry"`
}

type MandateTypeInformation2 struct {
	SvcLvl    *ServiceLevel8Choice          `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm *LocalInstrument2Choice       `xml:"LclInstrm,omitempty" json:",omitempty"`
	CtgyPurp  *CategoryPurpose1Choice       `xml:"CtgyPurp,omitempty" json:",omitempty"`
	Clssfctn  *MandateClassification1Choice `xml:"Clssfctn,omitempty" json:",omitempty"`
}

type OriginalGroupInformation29 struct {
	OrgnlMsgId   common.Max35Text    `xml:"OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text    `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm *common.ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
}

type OriginalTransactionReference31 struct {
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty" json:",omitempty"`
	Amt            *AmountType4Choice                            `xml:"Amt,omitempty" json:",omitempty"`
	IntrBkSttlmDt  *common.ISODate                               `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	ReqdColltnDt   *common.ISODate                               `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
	ReqdExctnDt    *DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	CdtrSchmeId    *PartyIdentification135                       `xml:"CdtrSchmeId,omitempty" json:",omitempty"`
	SttlmInf       *SettlementInstruction7                       `xml:"SttlmInf,omitempty" json:",omitempty"`
	PmtTpInf       *PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	PmtMtd         *PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:",omitempty"`
	MndtRltdInf    *MandateRelatedData1Choice                    `xml:"MndtRltdInf,omitempty" json:",omitempty"`
	RmtInf         *RemittanceInformation16                      `xml:"RmtInf,omitempty" json:",omitempty"`
	UltmtDbtr      *Party40Choice                                `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr           *Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct       *CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	DbtrAgtAcct    *CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:",omitempty"`
	CdtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	CdtrAgtAcct    *CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:",omitempty"`
	Cdtr           *Party40Choice                                `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct       *CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr      *Party40Choice                                `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	Purp           *Purpose2Choice                               `xml:"Purp,omitempty" json:",omitempty"`
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"Pty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

type PaymentTypeInformation27 struct {
	InstrPrty *Priority2Code          `xml:"InstrPrty,omitempty" json:",omitempty"`
	ClrChanl  *ClearingChannel2Code   `xml:"ClrChanl,omitempty" json:",omitempty"`
	SvcLvl    []ServiceLevel8Choice   `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:",omitempty"`
	SeqTp     *SequenceType3Code      `xml:"SeqTp,omitempty" json:",omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:",omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"Tp,omitempty" json:",omitempty"`
	Id common.Max2048Text       `xml:"Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text     `xml:"Prtry"`
}

type ReferredDocumentInformation7 struct {
	Tp       *ReferredDocumentType4     `xml:"Tp,omitempty" json:",omitempty"`
	Nb       *common.Max35Text          `xml:"Nb,omitempty" json:",omitempty"`
	RltdDt   *common.ISODate            `xml:"RltdDt,omitempty" json:",omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"LineDtls,omitempty" json:",omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"Cd"`
	Prtry common.Max35Text  `xml:"Prtry"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text           `xml:"Issr,omitempty" json:",omitempty"`
}

type RemittanceAmount2 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1           `xml:"DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1                `xml:"TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1              `xml:"AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
}

type RemittanceAmount3 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1           `xml:"DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1                `xml:"TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1              `xml:"AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
}

type RemittanceInformation16 struct {
	Ustrd []common.Max140Text                 `xml:"Ustrd,omitempty" json:",omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"Strd,omitempty" json:",omitempty"`
}

type RequestToModifyPaymentV07 struct {
	Assgnmt        CaseAssignment5              `xml:"Assgnmt"`
	Case           *Case5                       `xml:"Case,omitempty" json:",omitempty"`
	Undrlyg        UnderlyingTransaction6Choice `xml:"Undrlyg"`
	Mod            RequestedModification9       `xml:"Mod"`
	InstrForAssgne *InstructionForAssignee1     `xml:"InstrForAssgne,omitempty" json:",omitempty"`
	SplmtryData    []SupplementaryData1         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type RequestedModification9 struct {
	InstrId         *common.Max35Text                  `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId      *common.Max35Text                  `xml:"EndToEndId,omitempty" json:",omitempty"`
	TxId            *common.Max35Text                  `xml:"TxId,omitempty" json:",omitempty"`
	PmtTpInf        *PaymentTypeInformation27          `xml:"PmtTpInf,omitempty" json:",omitempty"`
	ReqdExctnDt     *DateAndDateTime2Choice            `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	ReqdColltnDt    *common.ISODate                    `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
	IntrBkSttlmDt   *common.ISODate                    `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	SttlmTmIndctn   *SettlementDateTimeIndication1     `xml:"SttlmTmIndctn,omitempty" json:",omitempty"`
	Amt             *AmountType4Choice                 `xml:"Amt,omitempty" json:",omitempty"`
	IntrBkSttlmAmt  *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty" json:",omitempty"`
	ChrgBr          *ChargeBearerType1Code             `xml:"ChrgBr,omitempty" json:",omitempty"`
	UltmtDbtr       *PartyIdentification135            `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr            *PartyIdentification135            `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct        *CashAccount38                     `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgtAcct     *CashAccount38                     `xml:"DbtrAgtAcct,omitempty" json:",omitempty"`
	SttlmInf        *SettlementInstruction6            `xml:"SttlmInf,omitempty" json:",omitempty"`
	CdtrAgtAcct     *CashAccount38                     `xml:"CdtrAgtAcct,omitempty" json:",omitempty"`
	Cdtr            *PartyIdentification135            `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct        *CashAccount38                     `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr       *PartyIdentification135            `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	Purp            *Purpose2Choice                    `xml:"Purp,omitempty" json:",omitempty"`
	InstrForDbtrAgt *common.Max140Text                 `xml:"InstrForDbtrAgt,omitempty" json:",omitempty"`
	InstrForNxtAgt  []InstructionForNextAgent1         `xml:"InstrForNxtAgt,omitempty" json:",omitempty"`
	InstrForCdtrAgt []InstructionForCreditorAgent3     `xml:"InstrForCdtrAgt,omitempty" json:",omitempty"`
	RmtInf          *RemittanceInformation16           `xml:"RmtInf,omitempty" json:",omitempty"`
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"Cd"`
	Prtry common.Max35Text          `xml:"Prtry"`
}

type SettlementDateTimeIndication1 struct {
	DbtDtTm *common.ISODateTime `xml:"DbtDtTm,omitempty" json:",omitempty"`
	CdtDtTm *common.ISODateTime `xml:"CdtDtTm,omitempty" json:",omitempty"`
}

type SettlementInstruction6 struct {
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt,omitempty" json:",omitempty"`
	InstgRmbrsmntAgtAcct *CashAccount38                                `xml:"InstgRmbrsmntAgtAcct,omitempty" json:",omitempty"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt,omitempty" json:",omitempty"`
	InstdRmbrsmntAgtAcct *CashAccount38                                `xml:"InstdRmbrsmntAgtAcct,omitempty" json:",omitempty"`
}

type SettlementInstruction7 struct {
	SttlmMtd             SettlementMethod1Code                         `xml:"SttlmMtd"`
	SttlmAcct            *CashAccount38                                `xml:"SttlmAcct,omitempty" json:",omitempty"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"ClrSys,omitempty" json:",omitempty"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt,omitempty" json:",omitempty"`
	InstgRmbrsmntAgtAcct *CashAccount38                                `xml:"InstgRmbrsmntAgtAcct,omitempty" json:",omitempty"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt,omitempty" json:",omitempty"`
	InstdRmbrsmntAgtAcct *CashAccount38                                `xml:"InstdRmbrsmntAgtAcct,omitempty" json:",omitempty"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification6 `xml:"ThrdRmbrsmntAgt,omitempty" json:",omitempty"`
	ThrdRmbrsmntAgtAcct  *CashAccount38                                `xml:"ThrdRmbrsmntAgtAcct,omitempty" json:",omitempty"`
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty" json:",omitempty"`
	RfrdDocAmt  *RemittanceAmount2             `xml:"RfrdDocAmt,omitempty" json:",omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty" json:",omitempty"`
	Invcr       *PartyIdentification135        `xml:"Invcr,omitempty" json:",omitempty"`
	Invcee      *PartyIdentification135        `xml:"Invcee,omitempty" json:",omitempty"`
	TaxRmt      *TaxInformation7               `xml:"TaxRmt,omitempty" json:",omitempty"`
	GrnshmtRmt  *Garnishment3                  `xml:"GrnshmtRmt,omitempty" json:",omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"AddtlRmtInf,omitempty" json:",omitempty"`
}

type TaxAmount2 struct {
	Rate         float64                            `xml:"Rate,omitempty" json:",omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty" json:",omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:",omitempty"`
	Dtls         []TaxRecordDetails2                `xml:"Dtls,omitempty" json:",omitempty"`
}

type TaxAmountAndType1 struct {
	Tp  *TaxAmountType1Choice             `xml:"Tp,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"Cd"`
	Prtry common.Max35Text           `xml:"Prtry"`
}

type TaxAuthorisation1 struct {
	Titl *common.Max35Text  `xml:"Titl,omitempty" json:",omitempty"`
	Nm   *common.Max140Text `xml:"Nm,omitempty" json:",omitempty"`
}

type TaxInformation7 struct {
	Cdtr            *TaxParty1                         `xml:"Cdtr,omitempty" json:",omitempty"`
	Dbtr            *TaxParty2                         `xml:"Dbtr,omitempty" json:",omitempty"`
	UltmtDbtr       *TaxParty2                         `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	AdmstnZone      *common.Max35Text                  `xml:"AdmstnZone,omitempty" json:",omitempty"`
	RefNb           *common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Mtd             *common.Max35Text                  `xml:"Mtd,omitempty" json:",omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:",omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:",omitempty"`
	Dt              *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	SeqNb           float64                            `xml:"SeqNb,omitempty" json:",omitempty"`
	Rcrd            []TaxRecord2                       `xml:"Rcrd,omitempty" json:",omitempty"`
}

type TaxParty1 struct {
	TaxId  *common.Max35Text `xml:"TaxId,omitempty" json:",omitempty"`
	RegnId *common.Max35Text `xml:"RegnId,omitempty" json:",omitempty"`
	TaxTp  *common.Max35Text `xml:"TaxTp,omitempty" json:",omitempty"`
}

type TaxParty2 struct {
	TaxId   *common.Max35Text  `xml:"TaxId,omitempty" json:",omitempty"`
	RegnId  *common.Max35Text  `xml:"RegnId,omitempty" json:",omitempty"`
	TaxTp   *common.Max35Text  `xml:"TaxTp,omitempty" json:",omitempty"`
	Authstn *TaxAuthorisation1 `xml:"Authstn,omitempty" json:",omitempty"`
}

type TaxPeriod2 struct {
	Yr     *common.ISODate       `xml:"Yr,omitempty" json:",omitempty"`
	Tp     *TaxRecordPeriod1Code `xml:"Tp,omitempty" json:",omitempty"`
	FrToDt *DatePeriod2          `xml:"FrToDt,omitempty" json:",omitempty"`
}

type TaxRecord2 struct {
	Tp       *common.Max35Text  `xml:"Tp,omitempty" json:",omitempty"`
	Ctgy     *common.Max35Text  `xml:"Ctgy,omitempty" json:",omitempty"`
	CtgyDtls *common.Max35Text  `xml:"CtgyDtls,omitempty" json:",omitempty"`
	DbtrSts  *common.Max35Text  `xml:"DbtrSts,omitempty" json:",omitempty"`
	CertId   *common.Max35Text  `xml:"CertId,omitempty" json:",omitempty"`
	FrmsCd   *common.Max35Text  `xml:"FrmsCd,omitempty" json:",omitempty"`
	Prd      *TaxPeriod2        `xml:"Prd,omitempty" json:",omitempty"`
	TaxAmt   *TaxAmount2        `xml:"TaxAmt,omitempty" json:",omitempty"`
	AddtlInf *common.Max140Text `xml:"AddtlInf,omitempty" json:",omitempty"`
}

type TaxRecordDetails2 struct {
	Prd *TaxPeriod2                       `xml:"Prd,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type UnderlyingGroupInformation1 struct {
	OrgnlMsgId         common.Max35Text    `xml:"OrgnlMsgId"`
	OrgnlMsgNmId       common.Max35Text    `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm       *common.ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	OrgnlMsgDlvryChanl *common.Max35Text   `xml:"OrgnlMsgDlvryChanl,omitempty" json:",omitempty"`
}

type UnderlyingPaymentInstruction6 struct {
	OrgnlGrpInf     *UnderlyingGroupInformation1      `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlPmtInfId   *common.Max35Text                 `xml:"OrgnlPmtInfId,omitempty" json:",omitempty"`
	OrgnlInstrId    *common.Max35Text                 `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId *common.Max35Text                 `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlUETR       *common.UUIDv4Identifier          `xml:"OrgnlUETR,omitempty" json:",omitempty"`
	OrgnlInstdAmt   ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt"`
	ReqdExctnDt     *DateAndDateTime2Choice           `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	ReqdColltnDt    *common.ISODate                   `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
	OrgnlTxRef      *OriginalTransactionReference31   `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
}

type UnderlyingPaymentTransaction5 struct {
	OrgnlGrpInf         *UnderlyingGroupInformation1      `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlInstrId        *common.Max35Text                 `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId     *common.Max35Text                 `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlTxId           *common.Max35Text                 `xml:"OrgnlTxId,omitempty" json:",omitempty"`
	OrgnlUETR           *common.UUIDv4Identifier          `xml:"OrgnlUETR,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt"`
	OrgnlIntrBkSttlmDt  common.ISODate                    `xml:"OrgnlIntrBkSttlmDt"`
	OrgnlTxRef          *OriginalTransactionReference31   `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
}

type UnderlyingStatementEntry3 struct {
	OrgnlGrpInf *OriginalGroupInformation29 `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlStmtId *common.Max35Text           `xml:"OrgnlStmtId,omitempty" json:",omitempty"`
	OrgnlNtryId *common.Max35Text           `xml:"OrgnlNtryId,omitempty" json:",omitempty"`
	OrgnlUETR   *common.UUIDv4Identifier    `xml:"OrgnlUETR,omitempty" json:",omitempty"`
}

type UnderlyingTransaction6Choice struct {
	Initn    UnderlyingPaymentInstruction6 `xml:"Initn"`
	IntrBk   UnderlyingPaymentTransaction5 `xml:"IntrBk"`
	StmtNtry UnderlyingStatementEntry3     `xml:"StmtNtry"`
}
