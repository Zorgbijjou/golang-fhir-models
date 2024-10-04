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

// CompartmentDefinition is documented here http://hl7.org/fhir/StructureDefinition/CompartmentDefinition
type CompartmentDefinition struct {
	ID                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                          `bson:"url" json:"url"`
	Version           *string                         `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                          `bson:"name" json:"name"`
	Status            PublicationStatus               `bson:"status" json:"status"`
	Experimental      *bool                           `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                         `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                         `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                 `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                         `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext                  `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Purpose           *string                         `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Code              CompartmentType                 `bson:"code" json:"code"`
	Search            bool                            `bson:"search" json:"search"`
	Resource          []CompartmentDefinitionResource `bson:"resource,omitempty" json:"resource,omitempty"`
}
type CompartmentDefinitionResource struct {
	ID                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              ResourceType `bson:"code" json:"code"`
	Param             []string     `bson:"param,omitempty" json:"param,omitempty"`
	Documentation     *string      `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type OtherCompartmentDefinition CompartmentDefinition

// MarshalJSON marshals the given CompartmentDefinition as JSON into a byte slice
func (r CompartmentDefinition) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(struct {
		OtherCompartmentDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherCompartmentDefinition: OtherCompartmentDefinition(r),
		ResourceType:               "CompartmentDefinition",
	})
	return buffer.Bytes(), err
}

// UnmarshalCompartmentDefinition unmarshals a CompartmentDefinition.
func UnmarshalCompartmentDefinition(b []byte) (CompartmentDefinition, error) {
	var compartmentDefinition CompartmentDefinition
	if err := json.Unmarshal(b, &compartmentDefinition); err != nil {
		return compartmentDefinition, err
	}
	return compartmentDefinition, nil
}
