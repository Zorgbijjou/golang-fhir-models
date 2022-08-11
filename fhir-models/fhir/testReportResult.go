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

// TestReportResult is documented here http://hl7.org/fhir/ValueSet/report-result-codes
type TestReportResult int

const (
	TestReportResultPass TestReportResult = iota
	TestReportResultFail
	TestReportResultPending
)

func (code TestReportResult) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *TestReportResult) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "pass":
		*code = TestReportResultPass
	case "fail":
		*code = TestReportResultFail
	case "pending":
		*code = TestReportResultPending
	default:
		return fmt.Errorf("unknown TestReportResult code `%s`", s)
	}
	return nil
}
func (code TestReportResult) String() string {
	return code.Code()
}
func (code TestReportResult) Code() string {
	switch code {
	case TestReportResultPass:
		return "pass"
	case TestReportResultFail:
		return "fail"
	case TestReportResultPending:
		return "pending"
	}
	return "<unknown>"
}
func (code TestReportResult) Display() string {
	switch code {
	case TestReportResultPass:
		return "Pass"
	case TestReportResultFail:
		return "Fail"
	case TestReportResultPending:
		return "Pending"
	}
	return "<unknown>"
}
func (code TestReportResult) Definition() string {
	switch code {
	case TestReportResultPass:
		return "All test operations successfully passed all asserts."
	case TestReportResultFail:
		return "One or more test operations failed one or more asserts."
	case TestReportResultPending:
		return "One or more test operations is pending execution completion."
	}
	return "<unknown>"
}
