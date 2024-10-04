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

// THIS FILE IS GENERATED BY https://github.com/Zorgbijjou/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// ChargeItemStatus is documented here http://hl7.org/fhir/ValueSet/chargeitem-status
type ChargeItemStatus int

const (
	ChargeItemStatusPlanned ChargeItemStatus = iota
	ChargeItemStatusBillable
	ChargeItemStatusNotBillable
	ChargeItemStatusAborted
	ChargeItemStatusBilled
	ChargeItemStatusEnteredInError
	ChargeItemStatusUnknown
)

func (code ChargeItemStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *ChargeItemStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "planned":
		*code = ChargeItemStatusPlanned
	case "billable":
		*code = ChargeItemStatusBillable
	case "not-billable":
		*code = ChargeItemStatusNotBillable
	case "aborted":
		*code = ChargeItemStatusAborted
	case "billed":
		*code = ChargeItemStatusBilled
	case "entered-in-error":
		*code = ChargeItemStatusEnteredInError
	case "unknown":
		*code = ChargeItemStatusUnknown
	default:
		return fmt.Errorf("unknown ChargeItemStatus code `%s`", s)
	}
	return nil
}
func (code ChargeItemStatus) String() string {
	return code.Code()
}
func (code ChargeItemStatus) Code() string {
	switch code {
	case ChargeItemStatusPlanned:
		return "planned"
	case ChargeItemStatusBillable:
		return "billable"
	case ChargeItemStatusNotBillable:
		return "not-billable"
	case ChargeItemStatusAborted:
		return "aborted"
	case ChargeItemStatusBilled:
		return "billed"
	case ChargeItemStatusEnteredInError:
		return "entered-in-error"
	case ChargeItemStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code ChargeItemStatus) Display() string {
	switch code {
	case ChargeItemStatusPlanned:
		return "Planned"
	case ChargeItemStatusBillable:
		return "Billable"
	case ChargeItemStatusNotBillable:
		return "Not billable"
	case ChargeItemStatusAborted:
		return "Aborted"
	case ChargeItemStatusBilled:
		return "Billed"
	case ChargeItemStatusEnteredInError:
		return "Entered in Error"
	case ChargeItemStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code ChargeItemStatus) Definition() string {
	switch code {
	case ChargeItemStatusPlanned:
		return "The charge item has been entered, but the charged service is not  yet complete, so it shall not be billed yet but might be used in the context of pre-authorization."
	case ChargeItemStatusBillable:
		return "The charge item is ready for billing."
	case ChargeItemStatusNotBillable:
		return "The charge item has been determined to be not billable (e.g. due to rules associated with the billing code)."
	case ChargeItemStatusAborted:
		return "The processing of the charge was aborted."
	case ChargeItemStatusBilled:
		return "The charge item has been billed (e.g. a billing engine has generated financial transactions by applying the associated ruled for the charge item to the context of the Encounter, and placed them into Claims/Invoices."
	case ChargeItemStatusEnteredInError:
		return "The charge item has been entered in error and should not be processed for billing."
	case ChargeItemStatusUnknown:
		return "The authoring system does not know which of the status values currently applies for this charge item  Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply, it's just not known which one."
	}
	return "<unknown>"
}
