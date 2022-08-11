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

// THIS FILE IS GENERATED BY https://github.com/lifebox-healthcare/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// ReferenceVersionRules is documented here http://hl7.org/fhir/ValueSet/reference-version-rules
type ReferenceVersionRules int

const (
	ReferenceVersionRulesEither ReferenceVersionRules = iota
	ReferenceVersionRulesIndependent
	ReferenceVersionRulesSpecific
)

func (code ReferenceVersionRules) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *ReferenceVersionRules) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "either":
		*code = ReferenceVersionRulesEither
	case "independent":
		*code = ReferenceVersionRulesIndependent
	case "specific":
		*code = ReferenceVersionRulesSpecific
	default:
		return fmt.Errorf("unknown ReferenceVersionRules code `%s`", s)
	}
	return nil
}
func (code ReferenceVersionRules) String() string {
	return code.Code()
}
func (code ReferenceVersionRules) Code() string {
	switch code {
	case ReferenceVersionRulesEither:
		return "either"
	case ReferenceVersionRulesIndependent:
		return "independent"
	case ReferenceVersionRulesSpecific:
		return "specific"
	}
	return "<unknown>"
}
func (code ReferenceVersionRules) Display() string {
	switch code {
	case ReferenceVersionRulesEither:
		return "Either Specific or independent"
	case ReferenceVersionRulesIndependent:
		return "Version independent"
	case ReferenceVersionRulesSpecific:
		return "Version Specific"
	}
	return "<unknown>"
}
func (code ReferenceVersionRules) Definition() string {
	switch code {
	case ReferenceVersionRulesEither:
		return "The reference may be either version independent or version specific."
	case ReferenceVersionRulesIndependent:
		return "The reference must be version independent."
	case ReferenceVersionRulesSpecific:
		return "The reference must be version specific."
	}
	return "<unknown>"
}
