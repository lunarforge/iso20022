// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v02

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAcmt02200102 struct {
	XMLName                 *xml.Name                           `json:",omitempty"`
	Xmlns                   string                              `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                `xml:",omitempty" json:",omitempty"`
	IdModAdvc               IdentificationModificationAdviceV02 `xml:"IdModAdvc"`
}

func (doc DocumentAcmt02200102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02200102) NameSpace() string {
	return utils.DocumentAcmt02200102NameSpace
}

func (doc DocumentAcmt02200102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		IdModAdvc IdentificationModificationAdviceV02 `xml:"IdModAdvc"`
	}
	output.IdModAdvc = doc.IdModAdvc
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt02300102 struct {
	XMLName                 *xml.Name                            `json:",omitempty"`
	Xmlns                   string                               `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                 `xml:",omitempty" json:",omitempty"`
	IdVrfctnReq             IdentificationVerificationRequestV02 `xml:"IdVrfctnReq"`
}

func (doc DocumentAcmt02300102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02300102) NameSpace() string {
	return utils.DocumentAcmt02300102NameSpace
}

func (doc DocumentAcmt02300102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		IdVrfctnReq IdentificationVerificationRequestV02 `xml:"IdVrfctnReq"`
	}
	output.IdVrfctnReq = doc.IdVrfctnReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt02400102 struct {
	XMLName                 *xml.Name                           `json:",omitempty"`
	Xmlns                   string                              `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                `xml:",omitempty" json:",omitempty"`
	IdVrfctnRpt             IdentificationVerificationReportV02 `xml:"IdVrfctnRpt"`
}

func (doc DocumentAcmt02400102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02400102) NameSpace() string {
	return utils.DocumentAcmt02400102NameSpace
}

func (doc DocumentAcmt02400102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		IdVrfctnRpt IdentificationVerificationReportV02 `xml:"IdVrfctnRpt"`
	}
	output.IdVrfctnRpt = doc.IdVrfctnRpt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt03000102 struct {
	XMLName                 *xml.Name                          `json:",omitempty"`
	Xmlns                   string                             `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                               `xml:",omitempty" json:",omitempty"`
	AcctSwtchReqRdrctn      AccountSwitchRequestRedirectionV02 `xml:"AcctSwtchReqRdrctn"`
}

func (doc DocumentAcmt03000102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03000102) NameSpace() string {
	return utils.DocumentAcmt03000102NameSpace
}

func (doc DocumentAcmt03000102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchReqRdrctn AccountSwitchRequestRedirectionV02 `xml:"AcctSwtchReqRdrctn"`
	}
	output.AcctSwtchReqRdrctn = doc.AcctSwtchReqRdrctn
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt03300102 struct {
	XMLName                     *xml.Name                                   `json:",omitempty"`
	Xmlns                       string                                      `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace     bool                                        `xml:",omitempty" json:",omitempty"`
	AcctSwtchNtfyAcctSwtchCmplt AccountSwitchNotifyAccountSwitchCompleteV02 `xml:"AcctSwtchNtfyAcctSwtchCmplt"`
}

func (doc DocumentAcmt03300102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03300102) NameSpace() string {
	return utils.DocumentAcmt03300102NameSpace
}

func (doc DocumentAcmt03300102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchNtfyAcctSwtchCmplt AccountSwitchNotifyAccountSwitchCompleteV02 `xml:"AcctSwtchNtfyAcctSwtchCmplt"`
	}
	output.AcctSwtchNtfyAcctSwtchCmplt = doc.AcctSwtchNtfyAcctSwtchCmplt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt03500102 struct {
	XMLName                 *xml.Name                       `json:",omitempty"`
	Xmlns                   string                          `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                            `xml:",omitempty" json:",omitempty"`
	AcctSwtchPmtRspn        AccountSwitchPaymentResponseV02 `xml:"AcctSwtchPmtRspn"`
}

func (doc DocumentAcmt03500102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03500102) NameSpace() string {
	return utils.DocumentAcmt03500102NameSpace
}

func (doc DocumentAcmt03500102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchPmtRspn AccountSwitchPaymentResponseV02 `xml:"AcctSwtchPmtRspn"`
	}
	output.AcctSwtchPmtRspn = doc.AcctSwtchPmtRspn
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt03700102 struct {
	XMLName                 *xml.Name                          `json:",omitempty"`
	Xmlns                   string                             `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                               `xml:",omitempty" json:",omitempty"`
	AcctSwtchTechRjctn      AccountSwitchTechnicalRejectionV02 `xml:"AcctSwtchTechRjctn"`
}

func (doc DocumentAcmt03700102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03700102) NameSpace() string {
	return utils.DocumentAcmt03700102NameSpace
}

func (doc DocumentAcmt03700102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchTechRjctn AccountSwitchTechnicalRejectionV02 `xml:"AcctSwtchTechRjctn"`
	}

	output.AcctSwtchTechRjctn = doc.AcctSwtchTechRjctn
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}
