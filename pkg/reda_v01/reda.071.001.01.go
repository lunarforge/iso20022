// Code generated by xsdgen. DO NOT EDIT.

package reda_v01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type ActivationHeader2 struct {
	MsgId    Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 MsgId"`
	CreDtTm  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CreDtTm"`
	MsgOrgtr RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 MsgOrgtr,omitempty"`
	MsgRcpt  RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 MsgRcpt,omitempty"`
	InitgPty RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 InitgPty"`
}

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Prtry"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 PrefrdMtd,omitempty"`
}

type ContractReference1 struct {
	Tp  DocumentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Tp,omitempty"`
	Ref Max500Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Ref"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DtTm"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CtryOfBirth"`
}

type DebtorActivation3 struct {
	DbtrActvtnId      Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DbtrActvtnId,omitempty"`
	DispNm            Max140Text              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DispNm,omitempty"`
	UltmtDbtr         RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 UltmtDbtr,omitempty"`
	Dbtr              RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Dbtr"`
	DbtrSolPrvdr      RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DbtrSolPrvdr"`
	CstmrId           []Party49Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CstmrId,omitempty"`
	CtrctFrmtTp       []DocumentFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CtrctFrmtTp,omitempty"`
	CtrctRef          []ContractReference1    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CtrctRef,omitempty"`
	Cdtr              RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Cdtr"`
	UltmtCdtr         RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 UltmtCdtr,omitempty"`
	ActvtnReqDlvryPty RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 ActvtnReqDlvryPty,omitempty"`
	StartDt           DateAndDateTime2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 StartDt,omitempty"`
	EndDt             DateAndDateTime2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 EndDt,omitempty"`
	DdctdActvtnCd     Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DdctdActvtnCd,omitempty"`
}

type DebtorActivation4 struct {
	DbtrActvtnId      Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DbtrActvtnId,omitempty"`
	DispNm            Max140Text              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DispNm,omitempty"`
	UltmtDbtr         RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 UltmtDbtr,omitempty"`
	Dbtr              RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Dbtr,omitempty"`
	DbtrSolPrvdr      RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DbtrSolPrvdr,omitempty"`
	CstmrId           []Party49Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CstmrId,omitempty"`
	CtrctFrmtTp       []DocumentFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CtrctFrmtTp,omitempty"`
	CtrctRef          []ContractReference1    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CtrctRef,omitempty"`
	Cdtr              RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Cdtr,omitempty"`
	UltmtCdtr         RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 UltmtCdtr,omitempty"`
	ActvtnReqDlvryPty RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 ActvtnReqDlvryPty,omitempty"`
	StartDt           DateAndDateTime2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 StartDt,omitempty"`
	EndDt             DateAndDateTime2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 EndDt,omitempty"`
	DdctdActvtnCd     Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DdctdActvtnCd,omitempty"`
}

type DebtorActivationAmendment3 struct {
	OrgnlBizInstr OriginalBusinessInstruction1     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 OrgnlBizInstr,omitempty"`
	AmdmntRsn     DebtorActivationAmendmentReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 AmdmntRsn,omitempty"`
	Amdmnt        DebtorActivationAmendment4       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Amdmnt"`
	OrgnlActvtn   OriginalActivation2Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 OrgnlActvtn"`
	SplmtryData   []SupplementaryData1             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 SplmtryData,omitempty"`
}

type DebtorActivationAmendment4 struct {
	DbtrActvtn      DebtorActivation4  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DbtrActvtn,omitempty"`
	ElctrncInvcData ElectronicInvoice1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 ElctrncInvcData,omitempty"`
}

type DebtorActivationAmendmentReason1Choice struct {
	Cd    ExternalDebtorActivationAmendmentReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Cd"`
	Prtry Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Prtry"`
}

type DebtorActivationAmendmentReason2 struct {
	Orgtr    RTPPartyIdentification1                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Orgtr,omitempty"`
	Rsn      DebtorActivationAmendmentReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Rsn"`
	AddtlInf []Max105Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 AddtlInf,omitempty"`
}

type Document struct {
	ReqToPayDbtrActvtnAmdmntReq RequestToPayDebtorActivationAmendmentRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 ReqToPayDbtrActvtnAmdmntReq"`
}

type DocumentFormat2Choice struct {
	Cd    ExternalDocumentFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Cd"`
	Prtry GenericIdentification1      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Prtry"`
}

type DocumentType1Choice struct {
	Cd    ExternalDocumentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Cd"`
	Prtry GenericIdentification1    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Prtry"`
}

type ElectronicInvoice1 struct {
	PresntmntTp PresentmentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 PresntmntTp"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalDebtorActivationAmendmentReason1Code string

// Must be at least 1 items long
type ExternalDocumentFormat1Code string

// Must be at least 1 items long
type ExternalDocumentType1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Issr,omitempty"`
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

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max128Text string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max256Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max4Text string

// Must be at least 1 items long
type Max500Text string

// Must be at least 1 items long
type Max70Text string

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type OrganisationIdentification37 struct {
	AnyBIC   AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 AnyBIC,omitempty"`
	LEI      LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 LEI,omitempty"`
	EmailAdr Max256Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 EmailAdr,omitempty"`
	Othr     []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Prtry"`
}

type OriginalActivation2Choice struct {
	OrgnlDbtrId     Party49Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 OrgnlDbtrId"`
	OrgnlActvtnData DebtorActivation3 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 OrgnlActvtnData"`
}

type OriginalBusinessInstruction1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 MsgId"`
	MsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 MsgNmId,omitempty"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CreDtTm,omitempty"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Id,omitempty"`
}

type Party49Choice struct {
	OrgId  OrganisationIdentification37 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 OrgId"`
	PrvtId PersonIdentification17       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 PrvtId"`
}

type PersonIdentification17 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DtAndPlcOfBirth,omitempty"`
	EmailAdr        Max256Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 EmailAdr,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

// May be one of FULL, PAYD
type PresentmentType1Code string

type RTPPartyIdentification1 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 PstlAdr,omitempty"`
	Id        Party49Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CtctDtls,omitempty"`
}

type RequestToPayDebtorActivationAmendmentRequestV01 struct {
	Hdr         ActivationHeader2            `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Hdr"`
	AmdmntData  []DebtorActivationAmendment3 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 AmdmntData"`
	SplmtryData []SupplementaryData1         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 SplmtryData,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Envlp"`
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
