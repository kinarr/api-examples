package examples

import (
	"testing"
)

func TestEmbedContentBasic(t *testing.T) {
	err := EmbedContentBasic()
	if err != nil {
		t.Errorf("EmbedContentBasic returned an error.")
	}
}

func TestBatchEmbedContents(t *testing.T) {
	err := BatchEmbedContents()
	if err != nil {
		t.Errorf("BatchEmbedContents returned an error.")
	}
}
