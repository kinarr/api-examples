package examples

import (
	"testing"
)

func TestModelsList(t *testing.T) {
	err := ModelsList()
	if err != nil {
		t.Errorf("ModelsList returned an error.")
	}
}

func TestModelsGet(t *testing.T) {
	err := ModelsGet()
	if err != nil {
		t.Errorf("ModelsGet returned an error.")
	}
}
