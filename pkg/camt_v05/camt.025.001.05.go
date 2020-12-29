// Code generated by xsdgen. DO NOT EDIT.

package camt_v05

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Prtry"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 PstlAdr,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 MmbId"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	Rct ReceiptV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Rct"`
}

// Must match the pattern [BEOVW]{1,1}[0-9]{2,2}|DUM
type EntryTypeIdentifier string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalEnquiryRequestType1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalPaymentControlRequestType1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Othr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Issr,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 SchmeNm,omitempty"`
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

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type LongPaymentIdentification2 struct {
	TxId           Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 TxId,omitempty"`
	UETR           UUIDv4Identifier                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 UETR,omitempty"`
	IntrBkSttlmAmt float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 IntrBkSttlmAmt"`
	IntrBkSttlmDt  ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 IntrBkSttlmDt"`
	PmtMtd         PaymentOrigin1Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 PmtMtd,omitempty"`
	InstgAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 InstgAgt"`
	InstdAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 InstdAgt"`
	NtryTp         EntryTypeIdentifier                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 NtryTp,omitempty"`
	EndToEndId     Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 EndToEndId,omitempty"`
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

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max70Text string

type MessageHeader9 struct {
	MsgId   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 MsgId"`
	CreDtTm ISODateTime        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 CreDtTm,omitempty"`
	ReqTp   RequestType4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 ReqTp,omitempty"`
}

type OriginalMessageAndIssuer1 struct {
	MsgId   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 MsgId"`
	MsgNmId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 MsgNmId,omitempty"`
	OrgtrNm Max70Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 OrgtrNm,omitempty"`
}

type PaymentIdentification6Choice struct {
	TxId      Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 TxId"`
	QId       QueueTransactionIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 QId"`
	LngBizId  LongPaymentIdentification2      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 LngBizId"`
	ShrtBizId ShortPaymentIdentification2     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 ShrtBizId"`
	PrtryId   Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 PrtryId"`
}

// May be one of BDT, BCT, CDT, CCT, CHK, BKT, DCP, CCP, RTI, CAN
type PaymentInstrument1Code string

type PaymentOrigin1Choice struct {
	FINMT    Max3NumericText        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 FINMT"`
	XMLMsgNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 XMLMsgNm"`
	Prtry    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Prtry"`
	Instrm   PaymentInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Instrm"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 AdrLine,omitempty"`
}

type QueueTransactionIdentification1 struct {
	QId    Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 QId"`
	PosInQ Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 PosInQ"`
}

type Receipt3 struct {
	OrgnlMsgId OriginalMessageAndIssuer1    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 OrgnlMsgId"`
	OrgnlPmtId PaymentIdentification6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 OrgnlPmtId,omitempty"`
	ReqHdlg    []RequestHandling1           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 ReqHdlg,omitempty"`
}

type ReceiptV05 struct {
	MsgHdr      MessageHeader9       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 MsgHdr"`
	RctDtls     []Receipt3           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 RctDtls"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 SplmtryData,omitempty"`
}

type RequestHandling1 struct {
	StsCd Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 StsCd"`
	Desc  Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Desc,omitempty"`
}

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 PmtCtrl"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Enqry"`
	Prtry   GenericIdentification1                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Prtry"`
}

type ShortPaymentIdentification2 struct {
	TxId          Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 TxId"`
	IntrBkSttlmDt ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 IntrBkSttlmDt"`
	InstgAgt      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 InstgAgt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
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
