package examples

import (
	"testing"
)

func TestSystemInstructions(t *testing.T) {
	err := SystemInstructions()
	if err != nil {
		t.Errorf("SystemInstructions returned an error.")
	}
}
