import { textGenTextOnlyPrompt } from './text_generation.js';
import assert from 'node:assert';
import { test, describe } from 'node:test';

describe('text_generation', () => {
  test('textGenTextOnlyPrompt', async () => {
    const text = await textGenTextOnlyPrompt();
    assert.ok(text.length > 0);
  });
});

