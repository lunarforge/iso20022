// Code generated by xsdgen. DO NOT EDIT.

package camt_v08

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Othr"`
}

type AccountOrBusinessError4Choice struct {
	Acct   CashAccount37    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Acct"`
	BizErr []ErrorHandling5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 BizErr"`
}

type AccountOrOperationalError4Choice struct {
	AcctRpt []AccountReport24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 AcctRpt"`
	OprlErr []ErrorHandling5  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 OprlErr"`
}

type AccountReport24 struct {
	AcctId    AccountIdentification4Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 AcctId"`
	AcctOrErr AccountOrBusinessError4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 AcctOrErr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prtry"`
}

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 AmtWthtCcy"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 AmtWthCcy"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BalanceRestrictionType1 struct {
	Tp     GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Tp"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Desc,omitempty"`
	PrcgTp ProcessingType1Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 PrcgTp,omitempty"`
}

// May be one of PDNG, STLD
type BalanceStatus1Code string

type BalanceType11Choice struct {
	Cd    ExternalSystemBalanceType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Cd"`
	Prtry Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prtry"`
}

type BalanceType9Choice struct {
	Cd    SystemBalanceType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Cd"`
	Prtry Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prtry"`
}

type BilateralLimit3 struct {
	CtrPtyId  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CtrPtyId"`
	LmtAmt    Amount2Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 LmtAmt"`
	CdtDbtInd CreditDebitCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CdtDbtInd"`
	BilBal    []CashBalance11                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 BilBal,omitempty"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 PstlAdr,omitempty"`
}

type CashAccount37 struct {
	Nm        Max70Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Nm,omitempty"`
	Tp        CashAccountType2Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Tp,omitempty"`
	Ccy       ActiveOrHistoricCurrencyCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Ccy,omitempty"`
	Prxy      ProxyAccountIdentification1                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prxy,omitempty"`
	CurMulLmt Limit5                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CurMulLmt,omitempty"`
	Ownr      PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Ownr,omitempty"`
	Svcr      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Svcr,omitempty"`
	MulBal    []CashBalance13                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 MulBal,omitempty"`
	CurBilLmt []BilateralLimit3                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CurBilLmt,omitempty"`
	StgOrdr   []StandingOrder6                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 StgOrdr,omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Id"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prtry"`
}

type CashBalance11 struct {
	Amt       float64                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Amt"`
	CdtDbtInd CreditDebitCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CdtDbtInd"`
	Tp        BalanceType9Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Tp,omitempty"`
	Sts       BalanceStatus1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Sts,omitempty"`
	ValDt     DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 ValDt,omitempty"`
	NbOfPmts  float64                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 NbOfPmts,omitempty"`
}

type CashBalance13 struct {
	Amt       float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Amt"`
	CdtDbtInd CreditDebitCode         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CdtDbtInd"`
	Tp        BalanceType11Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Tp,omitempty"`
	Sts       BalanceStatus1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Sts,omitempty"`
	ValDt     DateAndDateTime2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 ValDt,omitempty"`
	PrcgDt    DateAndDateTime2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 PrcgDt,omitempty"`
	NbOfPmts  float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 NbOfPmts,omitempty"`
	RstrctnTp BalanceRestrictionType1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 RstrctnTp,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 MmbId"`
}

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 PrefrdMtd,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 DtTm"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CtryOfBirth"`
}

type DatePeriodDetails1 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 ToDt,omitempty"`
}

type Document struct {
	RtrAcct ReturnAccountV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 RtrAcct"`
}

type ErrorHandling3Choice struct {
	Cd    ExternalSystemErrorHandling1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Cd"`
	Prtry Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prtry"`
}

type ErrorHandling5 struct {
	Err  ErrorHandling3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Err"`
	Desc Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Desc,omitempty"`
}

type EventType1Choice struct {
	Cd    ExternalSystemEventType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prtry"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type ExecutionType1Choice struct {
	Tm  ISOTime          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Tm"`
	Evt EventType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Evt"`
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalEnquiryRequestType1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPaymentControlRequestType1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

// Must be at least 1 items long
type ExternalSystemBalanceType1Code string

// Must be at least 1 items long
type ExternalSystemErrorHandling1Code string

// Must be at least 1 items long
type ExternalSystemEventType1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Othr,omitempty"`
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, OVNG
type Frequency2Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Issr,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Issr,omitempty"`
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

type ISOTime time.Time

func (t *ISOTime) UnmarshalText(text []byte) error {
	return (*xsdTime)(t).UnmarshalText(text)
}
func (t ISOTime) MarshalText() ([]byte, error) {
	return xsdTime(t).MarshalText()
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type Limit5 struct {
	Amt       Amount2Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Amt"`
	CdtDbtInd CreditDebitCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CdtDbtInd"`
}

// Must be at least 1 items long
type Max128Text string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max4Text string

// Must be at least 1 items long
type Max70Text string

type MessageHeader7 struct {
	MsgId       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 MsgId"`
	CreDtTm     ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CreDtTm,omitempty"`
	ReqTp       RequestType4Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 ReqTp,omitempty"`
	OrgnlBizQry OriginalBusinessQuery1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 OrgnlBizQry,omitempty"`
	QryNm       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 QryNm,omitempty"`
}

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prtry"`
}

type OriginalBusinessQuery1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 MsgId"`
	MsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 MsgNmId,omitempty"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CreDtTm,omitempty"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 OrgId"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 PrvtId"`
}

type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CtctDtls,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

type ProcessingType1Choice struct {
	Cd    ProcessingType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Cd"`
	Prtry Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prtry"`
}

