package examples

import (
	"testing"
)

func TestSafetySettings(t *testing.T) {
	err := SafetySettings()
	if err != nil {
		t.Errorf("SafetySettings returned an error.")
	}
}

func TestSafetySettingsMulti(t *testing.T) {
	err := SafetySettingsMulti()
	if err != nil {
		t.Errorf("SafetySettingsMulti returned an error.")
	}
}
