package fhir

import (
	"encoding/json"
	"testing"
)

// A primitive's extensions live in the "_"-prefixed sibling property. They used to be dropped,
// which silently stripped rendering extensions from questionnaires passing through the models.
func TestQuestionnaireItem_TextExtensionsSurviveRoundTrip(t *testing.T) {
	input := []byte(`{
	  "linkId": "71fc98ad-2dba-4ecd-91d4-a70807dc72ac",
	  "type": "choice",
	  "text": "Meetprotocol",
	  "_text": {
	    "extension": [
	      {
	        "url": "http://hl7.org/fhir/StructureDefinition/rendering-xhtml",
	        "valueString": "<style>div { padding: 0; }</style>Meetprotocol"
	      }
	    ]
	  }
	}`)

	var item QuestionnaireItem
	if err := json.Unmarshal(input, &item); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if item.TextPrimitiveElement == nil {
		t.Fatal("expected _text to be unmarshalled, got nil")
	}
	if len(item.TextPrimitiveElement.Extension) != 1 {
		t.Fatalf("expected 1 extension on _text, got %d", len(item.TextPrimitiveElement.Extension))
	}

	data, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var actual map[string]interface{}
	if err := json.Unmarshal(data, &actual); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}
	if _, ok := actual["_text"]; !ok {
		t.Errorf("expected _text in marshalled item, got %s", data)
	}
}

// A repeating primitive has an array sibling, positionally matched to the values, with null for
// the values that carry no id or extensions.
func TestPrimitiveElement_RepeatingPrimitiveKeepsPositions(t *testing.T) {
	input := []byte(`{
	  "status": "active",
	  "derivedFrom": ["Questionnaire/a", "Questionnaire/b"],
	  "_derivedFrom": [null, {"id": "second"}]
	}`)

	var questionnaire Questionnaire
	if err := json.Unmarshal(input, &questionnaire); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if len(questionnaire.DerivedFromPrimitiveElement) != 2 {
		t.Fatalf("expected 2 entries in _derivedFrom, got %d", len(questionnaire.DerivedFromPrimitiveElement))
	}
	if questionnaire.DerivedFromPrimitiveElement[0] != nil {
		t.Errorf("expected first entry to be null, got %v", questionnaire.DerivedFromPrimitiveElement[0])
	}
	if questionnaire.DerivedFromPrimitiveElement[1] == nil || *questionnaire.DerivedFromPrimitiveElement[1].Id != "second" {
		t.Errorf("expected second entry to have id 'second', got %v", questionnaire.DerivedFromPrimitiveElement[1])
	}
}