// May be one of RJCT, CVHD, RSVT, BLCK, EARM, EFAC, DLVR, COLD, CSDB
type ProcessingType1Code string

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Cd"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prtry"`
}

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 PmtCtrl"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Enqry"`
	Prtry   GenericIdentification1                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prtry"`
}

type ReturnAccountV08 struct {
	MsgHdr      MessageHeader7                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 MsgHdr"`
	RptOrErr    AccountOrOperationalError4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 RptOrErr"`
	SplmtryData []SupplementaryData1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 SplmtryData,omitempty"`
}

type StandingOrder6 struct {
	Amt             Amount2Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Amt"`
	CdtDbtInd       CreditDebitCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CdtDbtInd"`
	Ccy             ActiveCurrencyCode                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Ccy,omitempty"`
	Tp              StandingOrderType1Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Tp,omitempty"`
	AssoctdPoolAcct AccountIdentification4Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 AssoctdPoolAcct,omitempty"`
	Ref             Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Ref,omitempty"`
	Frqcy           Frequency2Code                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Frqcy,omitempty"`
	VldtyPrd        DatePeriodDetails1                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 VldtyPrd,omitempty"`
	SysMmb          BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 SysMmb,omitempty"`
	RspnsblPty      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 RspnsblPty,omitempty"`
	LkSetId         Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 LkSetId,omitempty"`
	LkSetOrdrId     Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 LkSetOrdrId,omitempty"`
	LkSetOrdrSeq    float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 LkSetOrdrSeq,omitempty"`
	ExctnTp         ExecutionType1Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 ExctnTp,omitempty"`
	Cdtr            BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Cdtr,omitempty"`
	CdtrAcct        CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CdtrAcct,omitempty"`
	Dbtr            BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Dbtr,omitempty"`
	DbtrAcct        CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 DbtrAcct,omitempty"`
	TtlsPerStgOrdr  StandingOrderTotalAmount1                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 TtlsPerStgOrdr,omitempty"`
	ZeroSweepInd    bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 ZeroSweepInd,omitempty"`
}

type StandingOrderTotalAmount1 struct {
	SetPrdfndOrdr TotalAmountAndCurrency1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 SetPrdfndOrdr"`
	PdgPrdfndOrdr TotalAmountAndCurrency1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 PdgPrdfndOrdr"`
	SetStgOrdr    TotalAmountAndCurrency1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 SetStgOrdr"`
	PdgStgOrdr    TotalAmountAndCurrency1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 PdgStgOrdr"`
}

type StandingOrderType1Choice struct {
	Cd    StandingOrderType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Cd"`
	Prtry GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Prtry"`
}

// May be one of USTO, PSTO
type StandingOrderType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of OPNG, INTM, CLSG, BOOK, CRRT, PDNG, LRLD, AVLB, LTSF, CRDT, EAST, PYMT, BLCK, XPCD, DLOD, XCRD, XDBT, ADJT, PRAV, DBIT, THRE, NOTE, FSET, BLOC, OTHB, CUST, FORC, COLC, FUND, PIPO, XCHG, CCPS, TOHB, COHB, DOHB, TPBL, CPBL, DPBL, FUTB, REJB, FCOL, FCOU, SCOL, SCOU, CUSA, XCHC, XCHN, DSET, LACK, NSET, OTCC, OTCG, OTCN, SAPD, SAPC, REPD, REPC, BSCD, BSCC, SAPP, IRLT, IRDR, DWRD, ADWR, AIDR
type SystemBalanceType2Code string

type TotalAmountAndCurrency1 struct {
	TtlAmt    float64            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 TtlAmt"`
	CdtDbtInd CreditDebitCode    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 CdtDbtInd,omitempty"`
	Ccy       ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 Ccy,omitempty"`
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

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("15:04:05.999999999")), nil
}
func (t xsdTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
