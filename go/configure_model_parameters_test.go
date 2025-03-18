package examples

import (
	"testing"
)

func TestConfigureModelParameters(t *testing.T) {
	_, err := ConfigureModelParameters()
	if err != nil {
		t.Errorf("ConfigureModelParameters returned an error.")
	}
}
