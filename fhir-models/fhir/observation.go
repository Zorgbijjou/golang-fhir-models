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

// THIS FILE IS GENERATED BY https://github.com/Zorgbijjou/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// Observation is documented here http://hl7.org/fhir/StructureDefinition/Observation
type Observation struct {
	ID                   *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn              []Reference                 `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf               []Reference                 `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status               ObservationStatus           `bson:"status" json:"status"`
	Category             []CodeableConcept           `bson:"category,omitempty" json:"category,omitempty"`
	Code                 CodeableConcept             `bson:"code" json:"code"`
	Subject              *Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Focus                []Reference                 `bson:"focus,omitempty" json:"focus,omitempty"`
	Encounter            *Reference                  `bson:"encounter,omitempty" json:"encounter,omitempty"`
	EffectiveDateTime    *string                     `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod      *Period                     `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	EffectiveTiming      *Timing                     `bson:"effectiveTiming,omitempty" json:"effectiveTiming,omitempty"`
	EffectiveInstant     *string                     `bson:"effectiveInstant,omitempty" json:"effectiveInstant,omitempty"`
	Issued               *string                     `bson:"issued,omitempty" json:"issued,omitempty"`
	Performer            []Reference                 `bson:"performer,omitempty" json:"performer,omitempty"`
	ValueQuantity        *Quantity                   `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept            `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueString          *string                     `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean         *bool                       `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueInteger         *int                        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueRange           *Range                      `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio                      `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData                `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueTime            *string                     `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueDateTime        *string                     `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod          *Period                     `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept            `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation       []CodeableConcept           `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	Note                 []Annotation                `bson:"note,omitempty" json:"note,omitempty"`
	BodySite             *CodeableConcept            `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Method               *CodeableConcept            `bson:"method,omitempty" json:"method,omitempty"`
	Specimen             *Reference                  `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Device               *Reference                  `bson:"device,omitempty" json:"device,omitempty"`
	ReferenceRange       []ObservationReferenceRange `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
	HasMember            []Reference                 `bson:"hasMember,omitempty" json:"hasMember,omitempty"`
	DerivedFrom          []Reference                 `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Component            []ObservationComponent      `bson:"component,omitempty" json:"component,omitempty"`
}
type ObservationReferenceRange struct {
	ID                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Low               *Quantity         `bson:"low,omitempty" json:"low,omitempty"`
	High              *Quantity         `bson:"high,omitempty" json:"high,omitempty"`
	Type              *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	AppliesTo         []CodeableConcept `bson:"appliesTo,omitempty" json:"appliesTo,omitempty"`
	Age               *Range            `bson:"age,omitempty" json:"age,omitempty"`
	Text              *string           `bson:"text,omitempty" json:"text,omitempty"`
}
type ObservationComponent struct {
	ID                   *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code                 CodeableConcept             `bson:"code" json:"code"`
	ValueQuantity        *Quantity                   `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept            `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueString          *string                     `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean         *bool                       `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueInteger         *int                        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueRange           *Range                      `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio                      `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData                `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueTime            *string                     `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueDateTime        *string                     `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod          *Period                     `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept            `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation       []CodeableConcept           `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	ReferenceRange       []ObservationReferenceRange `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
}
type OtherObservation Observation

// MarshalJSON marshals the given Observation as JSON into a byte slice
func (r Observation) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(struct {
		OtherObservation
		ResourceType string `json:"resourceType"`
	}{
		OtherObservation: OtherObservation(r),
		ResourceType:     "Observation",
	})
	return buffer.Bytes(), err
}

// UnmarshalObservation unmarshals a Observation.
func UnmarshalObservation(b []byte) (Observation, error) {
	var observation Observation
	if err := json.Unmarshal(b, &observation); err != nil {
		return observation, err
	}
	return observation, nil
}
