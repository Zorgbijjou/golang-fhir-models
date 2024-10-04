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

// LocationStatus is documented here http://hl7.org/fhir/ValueSet/location-status
type LocationStatus int

const (
	LocationStatusActive LocationStatus = iota
	LocationStatusSuspended
	LocationStatusInactive
)

func (code LocationStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *LocationStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = LocationStatusActive
	case "suspended":
		*code = LocationStatusSuspended
	case "inactive":
		*code = LocationStatusInactive
	default:
		return fmt.Errorf("unknown LocationStatus code `%s`", s)
	}
	return nil
}
func (code LocationStatus) String() string {
	return code.Code()
}
func (code LocationStatus) Code() string {
	switch code {
	case LocationStatusActive:
		return "active"
	case LocationStatusSuspended:
		return "suspended"
	case LocationStatusInactive:
		return "inactive"
	}
	return "<unknown>"
}
func (code LocationStatus) Display() string {
	switch code {
	case LocationStatusActive:
		return "Active"
	case LocationStatusSuspended:
		return "Suspended"
	case LocationStatusInactive:
		return "Inactive"
	}
	return "<unknown>"
}
func (code LocationStatus) Definition() string {
	switch code {
	case LocationStatusActive:
		return "The location is operational."
	case LocationStatusSuspended:
		return "The location is temporarily closed."
	case LocationStatusInactive:
		return "The location is no longer used."
	}
	return "<unknown>"
}
