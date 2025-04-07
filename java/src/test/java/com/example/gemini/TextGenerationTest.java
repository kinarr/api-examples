/*
 * Copyright 2025 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.example.gemini;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertFalse;


public class TextGenerationTest {
    @Test
    public void test_TextGenTextOnlyPrompt() {
        String result = assertDoesNotThrow(TextGeneration::TextGenTextOnlyPrompt,
                "TextGenTextOnlyPrompt returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_TextGenTextOnlyPromptStreaming() {
        String result = assertDoesNotThrow(TextGeneration::TextGenTextOnlyPromptStreaming,
                "TextGenTextOnlyPromptStreaming returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_TextGenMultimodalOneImagePrompt() {
        String result = assertDoesNotThrow(TextGeneration::TextGenMultimodalOneImagePrompt,
                "TextGenMultimodalOneImagePrompt returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_TextGenMultimodalOneImagePromptStreaming() {
        String result = assertDoesNotThrow(TextGeneration::TextGenMultimodalOneImagePromptStreaming,
                "TextGenMultimodalOneImagePromptStreaming returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_TextGenMultimodalMultiImagePrompt() {
        String result = assertDoesNotThrow(TextGeneration::TextGenMultimodalMultiImagePrompt,
                "TextGenMultimodalMultiImagePrompt returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_TextGenMultimodalMultiImagePromptStreaming() {
        String result = assertDoesNotThrow(TextGeneration::TextGenMultimodalMultiImagePromptStreaming,
                "TextGenMultimodalMultiImagePromptStreaming returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_TextGenMultimodalAudio() {
        String result = assertDoesNotThrow(TextGeneration::TextGenMultimodalAudio,
                "TextGenMultimodalAudio returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_TextGenMultimodalAudioStreaming() {
        String result = assertDoesNotThrow(TextGeneration::TextGenMultimodalAudioStreaming,
                "TextGenMultimodalAudioStreaming returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_TextGenMultimodalVideoPrompt() {
        String result = assertDoesNotThrow(TextGeneration::TextGenMultimodalVideoPrompt,
                "TextGenMultimodalVideoPrompt returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_TextGenMultimodalVideoPromptStreaming() {
        String result = assertDoesNotThrow(TextGeneration::TextGenMultimodalVideoPromptStreaming,
                "TextGenMultimodalVideoPromptStreaming returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_TextGenMultimodalPdf() {
        String result = assertDoesNotThrow(TextGeneration::TextGenMultimodalPdf,
                "TextGenMultimodalPdf returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_TextGenMultimodalPdfStreaming() {
        String result = assertDoesNotThrow(TextGeneration::TextGenMultimodalPdfStreaming,
                "TextGenMultimodalPdfStreaming returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }
}
