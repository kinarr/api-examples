package examples

import (
	"testing"
)

func TestTextGenTextOnlyPrompt(t *testing.T) {
	_, err := TextGenTextOnlyPrompt()
	if err != nil {
		t.Errorf("TextGenTextOnlyPrompt returned an error.")
	}
}

func TestTextGenTextOnlyPromptStreaming(t *testing.T) {
	err := TextGenTextOnlyPromptStreaming()
	if err != nil {
		t.Errorf("TextGenTextOnlyPromptStreaming returned an error.")
	}
}

func TestTextGenMultimodalOneImagePrompt(t *testing.T) {
	_, err := TextGenMultimodalOneImagePrompt()
	if err != nil {
		t.Errorf("TextGenMultimodalOneImagePrompt returned an error.")
	}
}

func TestTextGenMultimodalOneImagePromptStreaming(t *testing.T) {
	err := TextGenMultimodalOneImagePromptStreaming()
	if err != nil {
		t.Errorf("TextGenMultimodalOneImagePromptStreaming returned an error.")
	}
}

func TestTextGenMultimodalMultiImagePrompt(t *testing.T) {
	_, err := TextGenMultimodalMultiImagePrompt()
	if err != nil {
		t.Errorf("TextGenMultimodalMultiImagePrompt returned an error.")
	}
}

func TestTextGenMultimodalMultiImagePromptStreaming(t *testing.T) {
	err := TextGenMultimodalMultiImagePromptStreaming()
	if err != nil {
		t.Errorf("TextGenMultimodalMultiImagePromptStreaming returned an error.")
	}
}

func TestTextGenMultimodalAudio(t *testing.T) {
	_, err := TextGenMultimodalAudio()
	if err != nil {
		t.Errorf("TextGenMultimodalAudio returned an error.")
	}
}

func TestTextGenMultimodalAudioStreaming(t *testing.T) {
	err := TextGenMultimodalAudioStreaming()
	if err != nil {
		t.Errorf("TextGenMultimodalAudioStreaming returned an error.")
	}
}

func TestTextGenMultimodalVideoPrompt(t *testing.T) {
	_, err := TextGenMultimodalVideoPrompt()
	if err != nil {
		t.Errorf("TextGenMultimodalVideoPrompt returned an error.")
	}
}

func TestTextGenMultimodalVideoPromptStreaming(t *testing.T) {
	err := TextGenMultimodalVideoPromptStreaming()
	if err != nil {
		t.Errorf("TextGenMultimodalVideoPromptStreaming returned an error.")
	}
}

func TestTextGenMultimodalPdf(t *testing.T) {
	_, err := TextGenMultimodalPdf()
	if err != nil {
		t.Errorf("TextGenMultimodalPdf returned an error.")
	}
}

func TestTextGenMultimodalPdfStreaming(t *testing.T) {
	err := TextGenMultimodalPdfStreaming()
	if err != nil {
		t.Errorf("TextGenMultimodalPdfStreaming returned an error.")
	}
}
