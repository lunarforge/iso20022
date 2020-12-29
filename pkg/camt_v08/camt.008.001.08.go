// Code generated by xsdgen. DO NOT EDIT.

package camt_v08

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Prtry"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Prtry"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 PstlAdr,omitempty"`
}

type CancelTransactionV08 struct {
	MsgHdr      MessageHeader9               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 MsgHdr"`
	PmtId       PaymentIdentification6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 PmtId"`
	CshAcct     CashAccount38                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 CshAcct,omitempty"`
	CxlRsn      PaymentCancellationReason5   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 CxlRsn,omitempty"`
	SplmtryData []SupplementaryData1         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 SplmtryData,omitempty"`
}

type CancellationReason33Choice struct {
	Cd    ExternalCancellationReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Cd"`
	Prtry Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Prtry"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Id"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 MmbId"`
}

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 PrefrdMtd,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 CtryOfBirth"`
}

type Document struct {
	CclTx CancelTransactionV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 CclTx"`
}

// Must match the pattern [BEOVW]{1,1}[0-9]{2,2}|DUM
type EntryTypeIdentifier string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCancellationReason1Code string

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

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Issr,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Issr,omitempty"`
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

type LongPaymentIdentification2 struct {
	TxId           Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 TxId,omitempty"`
	UETR           UUIDv4Identifier                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 UETR,omitempty"`
	IntrBkSttlmAmt float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 IntrBkSttlmAmt"`
	IntrBkSttlmDt  ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 IntrBkSttlmDt"`
	PmtMtd         PaymentOrigin1Choice                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 PmtMtd,omitempty"`
	InstgAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 InstgAgt"`
	InstdAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 InstdAgt"`
	NtryTp         EntryTypeIdentifier                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 NtryTp,omitempty"`
	EndToEndId     Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 EndToEndId,omitempty"`
}

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
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must be at least 1 items long
type Max4Text string

// Must be at least 1 items long
type Max70Text string

type MessageHeader9 struct {
	MsgId   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 MsgId"`
	CreDtTm ISODateTime        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 CreDtTm,omitempty"`
	ReqTp   RequestType4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 ReqTp,omitempty"`
}

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Prtry"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 OrgId"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 PrvtId"`
}

type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 CtctDtls,omitempty"`
}

type PaymentCancellationReason5 struct {
	Orgtr    PartyIdentification135     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Orgtr,omitempty"`
	Rsn      CancellationReason33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Rsn,omitempty"`
	AddtlInf []Max105Text               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 AddtlInf,omitempty"`
}

type PaymentIdentification6Choice struct {
	TxId      Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 TxId"`
	QId       QueueTransactionIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 QId"`
	LngBizId  LongPaymentIdentification2      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 LngBizId"`
	ShrtBizId ShortPaymentIdentification2     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 ShrtBizId"`
	PrtryId   Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 PrtryId"`
}

// May be one of BDT, BCT, CDT, CCT, CHK, BKT, DCP, CCP, RTI, CAN
type PaymentInstrument1Code string

type PaymentOrigin1Choice struct {
	FINMT    Max3NumericText        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 FINMT"`
	XMLMsgNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 XMLMsgNm"`
	Prtry    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Prtry"`
	Instrm   PaymentInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Instrm"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Cd"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Prtry"`
}

type QueueTransactionIdentification1 struct {
	QId    Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 QId"`
	PosInQ Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 PosInQ"`
}

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 PmtCtrl"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Enqry"`
	Prtry   GenericIdentification1                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Prtry"`
}

type ShortPaymentIdentification2 struct {
	TxId          Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 TxId"`
	IntrBkSttlmDt ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 IntrBkSttlmDt"`
	InstgAgt      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 InstgAgt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 Envlp"`
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
