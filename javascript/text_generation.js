import {GoogleGenAI} from '@google/genai';

export async function textGenTextOnlyPrompt() {
  // [START text_gen_text_only_prompt]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const GEMINI_API_KEY = process.env.GEMINI_API_KEY;
  const ai = new GoogleGenAI({apiKey: GEMINI_API_KEY});

  const response = await ai.models.generateContent({
    model: 'gemini-2.0-flash-001',
    contents: 'Write a story about a magic backpack.',
  });
  console.log(response.text);
  // [END text_gen_text_only_prompt]
  return response.text;
}
