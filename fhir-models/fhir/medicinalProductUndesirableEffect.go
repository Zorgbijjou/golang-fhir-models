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

// THIS FILE IS GENERATED BY https://github.com/michaelsauter/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// MedicinalProductUndesirableEffect is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductUndesirableEffect
type MedicinalProductUndesirableEffect struct {
	ID                     *string          `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string          `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative       `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Subject                []Reference      `bson:"subject,omitempty" json:"subject,omitempty"`
	SymptomConditionEffect *CodeableConcept `bson:"symptomConditionEffect,omitempty" json:"symptomConditionEffect,omitempty"`
	Classification         *CodeableConcept `bson:"classification,omitempty" json:"classification,omitempty"`
	FrequencyOfOccurrence  *CodeableConcept `bson:"frequencyOfOccurrence,omitempty" json:"frequencyOfOccurrence,omitempty"`
	Population             []Population     `bson:"population,omitempty" json:"population,omitempty"`
}
type OtherMedicinalProductUndesirableEffect MedicinalProductUndesirableEffect

// MarshalJSON marshals the given MedicinalProductUndesirableEffect as JSON into a byte slice
func (r MedicinalProductUndesirableEffect) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(struct {
		OtherMedicinalProductUndesirableEffect
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductUndesirableEffect: OtherMedicinalProductUndesirableEffect(r),
		ResourceType:                           "MedicinalProductUndesirableEffect",
	})
	return buffer.Bytes(), err
}

// UnmarshalMedicinalProductUndesirableEffect unmarshals a MedicinalProductUndesirableEffect.
func UnmarshalMedicinalProductUndesirableEffect(b []byte) (MedicinalProductUndesirableEffect, error) {
	var medicinalProductUndesirableEffect MedicinalProductUndesirableEffect
	if err := json.Unmarshal(b, &medicinalProductUndesirableEffect); err != nil {
		return medicinalProductUndesirableEffect, err
	}
	return medicinalProductUndesirableEffect, nil
}
