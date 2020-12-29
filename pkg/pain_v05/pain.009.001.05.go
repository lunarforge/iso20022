// Code generated by xsdgen. DO NOT EDIT.

package pain_v05

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type AuthenticationChannel1Choice struct {
	Cd    ExternalAuthenticationChannel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type Authorisation1Choice struct {
	Cd    Authorisation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry Max128Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

// May be one of AUTH, FDET, FSUM, ILEV
type Authorisation1Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICFIIdentifier string

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Id,omitempty"`
	Nm      Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Nm,omitempty"`
	PstlAdr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PstlAdr,omitempty"`
}

type CashAccount24 struct {
	Id  AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Id"`
	Tp  CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Tp,omitempty"`
	Ccy ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Ccy,omitempty"`
	Nm  Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Nm,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MmbId"`
}

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Othr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CtryOfBirth"`
}

type DatePeriodDetails1 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 ToDt,omitempty"`
}

type Document struct {
	MndtInitnReq MandateInitiationRequestV05 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MndtInitnReq"`
}

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT, PUOR
type DocumentType6Code string

// Must match the pattern [0-9]{2}
type Exact2NumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalAuthenticationChannel1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalCategoryPurpose1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalLocalInstrument1Code string

// Must be at least 1 items long
type ExternalMandateSetupReason1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// Must be at least 1 items long
type ExternalServiceLevel1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type FinancialInstitutionIdentification8 struct {
	BICFI       BICFIIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 ClrSysMmbId,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Othr,omitempty"`
}

// May be one of NEVR, YEAR, RATE, MIAN, QURT
type Frequency10Code string

type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Tp"`
	Prd    FrequencyPeriod1    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prd"`
	PtInTm FrequencyAndMoment1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PtInTm"`
}

type Frequency37Choice struct {
	Cd    Frequency10Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, FRTN
type Frequency6Code string

type FrequencyAndMoment1 struct {
	Tp     Frequency6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Tp"`
	PtInTm Exact2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PtInTm"`
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Tp"`
	CntPerPrd float64        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CntPerPrd"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Issr,omitempty"`
}

type GroupHeader47 struct {
	MsgId    Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MsgId"`
	CreDtTm  ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CreDtTm"`
	Authstn  []Authorisation1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Authstn,omitempty"`
	InitgPty PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 InitgPty,omitempty"`
	InstgAgt BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 InstgAgt,omitempty"`
	InstdAgt BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 InstdAgt,omitempty"`
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

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type Mandate10 struct {
	MndtId        []Max35Text                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MndtId,omitempty"`
	MndtReqId     Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MndtReqId"`
	Authntcn      MandateAuthentication1                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Authntcn,omitempty"`
	Tp            MandateTypeInformation2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Tp,omitempty"`
	Ocrncs        MandateOccurrences4                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Ocrncs,omitempty"`
	TrckgInd      bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 TrckgInd"`
	FrstColltnAmt ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 FrstColltnAmt,omitempty"`
	ColltnAmt     ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 ColltnAmt,omitempty"`
	MaxAmt        ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MaxAmt,omitempty"`
	Adjstmnt      MandateAdjustment1                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Adjstmnt,omitempty"`
	Rsn           MandateSetupReason1Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Rsn,omitempty"`
	CdtrSchmeId   PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CdtrSchmeId,omitempty"`
	Cdtr          PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cdtr"`
	CdtrAcct      CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CdtrAcct,omitempty"`
	CdtrAgt       BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CdtrAgt,omitempty"`
	UltmtCdtr     PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 UltmtCdtr,omitempty"`
	Dbtr          PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Dbtr"`
	DbtrAcct      CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 DbtrAcct,omitempty"`
	DbtrAgt       BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 DbtrAgt"`
	UltmtDbtr     PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 UltmtDbtr,omitempty"`
	MndtRef       Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MndtRef,omitempty"`
	RfrdDoc       []ReferredMandateDocument1                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 RfrdDoc,omitempty"`
	SplmtryData   []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SplmtryData,omitempty"`
}

type MandateAdjustment1 struct {
	DtAdjstmntRuleInd bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 DtAdjstmntRuleInd"`
	Ctgy              Frequency37Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Ctgy,omitempty"`
	Amt               ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Amt,omitempty"`
	Rate              float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Rate,omitempty"`
}

type MandateAuthentication1 struct {
	MsgAuthntcnCd Max16Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MsgAuthntcnCd,omitempty"`
	Dt            ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Dt,omitempty"`
	Chanl         AuthenticationChannel1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Chanl,omitempty"`
}

type MandateClassification1Choice struct {
	Cd    MandateClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

// May be one of FIXE, USGB, VARI
type MandateClassification1Code string

type MandateInitiationRequestV05 struct {
	GrpHdr      GroupHeader47        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 GrpHdr"`
	Mndt        []Mandate10          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Mndt"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SplmtryData,omitempty"`
}

type MandateOccurrences4 struct {
	SeqTp        SequenceType2Code  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SeqTp"`
	Frqcy        Frequency36Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Frqcy,omitempty"`
	Drtn         DatePeriodDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Drtn,omitempty"`
	FrstColltnDt ISODate            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 FrstColltnDt,omitempty"`
	FnlColltnDt  ISODate            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 FnlColltnDt,omitempty"`
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type MandateTypeInformation2 struct {
	SvcLvl    ServiceLevel8Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CtgyPurp,omitempty"`
	Clssfctn  MandateClassification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Clssfctn,omitempty"`
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
type Max70Text string

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type OrganisationIdentification8 struct {
	AnyBIC AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 AnyBIC,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type Party11Choice struct {
	OrgId  OrganisationIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PrvtId"`
}

type PartyIdentification43 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Nm,omitempty"`
	PstlAdr   PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PstlAdr,omitempty"`
	Id        Party11Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CtctDtls,omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 AdrLine,omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CdOrPrtry"`
	Issr      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Issr,omitempty"`
}

type ReferredMandateDocument1 struct {
	Tp      ReferredDocumentType4 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Tp,omitempty"`
	Nb      Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Nb,omitempty"`
	CdtrRef Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CdtrRef,omitempty"`
	RltdDt  ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 RltdDt,omitempty"`
}

// May be one of RCUR, OOFF
type SequenceType2Code string

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
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
