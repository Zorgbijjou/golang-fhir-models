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

// ConditionalDeleteStatus is documented here http://hl7.org/fhir/ValueSet/conditional-delete-status
type ConditionalDeleteStatus int

const (
	ConditionalDeleteStatusNotSupported ConditionalDeleteStatus = iota
	ConditionalDeleteStatusSingle
	ConditionalDeleteStatusMultiple
)

func (code ConditionalDeleteStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *ConditionalDeleteStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "not-supported":
		*code = ConditionalDeleteStatusNotSupported
	case "single":
		*code = ConditionalDeleteStatusSingle
	case "multiple":
		*code = ConditionalDeleteStatusMultiple
	default:
		return fmt.Errorf("unknown ConditionalDeleteStatus code `%s`", s)
	}
	return nil
}
func (code ConditionalDeleteStatus) String() string {
	return code.Code()
}
func (code ConditionalDeleteStatus) Code() string {
	switch code {
	case ConditionalDeleteStatusNotSupported:
		return "not-supported"
	case ConditionalDeleteStatusSingle:
		return "single"
	case ConditionalDeleteStatusMultiple:
		return "multiple"
	}
	return "<unknown>"
}
func (code ConditionalDeleteStatus) Display() string {
	switch code {
	case ConditionalDeleteStatusNotSupported:
		return "Not Supported"
	case ConditionalDeleteStatusSingle:
		return "Single Deletes Supported"
	case ConditionalDeleteStatusMultiple:
		return "Multiple Deletes Supported"
	}
	return "<unknown>"
}
func (code ConditionalDeleteStatus) Definition() string {
	switch code {
	case ConditionalDeleteStatusNotSupported:
		return "No support for conditional deletes."
	case ConditionalDeleteStatusSingle:
		return "Conditional deletes are supported, but only single resources at a time."
	case ConditionalDeleteStatusMultiple:
		return "Conditional deletes are supported, and multiple resources can be deleted in a single interaction."
	}
	return "<unknown>"
}
