package examples

import (
	"testing"
)

func TestFunctionCalling(t *testing.T) {
	err := FunctionCalling()
	if err != nil {
		t.Errorf("FunctionCalling returned an error.")
	}
}
