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

// FlagStatus is documented here http://hl7.org/fhir/ValueSet/flag-status
type FlagStatus int

const (
	FlagStatusActive FlagStatus = iota
	FlagStatusInactive
	FlagStatusEnteredInError
)

func (code FlagStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *FlagStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = FlagStatusActive
	case "inactive":
		*code = FlagStatusInactive
	case "entered-in-error":
		*code = FlagStatusEnteredInError
	default:
		return fmt.Errorf("unknown FlagStatus code `%s`", s)
	}
	return nil
}
func (code FlagStatus) String() string {
	return code.Code()
}
func (code FlagStatus) Code() string {
	switch code {
	case FlagStatusActive:
		return "active"
	case FlagStatusInactive:
		return "inactive"
	case FlagStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code FlagStatus) Display() string {
	switch code {
	case FlagStatusActive:
		return "Active"
	case FlagStatusInactive:
		return "Inactive"
	case FlagStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code FlagStatus) Definition() string {
	switch code {
	case FlagStatusActive:
		return "A current flag that should be displayed to a user. A system may use the category to determine which user roles should view the flag."
	case FlagStatusInactive:
		return "The flag no longer needs to be displayed."
	case FlagStatusEnteredInError:
		return "The flag was added in error and should no longer be displayed."
	}
	return "<unknown>"
}
