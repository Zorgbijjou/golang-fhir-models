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

// GraphDefinition is documented here http://hl7.org/fhir/StructureDefinition/GraphDefinition
type GraphDefinition struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string               `bson:"url,omitempty" json:"url,omitempty"`
	Version           *string               `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                `bson:"name" json:"name"`
	Status            PublicationStatus     `bson:"status" json:"status"`
	Experimental      *bool                 `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string               `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string               `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail       `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string               `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext        `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept     `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string               `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Start             ResourceType          `bson:"start" json:"start"`
	Profile           *string               `bson:"profile,omitempty" json:"profile,omitempty"`
	Link              []GraphDefinitionLink `bson:"link,omitempty" json:"link,omitempty"`
}
type GraphDefinitionLink struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Path              *string                     `bson:"path,omitempty" json:"path,omitempty"`
	SliceName         *string                     `bson:"sliceName,omitempty" json:"sliceName,omitempty"`
	Min               *int                        `bson:"min,omitempty" json:"min,omitempty"`
	Max               *string                     `bson:"max,omitempty" json:"max,omitempty"`
	Description       *string                     `bson:"description,omitempty" json:"description,omitempty"`
	Target            []GraphDefinitionLinkTarget `bson:"target,omitempty" json:"target,omitempty"`
}
type GraphDefinitionLinkTarget struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              ResourceType                           `bson:"type" json:"type"`
	Params            *string                                `bson:"params,omitempty" json:"params,omitempty"`
	Profile           *string                                `bson:"profile,omitempty" json:"profile,omitempty"`
	Compartment       []GraphDefinitionLinkTargetCompartment `bson:"compartment,omitempty" json:"compartment,omitempty"`
	Link              []GraphDefinitionLink                  `bson:"link,omitempty" json:"link,omitempty"`
}
type GraphDefinitionLinkTargetCompartment struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Use               GraphCompartmentUse  `bson:"use" json:"use"`
	Code              CompartmentType      `bson:"code" json:"code"`
	Rule              GraphCompartmentRule `bson:"rule" json:"rule"`
	Expression        *string              `bson:"expression,omitempty" json:"expression,omitempty"`
	Description       *string              `bson:"description,omitempty" json:"description,omitempty"`
}
type OtherGraphDefinition GraphDefinition

// MarshalJSON marshals the given GraphDefinition as JSON into a byte slice
func (r GraphDefinition) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(struct {
		OtherGraphDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherGraphDefinition: OtherGraphDefinition(r),
		ResourceType:         "GraphDefinition",
	})
	return buffer.Bytes(), err
}

// UnmarshalGraphDefinition unmarshals a GraphDefinition.
func UnmarshalGraphDefinition(b []byte) (GraphDefinition, error) {
	var graphDefinition GraphDefinition
	if err := json.Unmarshal(b, &graphDefinition); err != nil {
		return graphDefinition, err
	}
	return graphDefinition, nil
}
