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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/lifebox-healthcare/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// Invoice is documented here http://hl7.org/fhir/StructureDefinition/Invoice
type Invoice struct {
	Id                  *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	Extension           []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              InvoiceStatus                   `bson:"status" json:"status"`
	CancelledReason     *string                         `bson:"cancelledReason,omitempty" json:"cancelledReason,omitempty"`
	Type                *CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	Subject             *Reference                      `bson:"subject,omitempty" json:"subject,omitempty"`
	Recipient           *Reference                      `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Date                *string                         `bson:"date,omitempty" json:"date,omitempty"`
	Participant         []InvoiceParticipant            `bson:"participant,omitempty" json:"participant,omitempty"`
	Issuer              *Reference                      `bson:"issuer,omitempty" json:"issuer,omitempty"`
	Account             *Reference                      `bson:"account,omitempty" json:"account,omitempty"`
	LineItem            []InvoiceLineItem               `bson:"lineItem,omitempty" json:"lineItem,omitempty"`
	TotalPriceComponent []InvoiceLineItemPriceComponent `bson:"totalPriceComponent,omitempty" json:"totalPriceComponent,omitempty"`
	TotalNet            *Money                          `bson:"totalNet,omitempty" json:"totalNet,omitempty"`
	TotalGross          *Money                          `bson:"totalGross,omitempty" json:"totalGross,omitempty"`
	PaymentTerms        *string                         `bson:"paymentTerms,omitempty" json:"paymentTerms,omitempty"`
	Note                []Annotation                    `bson:"note,omitempty" json:"note,omitempty"`
}
type InvoiceParticipant struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Actor             Reference        `bson:"actor" json:"actor"`
}
type InvoiceLineItem struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          *int                            `bson:"sequence,omitempty" json:"sequence,omitempty"`
	PriceComponent    []InvoiceLineItemPriceComponent `bson:"priceComponent,omitempty" json:"priceComponent,omitempty"`
}
type InvoiceLineItemPriceComponent struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              InvoicePriceComponentType `bson:"type" json:"type"`
	Code              *CodeableConcept          `bson:"code,omitempty" json:"code,omitempty"`
	Factor            *string                   `bson:"factor,omitempty" json:"factor,omitempty"`
	Amount            *Money                    `bson:"amount,omitempty" json:"amount,omitempty"`
}
type OtherInvoice Invoice

// MarshalJSON marshals the given Invoice as JSON into a byte slice
func (r Invoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherInvoice
		ResourceType string `json:"resourceType"`
	}{
		OtherInvoice: OtherInvoice(r),
		ResourceType: "Invoice",
	})
}

// UnmarshalInvoice unmarshals a Invoice.
func UnmarshalInvoice(b []byte) (Invoice, error) {
	var invoice Invoice
	if err := json.Unmarshal(b, &invoice); err != nil {
		return invoice, err
	}
	return invoice, nil
}
