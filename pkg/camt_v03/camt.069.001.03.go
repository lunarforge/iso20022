// Code generated by xsdgen. DO NOT EDIT.

package camt_v03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Prtry"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Prtry"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 PstlAdr,omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Id"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 MmbId"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DatePeriod2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 ToDt"`
}

type DatePeriod2Choice struct {
	FrDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 FrDt"`
	ToDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 ToDt"`
	FrToDt DatePeriod2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 FrToDt"`
}

type Document struct {
	GetStgOrdr GetStandingOrderV03 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 GetStgOrdr"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Issr,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 SchmeNm,omitempty"`
}

type GetStandingOrderV03 struct {
	MsgHdr        MessageHeader4       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 MsgHdr"`
	StgOrdrQryDef StandingOrderQuery3  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 StgOrdrQryDef,omitempty"`
	SplmtryData   []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 SplmtryData,omitempty"`
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

type MessageHeader4 struct {
	MsgId   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 MsgId"`
	CreDtTm ISODateTime        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 CreDtTm,omitempty"`
	ReqTp   RequestType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 ReqTp,omitempty"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 AdrLine,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Cd"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Prtry"`
}

// May be one of ALLL, CHNG, MODF, DELD
type QueryType2Code string

type RequestType3Choice struct {
	Cd    StandingOrderQueryType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Cd"`
	Prtry GenericIdentification1      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Prtry"`
}

type StandingOrderCriteria3 struct {
	NewQryNm Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 NewQryNm,omitempty"`
	SchCrit  []StandingOrderSearchCriteria3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 SchCrit,omitempty"`
	RtrCrit  StandingOrderReturnCriteria1   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 RtrCrit,omitempty"`
}

type StandingOrderCriteria3Choice struct {
	QryNm   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 QryNm"`
	NewCrit StandingOrderCriteria3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 NewCrit"`
}

type StandingOrderQuery3 struct {
	QryTp       QueryType2Code               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 QryTp,omitempty"`
	StgOrdrCrit StandingOrderCriteria3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 StgOrdrCrit,omitempty"`
}

// May be one of SLST, SDTL, TAPS, SLSL, SWLS
type StandingOrderQueryType1Code string

type StandingOrderReturnCriteria1 struct {
	StgOrdrIdInd    bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 StgOrdrIdInd,omitempty"`
	TpInd           bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 TpInd,omitempty"`
	SysMmbInd       bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 SysMmbInd,omitempty"`
	RspnsblPtyInd   bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 RspnsblPtyInd,omitempty"`
	CcyInd          bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 CcyInd,omitempty"`
	DbtrAcctInd     bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 DbtrAcctInd,omitempty"`
	CdtrAcctInd     bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 CdtrAcctInd,omitempty"`
	AssoctdPoolAcct bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 AssoctdPoolAcct,omitempty"`
	FrqcyInd        bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 FrqcyInd,omitempty"`
	ExctnTpInd      bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 ExctnTpInd,omitempty"`
	VldtyFrInd      bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 VldtyFrInd,omitempty"`
	VldToInd        bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 VldToInd,omitempty"`
	LkSetIdInd      bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 LkSetIdInd,omitempty"`
	LkSetOrdrIdInd  bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 LkSetOrdrIdInd,omitempty"`
	LkSetOrdrSeqInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 LkSetOrdrSeqInd,omitempty"`
	TtlAmtInd       bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 TtlAmtInd,omitempty"`
	ZeroSweepInd    bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 ZeroSweepInd,omitempty"`
}

type StandingOrderSearchCriteria3 struct {
	KeyAttrbtsInd   bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 KeyAttrbtsInd,omitempty"`
	StgOrdrId       Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 StgOrdrId,omitempty"`
	Tp              StandingOrderType1Choice                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Tp,omitempty"`
	Acct            CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Acct,omitempty"`
	Ccy             ActiveCurrencyCode                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Ccy,omitempty"`
	VldtyPrd        DatePeriod2Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 VldtyPrd,omitempty"`
	SysMmb          BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 SysMmb,omitempty"`
	RspnsblPty      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 RspnsblPty,omitempty"`
	AssoctdPoolAcct AccountIdentification4Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 AssoctdPoolAcct,omitempty"`
	LkSetId         Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 LkSetId,omitempty"`
	LkSetOrdrId     Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 LkSetOrdrId,omitempty"`
	LkSetOrdrSeq    float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 LkSetOrdrSeq,omitempty"`
	ZeroSweepInd    bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 ZeroSweepInd,omitempty"`
}

type StandingOrderType1Choice struct {
	Cd    StandingOrderType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Cd"`
	Prtry GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Prtry"`
}

// May be one of USTO, PSTO
type StandingOrderType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 Envlp"`
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
