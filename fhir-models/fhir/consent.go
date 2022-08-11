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

// Consent is documented here http://hl7.org/fhir/StructureDefinition/Consent
type Consent struct {
	ID                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            ConsentState          `bson:"status" json:"status"`
	Scope             CodeableConcept       `bson:"scope" json:"scope"`
	Category          []CodeableConcept     `bson:"category" json:"category"`
	Patient           *Reference            `bson:"patient,omitempty" json:"patient,omitempty"`
	DateTime          *string               `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Performer         []Reference           `bson:"performer,omitempty" json:"performer,omitempty"`
	Organization      []Reference           `bson:"organization,omitempty" json:"organization,omitempty"`
	SourceAttachment  *Attachment           `bson:"sourceAttachment,omitempty" json:"sourceAttachment,omitempty"`
	SourceReference   *Reference            `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
	Policy            []ConsentPolicy       `bson:"policy,omitempty" json:"policy,omitempty"`
	PolicyRule        *CodeableConcept      `bson:"policyRule,omitempty" json:"policyRule,omitempty"`
	Verification      []ConsentVerification `bson:"verification,omitempty" json:"verification,omitempty"`
	Provision         *ConsentProvision     `bson:"provision,omitempty" json:"provision,omitempty"`
}
type ConsentPolicy struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Authority         *string     `bson:"authority,omitempty" json:"authority,omitempty"`
	Uri               *string     `bson:"uri,omitempty" json:"uri,omitempty"`
}
type ConsentVerification struct {
	ID                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Verified          bool        `bson:"verified" json:"verified"`
	VerifiedWith      *Reference  `bson:"verifiedWith,omitempty" json:"verifiedWith,omitempty"`
	VerificationDate  *string     `bson:"verificationDate,omitempty" json:"verificationDate,omitempty"`
}
type ConsentProvision struct {
	ID                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *ConsentProvisionType   `bson:"type,omitempty" json:"type,omitempty"`
	Period            *Period                 `bson:"period,omitempty" json:"period,omitempty"`
	Actor             []ConsentProvisionActor `bson:"actor,omitempty" json:"actor,omitempty"`
	Action            []CodeableConcept       `bson:"action,omitempty" json:"action,omitempty"`
	SecurityLabel     []Coding                `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Purpose           []Coding                `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Class             []Coding                `bson:"class,omitempty" json:"class,omitempty"`
	Code              []CodeableConcept       `bson:"code,omitempty" json:"code,omitempty"`
	DataPeriod        *Period                 `bson:"dataPeriod,omitempty" json:"dataPeriod,omitempty"`
	Data              []ConsentProvisionData  `bson:"data,omitempty" json:"data,omitempty"`
	Provision         []ConsentProvision      `bson:"provision,omitempty" json:"provision,omitempty"`
}
type ConsentProvisionActor struct {
	ID                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              CodeableConcept `bson:"role" json:"role"`
	Reference         Reference       `bson:"reference" json:"reference"`
}
type ConsentProvisionData struct {
	ID                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Meaning           ConsentDataMeaning `bson:"meaning" json:"meaning"`
	Reference         Reference          `bson:"reference" json:"reference"`
}
type OtherConsent Consent

// MarshalJSON marshals the given Consent as JSON into a byte slice
func (r Consent) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(struct {
		OtherConsent
		ResourceType string `json:"resourceType"`
	}{
		OtherConsent: OtherConsent(r),
		ResourceType: "Consent",
	})
	return buffer.Bytes(), err
}

// UnmarshalConsent unmarshals a Consent.
func UnmarshalConsent(b []byte) (Consent, error) {
	var consent Consent
	if err := json.Unmarshal(b, &consent); err != nil {
		return consent, err
	}
	return consent, nil
}
