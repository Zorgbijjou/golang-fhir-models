package fhir

// THIS FILE IS NOT GENERATED, IT IS WRITTEN BY HAND

// PrimitiveElement carries the id and extensions of a primitive element. In JSON these are not
// stored on the element itself - a primitive is a plain JSON value and has nowhere to put them -
// but in a sibling property with the same name prefixed with an underscore:
//
//	"text": "Meetprotocol",
//	"_text": { "extension": [ ... ] }
//
// See https://hl7.org/fhir/R4/json.html#primitive. Every primitive field in the generated models
// has a matching `<Field>PrimitiveElement` so these extensions survive a marshal/unmarshal round
// trip instead of being silently dropped.
//
// For a repeating primitive the sibling is an array, positionally matched to the values, with null
// for the values that have no id or extensions - hence the pointers in `[]*PrimitiveElement`.
type PrimitiveElement struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
}
