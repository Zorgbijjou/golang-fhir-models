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

// GraphCompartmentUse is documented here http://hl7.org/fhir/ValueSet/graph-compartment-use
type GraphCompartmentUse int

const (
	GraphCompartmentUseCondition GraphCompartmentUse = iota
	GraphCompartmentUseRequirement
)

func (code GraphCompartmentUse) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *GraphCompartmentUse) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "condition":
		*code = GraphCompartmentUseCondition
	case "requirement":
		*code = GraphCompartmentUseRequirement
	default:
		return fmt.Errorf("unknown GraphCompartmentUse code `%s`", s)
	}
	return nil
}
func (code GraphCompartmentUse) String() string {
	return code.Code()
}
func (code GraphCompartmentUse) Code() string {
	switch code {
	case GraphCompartmentUseCondition:
		return "condition"
	case GraphCompartmentUseRequirement:
		return "requirement"
	}
	return "<unknown>"
}
func (code GraphCompartmentUse) Display() string {
	switch code {
	case GraphCompartmentUseCondition:
		return "Condition"
	case GraphCompartmentUseRequirement:
		return "Requirement"
	}
	return "<unknown>"
}
func (code GraphCompartmentUse) Definition() string {
	switch code {
	case GraphCompartmentUseCondition:
		return "This compartment rule is a condition for whether the rule applies."
	case GraphCompartmentUseRequirement:
		return "This compartment rule is enforced on any relationships that meet the conditions."
	}
	return "<unknown>"
}
