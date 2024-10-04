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

// AuditEventAction is documented here http://hl7.org/fhir/ValueSet/audit-event-action
type AuditEventAction int

const (
	AuditEventActionC AuditEventAction = iota
	AuditEventActionR
	AuditEventActionU
	AuditEventActionD
	AuditEventActionE
)

func (code AuditEventAction) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *AuditEventAction) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "C":
		*code = AuditEventActionC
	case "R":
		*code = AuditEventActionR
	case "U":
		*code = AuditEventActionU
	case "D":
		*code = AuditEventActionD
	case "E":
		*code = AuditEventActionE
	default:
		return fmt.Errorf("unknown AuditEventAction code `%s`", s)
	}
	return nil
}
func (code AuditEventAction) String() string {
	return code.Code()
}
func (code AuditEventAction) Code() string {
	switch code {
	case AuditEventActionC:
		return "C"
	case AuditEventActionR:
		return "R"
	case AuditEventActionU:
		return "U"
	case AuditEventActionD:
		return "D"
	case AuditEventActionE:
		return "E"
	}
	return "<unknown>"
}
func (code AuditEventAction) Display() string {
	switch code {
	case AuditEventActionC:
		return "Create"
	case AuditEventActionR:
		return "Read/View/Print"
	case AuditEventActionU:
		return "Update"
	case AuditEventActionD:
		return "Delete"
	case AuditEventActionE:
		return "Execute"
	}
	return "<unknown>"
}
func (code AuditEventAction) Definition() string {
	switch code {
	case AuditEventActionC:
		return "Create a new database object, such as placing an order."
	case AuditEventActionR:
		return "Display or print data, such as a doctor census."
	case AuditEventActionU:
		return "Update data, such as revise patient information."
	case AuditEventActionD:
		return "Delete items, such as a doctor master file record."
	case AuditEventActionE:
		return "Perform a system or application function such as log-on, program execution or use of an object's method, or perform a query/search operation."
	}
	return "<unknown>"
}
