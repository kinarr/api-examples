package examples

import (
	"testing"
)

func TestSystemInstruction(t *testing.T) {
	err := SystemInstruction()
	if err != nil {
		t.Errorf("SystemInstruction returned an error.")
	}
}
