// Copyright 2019 The Samply Development Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/michaelsauter/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// AccountStatus is documented here http://hl7.org/fhir/ValueSet/account-status
type AccountStatus int

const (
	AccountStatusActive AccountStatus = iota
	AccountStatusInactive
	AccountStatusEnteredInError
	AccountStatusOnHold
	AccountStatusUnknown
)

func (code AccountStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *AccountStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = AccountStatusActive
	case "inactive":
		*code = AccountStatusInactive
	case "entered-in-error":
		*code = AccountStatusEnteredInError
	case "on-hold":
		*code = AccountStatusOnHold
	case "unknown":
		*code = AccountStatusUnknown
	default:
		return fmt.Errorf("unknown AccountStatus code `%s`", s)
	}
	return nil
}
func (code AccountStatus) String() string {
	return code.Code()
}
func (code AccountStatus) Code() string {
	switch code {
	case AccountStatusActive:
		return "active"
	case AccountStatusInactive:
		return "inactive"
	case AccountStatusEnteredInError:
		return "entered-in-error"
	case AccountStatusOnHold:
		return "on-hold"
	case AccountStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code AccountStatus) Display() string {
	switch code {
	case AccountStatusActive:
		return "Active"
	case AccountStatusInactive:
		return "Inactive"
	case AccountStatusEnteredInError:
		return "Entered in error"
	case AccountStatusOnHold:
		return "On Hold"
	case AccountStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code AccountStatus) Definition() string {
	switch code {
	case AccountStatusActive:
		return "This account is active and may be used."
	case AccountStatusInactive:
		return "This account is inactive and should not be used to track financial information."
	case AccountStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	case AccountStatusOnHold:
		return "This account is on hold."
	case AccountStatusUnknown:
		return "The account status is unknown."
	}
	return "<unknown>"
}
