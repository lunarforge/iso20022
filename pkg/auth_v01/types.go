// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v01

import (
	"reflect"
	"regexp"

	"github.com/moov-io/iso20022/pkg/utils"
)

// Must match the pattern [0-9]{8,28}
type Min8Max28NumericText string

func (r Min8Max28NumericText) Validate() error {
	reg := regexp.MustCompile(`[0-9]{8,28}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("Min8Max28NumericText")
	}
	return nil
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

func (r NamePrefix1Code) Validate() error {
	for _, vv := range []string{
		"DOCT", "MIST", "MISS", "MADM",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("NamePrefix1Code")
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

func (r ActiveOrHistoricCurrencyCode) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{3,3}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("ActiveOrHistoricCurrencyCode")
	}
	return nil
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

func (r AnyBICIdentifier) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("AnyBICIdentifier")
	}
	return nil
}

// May be one of FOUN, NFOU, NOAP
type InvestigationStatus1Code string

func (r InvestigationStatus1Code) Validate() error {
	for _, vv := range []string{
		"FOUN", "NFOU", "NOAP",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("InvestigationStatus1Code")
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICFIIdentifier string

func (r BICFIIdentifier) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("BICFIIdentifier")
	}
	return nil
}

// May be one of DTTX, OREC
type TransactionRequestType1Code string

func (r TransactionRequestType1Code) Validate() error {
	for _, vv := range []string{
		"DTTX", "OREC",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("TransactionRequestType1Code")
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

func (r CountryCode) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{2,2}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("CountryCode")
	}
	return nil
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

func (r ExternalAccountIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalAccountIdentification1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalCashAccountType1Code string

func (r ExternalCashAccountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalCashAccountType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

func (r ExternalClearingSystemIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 5 {
		return utils.NewErrTextLengthInvalid("ExternalClearingSystemIdentification1Code", 1, 5)
	}
	return nil
}

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

func (r ExternalFinancialInstitutionIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalFinancialInstitutionIdentification1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

func (r ExternalOrganisationIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalOrganisationIdentification1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

func (r ExternalPersonIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalPersonIdentification1Code", 1, 4)
	}
	return nil
}

// May be one of NRES, PART, COMP
type StatusResponse1Code string

func (r StatusResponse1Code) Validate() error {
	for _, vv := range []string{
		"NRES", "PART", "COMP",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("StatusResponse1Code")
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

func (r IBAN2007Identifier) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("IBAN2007Identifier")
	}
	return nil
}

// May be one of ALLP, OWNE
type InvestigatedParties1Code string

func (r InvestigatedParties1Code) Validate() error {
	for _, vv := range []string{
		"ALLP", "OWNE",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("InvestigatedParties1Code")
}
