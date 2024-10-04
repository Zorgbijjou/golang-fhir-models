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

// TestScriptRequestMethodCode is documented here http://hl7.org/fhir/ValueSet/http-operations
type TestScriptRequestMethodCode int

const (
	TestScriptRequestMethodCodeDelete TestScriptRequestMethodCode = iota
	TestScriptRequestMethodCodeGet
	TestScriptRequestMethodCodeOptions
	TestScriptRequestMethodCodePatch
	TestScriptRequestMethodCodePost
	TestScriptRequestMethodCodePut
	TestScriptRequestMethodCodeHead
)

func (code TestScriptRequestMethodCode) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *TestScriptRequestMethodCode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "delete":
		*code = TestScriptRequestMethodCodeDelete
	case "get":
		*code = TestScriptRequestMethodCodeGet
	case "options":
		*code = TestScriptRequestMethodCodeOptions
	case "patch":
		*code = TestScriptRequestMethodCodePatch
	case "post":
		*code = TestScriptRequestMethodCodePost
	case "put":
		*code = TestScriptRequestMethodCodePut
	case "head":
		*code = TestScriptRequestMethodCodeHead
	default:
		return fmt.Errorf("unknown TestScriptRequestMethodCode code `%s`", s)
	}
	return nil
}
func (code TestScriptRequestMethodCode) String() string {
	return code.Code()
}
func (code TestScriptRequestMethodCode) Code() string {
	switch code {
	case TestScriptRequestMethodCodeDelete:
		return "delete"
	case TestScriptRequestMethodCodeGet:
		return "get"
	case TestScriptRequestMethodCodeOptions:
		return "options"
	case TestScriptRequestMethodCodePatch:
		return "patch"
	case TestScriptRequestMethodCodePost:
		return "post"
	case TestScriptRequestMethodCodePut:
		return "put"
	case TestScriptRequestMethodCodeHead:
		return "head"
	}
	return "<unknown>"
}
func (code TestScriptRequestMethodCode) Display() string {
	switch code {
	case TestScriptRequestMethodCodeDelete:
		return "DELETE"
	case TestScriptRequestMethodCodeGet:
		return "GET"
	case TestScriptRequestMethodCodeOptions:
		return "OPTIONS"
	case TestScriptRequestMethodCodePatch:
		return "PATCH"
	case TestScriptRequestMethodCodePost:
		return "POST"
	case TestScriptRequestMethodCodePut:
		return "PUT"
	case TestScriptRequestMethodCodeHead:
		return "HEAD"
	}
	return "<unknown>"
}
func (code TestScriptRequestMethodCode) Definition() string {
	switch code {
	case TestScriptRequestMethodCodeDelete:
		return "HTTP DELETE operation."
	case TestScriptRequestMethodCodeGet:
		return "HTTP GET operation."
	case TestScriptRequestMethodCodeOptions:
		return "HTTP OPTIONS operation."
	case TestScriptRequestMethodCodePatch:
		return "HTTP PATCH operation."
	case TestScriptRequestMethodCodePost:
		return "HTTP POST operation."
	case TestScriptRequestMethodCodePut:
		return "HTTP PUT operation."
	case TestScriptRequestMethodCodeHead:
		return "HTTP HEAD operation."
	}
	return "<unknown>"
}
