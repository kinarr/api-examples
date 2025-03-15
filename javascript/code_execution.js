/**
 * @license
 * Copyright 2025 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import {GoogleGenAI} from '@google/genai';

export async function codeExecutionBasic() {
  // [START code_execution_basic]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});

  const response = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents:
      'What is the sum of the first 50 prime numbers? Generate and run code for the calculation, and make sure you get all 50.',
  });

  // Each part may contain text, executable code, or an execution result.
  for (const part of response.candidates[0].content.parts) {
    console.log(part);
    console.log('\n');
  }

  console.log('-'.repeat(80));
  // The `.text` accessor concatenates the parts into a markdown-formatted text.
  console.log('\n', response.text);
  // [END code_execution_basic]

  return {
    parts: response.candidates[0].content.parts,
    text: response.text,
  };
}

/**
 * codeExecutionRequestOverride:
 * Similar to codeExecutionBasic, but illustrates an override in the request.
 * Prints out the executable code and its execution result.
 */
export async function codeExecutionRequestOverride() {
  // [START code_execution_request_override]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});

  const response = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents:
      'What is the sum of the first 50 prime numbers? Generate and run code for the calculation, and make sure you get all 50.',
    config: {
      tools: [{codeExecution: {}}],
    },
  });

  console.log('\n', response.executableCode);
  console.log('\n', response.codeExecutionResult);
  // [END code_execution_request_override]

  return {
    executableCode: response.executableCode,
    codeExecutionResult: response.codeExecutionResult,
  };
}
