package to

import (
	"testing"
)

func TestValue(t *testing.T) {
	t.Run("nil string", func(t *testing.T) {
		expected := ""
		actual := Value((*string)(nil))
		if actual != expected {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})
	t.Run("string", func(t *testing.T) {
		expected := "hello"
		actual := Value(&expected)
		if actual != expected {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})
}
