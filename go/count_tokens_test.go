package examples

import (
	"testing"
)

func TestTokensContextWindow(t *testing.T) {
	err := TokensContextWindow()
	if err != nil {
		t.Errorf("TokensContextWindow returned an error.")
	}
}

func TestTokensTextOnly(t *testing.T) {
	err := TokensTextOnly()
	if err != nil {
		t.Errorf("TokensTextOnly returned an error.")
	}
}

func TestTokensMultimodalImageInline(t *testing.T) {
	err := TokensMultimodalImageInline()
	if err != nil {
		t.Errorf("TokensMultimodalImageInline returned an error.")
	}
}

func TestTokensMultimodalVideoAudioInline(t *testing.T) {
	err := TokensMultimodalVideoAudioInline()
	if err != nil {
		t.Errorf("TokensMultimodalVideoAudioInline returned an error.")
	}
}

func TestTokensMultimodalPdfInline(t *testing.T) {
	err := TokensMultimodalPdfInline()
	if err != nil {
		t.Errorf("TokensMultimodalPdfInline returned an error.")
	}
}

func TestTokensCachedContent(t *testing.T) {
	err := TokensCachedContent()
	if err != nil {
		t.Errorf("TokensCachedContent returned an error.")
	}
}
