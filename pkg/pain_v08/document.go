// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v08

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentPain01400108 struct {
	CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV08 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 CdtrPmtActvtnReqStsRpt"`
}

func (doc DocumentPain01400108) Validate() error {
	return utils.Validate(&doc)
}

type DocumentPain01300108 struct {
	CdtrPmtActvtnReq CreditorPaymentActivationRequestV08 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CdtrPmtActvtnReq"`
}

func (doc DocumentPain01300108) Validate() error {
	return utils.Validate(&doc)
}
