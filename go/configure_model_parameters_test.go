package examples

import (
	"testing"
)

func TestConfigureModel(t *testing.T) {
	_, err := ConfigureModel()
	if err != nil {
		t.Errorf("ConfigureModel returned an error.")
	}
}
