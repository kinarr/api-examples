package examples

import (
	"testing"
)

func TestGenerateContentTextOnly(t *testing.T) {
	_, err := GenerateContentTextOnly()
	if err != nil {
		t.Errorf("GenerateContentTextOnly returned an error.")
	}
}
