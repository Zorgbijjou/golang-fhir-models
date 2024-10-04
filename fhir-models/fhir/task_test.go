package fhir

import (
	"encoding/json"
	"testing"
)

func TestTaskInput_ValuePolymorphism(t *testing.T) {
	// Expected JSON:
	// {
	//	 "type": {},
	//	 "valueInteger": 5
	// }

	expected := 5
	taskInput := TaskInput{ValueInteger: &expected}
	data, _ := json.MarshalIndent(taskInput, "", "  ")
	actual := map[string]interface{}{}
	_ = json.Unmarshal(data, &actual)
	if len(actual) != 2 {
		t.Errorf("expected 2 fields, got %d", len(actual))
	}
	if int(actual["valueInteger"].(float64)) != expected {
		t.Errorf("expected %d (%T), got %v (%T)", expected, expected, actual["valueInteger"], actual["valueInteger"])
	}
}
