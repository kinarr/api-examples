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
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;

public class ControlledGenerationTest {
    @Test
    public void test_JsonControlledGeneration() {
        String result = assertDoesNotThrow(ControlledGeneration::JsonControlledGeneration,
                "JsonControlledGeneration returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_JsonNoSchema() {
        String result = assertDoesNotThrow(ControlledGeneration::JsonNoSchema,
                "JsonNoSchema returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_JsonEnum() {
        String result = assertDoesNotThrow(ControlledGeneration::JsonEnum,
                "JsonEnum returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
        assertTrue(result.trim().contains("Keyboard"), "Response should contain word `Keyboard`");
    }

    @Test
    public void test_EnumInJson() {
        String result = assertDoesNotThrow(ControlledGeneration::EnumInJson,
                "EnumInJson returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_XEnum() {
        String result = assertDoesNotThrow(ControlledGeneration::XEnum,
                "XEnum returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }
}
