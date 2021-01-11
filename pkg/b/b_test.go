package b

import "testing"

func TestB(t *testing.T) {
	expected := "bV2/c"
	actual := B()
	if actual != expected {
		t.Errorf("got %v but expected %v", actual, expected)
	}
}
