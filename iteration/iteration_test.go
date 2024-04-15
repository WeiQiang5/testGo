package iteration

import "testing"

func TestIteration(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("repeated: '%s' expected '%s'", repeated, expected)
	}
}
