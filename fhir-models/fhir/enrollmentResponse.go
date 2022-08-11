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
)

// THIS FILE IS GENERATED BY https://github.com/lifebox-healthcare/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// EnrollmentResponse is documented here http://hl7.org/fhir/StructureDefinition/EnrollmentResponse
type EnrollmentResponse struct {
	ID                *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                       `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                    `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *FinancialResourceStatusCodes `bson:"status,omitempty" json:"status,omitempty"`
	Request           *Reference                    `bson:"request,omitempty" json:"request,omitempty"`
	Outcome           *ClaimProcessingCodes         `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition       *string                       `bson:"disposition,omitempty" json:"disposition,omitempty"`
	Created           *string                       `bson:"created,omitempty" json:"created,omitempty"`
	Organization      *Reference                    `bson:"organization,omitempty" json:"organization,omitempty"`
	RequestProvider   *Reference                    `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
}
type OtherEnrollmentResponse EnrollmentResponse

// MarshalJSON marshals the given EnrollmentResponse as JSON into a byte slice
func (r EnrollmentResponse) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(struct {
		OtherEnrollmentResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherEnrollmentResponse: OtherEnrollmentResponse(r),
		ResourceType:            "EnrollmentResponse",
	})
	return buffer.Bytes(), err
}

// UnmarshalEnrollmentResponse unmarshals a EnrollmentResponse.
func UnmarshalEnrollmentResponse(b []byte) (EnrollmentResponse, error) {
	var enrollmentResponse EnrollmentResponse
	if err := json.Unmarshal(b, &enrollmentResponse); err != nil {
		return enrollmentResponse, err
	}
	return enrollmentResponse, nil
}
