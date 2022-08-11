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

// MedicinalProductManufactured is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductManufactured
type MedicinalProductManufactured struct {
	ID                      *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta                    *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules           *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text                    *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	Extension               []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ManufacturedDoseForm    CodeableConcept     `bson:"manufacturedDoseForm" json:"manufacturedDoseForm"`
	UnitOfPresentation      *CodeableConcept    `bson:"unitOfPresentation,omitempty" json:"unitOfPresentation,omitempty"`
	Quantity                Quantity            `bson:"quantity" json:"quantity"`
	Manufacturer            []Reference         `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Ingredient              []Reference         `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	PhysicalCharacteristics *ProdCharacteristic `bson:"physicalCharacteristics,omitempty" json:"physicalCharacteristics,omitempty"`
	OtherCharacteristics    []CodeableConcept   `bson:"otherCharacteristics,omitempty" json:"otherCharacteristics,omitempty"`
}
type OtherMedicinalProductManufactured MedicinalProductManufactured

// MarshalJSON marshals the given MedicinalProductManufactured as JSON into a byte slice
func (r MedicinalProductManufactured) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(struct {
		OtherMedicinalProductManufactured
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductManufactured: OtherMedicinalProductManufactured(r),
		ResourceType:                      "MedicinalProductManufactured",
	})
	return buffer.Bytes(), err
}

// UnmarshalMedicinalProductManufactured unmarshals a MedicinalProductManufactured.
func UnmarshalMedicinalProductManufactured(b []byte) (MedicinalProductManufactured, error) {
	var medicinalProductManufactured MedicinalProductManufactured
	if err := json.Unmarshal(b, &medicinalProductManufactured); err != nil {
		return medicinalProductManufactured, err
	}
	return medicinalProductManufactured, nil
}
