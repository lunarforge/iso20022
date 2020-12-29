// Code generated by xsdgen. DO NOT EDIT.

package camt_v08

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Prtry"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 PstlAdr,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 MmbId"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateTimePeriod1 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 ToDtTm"`
}

type DateTimePeriod1Choice struct {
	FrDtTm ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 FrDtTm"`
	ToDtTm ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 ToDtTm"`
	DtTmRg DateTimePeriod1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 DtTmRg"`
}

type Document struct {
	ModfyTx ModifyTransactionV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 ModfyTx"`
}

// Must match the pattern [BEOVW]{1,1}[0-9]{2,2}|DUM
type EntryTypeIdentifier string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Othr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 SchmeNm,omitempty"`
}

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

// May be one of PBEN, TTIL, TFRO
type Instruction1Code string

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type LongPaymentIdentification2 struct {
	TxId           Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 TxId,omitempty"`
	UETR           UUIDv4Identifier                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 UETR,omitempty"`
	IntrBkSttlmAmt float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 IntrBkSttlmAmt"`
	IntrBkSttlmDt  ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 IntrBkSttlmDt"`
	PmtMtd         PaymentOrigin1Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 PmtMtd,omitempty"`
	InstgAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 InstgAgt"`
	InstdAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 InstdAgt"`
	NtryTp         EntryTypeIdentifier                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 NtryTp,omitempty"`
	EndToEndId     Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 EndToEndId,omitempty"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must be at least 1 items long
type Max70Text string

type MessageHeader1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 MsgId"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 CreDtTm,omitempty"`
}

type ModifyTransactionV08 struct {
	MsgHdr      MessageHeader1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 MsgHdr"`
	Mod         []TransactionModification5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Mod"`
	SplmtryData []SupplementaryData1       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 SplmtryData,omitempty"`
}

type PaymentIdentification6Choice struct {
	TxId      Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 TxId"`
	QId       QueueTransactionIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 QId"`
	LngBizId  LongPaymentIdentification2      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 LngBizId"`
	ShrtBizId ShortPaymentIdentification2     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 ShrtBizId"`
	PrtryId   Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 PrtryId"`
}

type PaymentInstruction33 struct {
	Instr       Instruction1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Instr,omitempty"`
	Tp          PaymentType4Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Tp,omitempty"`
	Prty        Priority1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Prty,omitempty"`
	PrcgVldtyTm DateTimePeriod1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 PrcgVldtyTm,omitempty"`
}

// May be one of BDT, BCT, CDT, CCT, CHK, BKT, DCP, CCP, RTI, CAN
type PaymentInstrument1Code string

type PaymentOrigin1Choice struct {
	FINMT    Max3NumericText        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 FINMT"`
	XMLMsgNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 XMLMsgNm"`
	Prtry    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Prtry"`
	Instrm   PaymentInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Instrm"`
}

// May be one of CBS, BCK, BAL, CLS, CTR, CBH, CBP, DPG, DPN, EXP, TCH, LMT, LIQ, DPP, DPH, DPS, STF, TRP, TCS, LOA, LOR, TCP, OND, MGL
type PaymentType3Code string

type PaymentType4Choice struct {
	Cd    PaymentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Cd"`
	Prtry Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Prtry"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 AdrLine,omitempty"`
}

type Priority1Choice struct {
	Cd    Priority5Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Cd"`
	Prtry Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Prtry"`
}

// May be one of HIGH, LOWW, NORM, URGT
type Priority5Code string

type QueueTransactionIdentification1 struct {
	QId    Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 QId"`
	PosInQ Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 PosInQ"`
}

type ShortPaymentIdentification2 struct {
	TxId          Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 TxId"`
	IntrBkSttlmDt ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 IntrBkSttlmDt"`
	InstgAgt      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 InstgAgt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TransactionModification5 struct {
	PmtId        PaymentIdentification6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 PmtId"`
	NewPmtValSet PaymentInstruction33         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 NewPmtValSet"`
}

// Must match the pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}
type UUIDv4Identifier string

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
