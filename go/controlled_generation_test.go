package examples

import (
	"testing"
)

func TestJsonControlledGeneration(t *testing.T) {
	_, err := JsonControlledGeneration()
	if err != nil {
		t.Errorf("JsonControlledGeneration returned an error.")
	}
}

func TestJsonNoSchema(t *testing.T) {
	_, err := JsonNoSchema()
	if err != nil {
		t.Errorf("JsonNoSchema returned an error.")
	}
}

func TestJsonEnum(t *testing.T) {
	_, err := JsonEnum()
	if err != nil {
		t.Errorf("JsonEnum returned an error.")
	}
}

func TestEnumInJson(t *testing.T) {
	_, err := EnumInJson()
	if err != nil {
		t.Errorf("EnumInJson returned an error.")
	}
}

func TestJsonEnumRaw(t *testing.T) {
	_, err := JsonEnumRaw()
	if err != nil {
		t.Errorf("JsonEnumRaw returned an error.")
	}
}

func TestXEnum(t *testing.T) {
	_, err := XEnum()
	if err != nil {
		t.Errorf("XEnum returned an error.")
	}
}

func TestXEnumRaw(t *testing.T) {
	_, err := XEnumRaw()
	if err != nil {
		t.Errorf("XEnumRaw returned an error.")
	}
}
