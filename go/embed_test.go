package examples

import (
	"testing"
)

func TestEmbedContent(t *testing.T) {
	err := EmbedContent()
	if err != nil {
		t.Errorf("EmbedContent returned an error.")
	}
}

func TestBatchEmbedContents(t *testing.T) {
	err := BatchEmbedContents()
	if err != nil {
		t.Errorf("BatchEmbedContents returned an error.")
	}
}
