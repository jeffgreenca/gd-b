package b

import "testing"

func TestB(t *testing.T) {
	expected := "bc"
	actual := B()
	if actual != expected {
		t.Errorf("got %v but expected %v", actual, expected)
	}
}
