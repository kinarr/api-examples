# -*- coding: utf-8 -*-
# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
from absl.testing import absltest

class UnitTests(absltest.TestCase):
    def test_code_execution_basic(self):
        # [START code_execution_basic]
        from google import genai
        from google.genai import types
        client = genai.Client()
        response = client.models.generate_content(
            model="gemini-2.0-flash",
            contents=(
                "What is the sum of the first 50 prime numbers? "
                "Generate and run code for the calculation, and make sure you get all 50."
            ),
            config=types.GenerateContentConfig(
                tools=[types.Tool(code_execution=types.ToolCodeExecution())],
            )
        )
        # Each part may contain text, executable code, or an execution result.
        for part in response.candidates[0].content.parts:
            print(part, "\n")

        print("-" * 80)
        # The .text accessor concatenates the parts into a markdown-formatted text.
        print("\n", response.executable_code)
        print("\n", response.code_execution_result)
        # [END code_execution_basic]

        # [START code_execution_basic_return]
        # Expected output:
        # video_metadata=None thought=None code_execution_result=None executable_code=None file_data=None function_call=None function_response=None inline_data=None text="Okay, I understand. I need to calculate the sum of the first 50 prime numbers. I'll use a Python code block to generate the prime numbers and calculate their sum. I'll make sure the code generates the first 50 primes correctly before summing them.\n\n"

        # video_metadata=None thought=None code_execution_result=None executable_code=ExecutableCode(code='def is_prime(n):\n  """Returns True if n is a prime number, False otherwise."""\n  if n <= 1:\n    return False\n  for i in range(2, int(n**0.5) + 1):\n    if n % i == 0:\n      return False\n  return True\n\nprimes = []\nnum = 2\nwhile len(primes) < 50:\n  if is_prime(num):\n    primes.append(num)\n  num += 1\n\nprint(f\'{primes=}\')\nprint(f\'{sum(primes)=}\')\n', language=<Language.PYTHON: 'PYTHON'>) file_data=None function_call=None function_response=None inline_data=None text=None

        # video_metadata=None thought=None code_execution_result=CodeExecutionResult(outcome=<Outcome.OUTCOME_OK: 'OUTCOME_OK'>, output='primes=[2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229]\nsum(primes)=5117\n') executable_code=None file_data=None function_call=None function_response=None inline_data=None text=None

        # video_metadata=None thought=None code_execution_result=None executable_code=None file_data=None function_call=None function_response=None inline_data=None text='The sum of the first 50 prime numbers is 5117.\n'

        # --------------------------------------------------------------------------------

        # def is_prime(n):
        #   """Returns True if n is a prime number, False otherwise."""
        #   if n <= 1:
        #     return False
        #   for i in range(2, int(n**0.5) + 1):
        #     if n % i == 0:
        #       return False
        #   return True

        # primes = []
        # num = 2
        # while len(primes) < 50:
        #   if is_prime(num):
        #     primes.append(num)
        #   num += 1

        # print(f'{primes=}')
        # print(f'{sum(primes)=}')


        # primes=[2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229]
        # sum(primes)=5117
        # [END code_execution_basic_return]

    def test_code_execution_request_override(self):
        # [START code_execution_request_override]
        from google import genai
        from google.genai import types
        client = genai.Client()
        response = client.models.generate_content(
            model="gemini-2.0-flash",
            contents=(
                "What is the sum of the first 50 prime numbers? "
                "Generate and run code for the calculation, and make sure you get all 50."
            ),
            config=types.GenerateContentConfig(
                tools=[types.Tool(code_execution=types.ToolCodeExecution())],
            )
        )
        print(response.executable_code)
        print(response.code_execution_result)
        # [END code_execution_request_override]

        # [START code_execution_request_override_return]
        # Expected output:
        # def is_prime(n):
        #     """Check if a number is prime."""
        #     if n <= 1:
        #         return False
        #     for i in range(2, int(n**0.5) + 1):
        #         if n % i == 0:
        #             return False
        #     return True

        # primes = []
        # num = 2
        # while len(primes) < 50:
        #     if is_prime(num):
        #         primes.append(num)
        #     num += 1

        # sum_of_primes = sum(primes)
        # print(f'{primes=}')
        # print(f'{sum_of_primes=}')

        # primes=[2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229]
        # sum_of_primes=5117
        # [END code_execution_request_override_return]

    def test_code_execution_chat(self):
        # [START code_execution_chat]
        from google import genai
        from google.genai import types
        client = genai.Client()
        chat = client.chats.create(
            model="gemini-2.0-flash",
            config=types.GenerateContentConfig(
                tools=[types.Tool(code_execution=types.ToolCodeExecution())],
            )
        )
        # First, a simple chat message.
        response = chat.send_message(message='Can you print "Hello world!"?')
        # Then, a code-execution request.
        response = chat.send_message(
            message=(
                "What is the sum of the first 50 prime numbers? "
                "Generate and run code for the calculation, and make sure you get all 50."
            )
        )
        print(response.executable_code)
        print(response.code_execution_result)
        # [END code_execution_chat]

        # [START code_execution_chat_return]
        # Expected output:
        # def is_prime(n):
        #   """Returns True if n is a prime number, False otherwise."""
        #   if n <= 1:
        #     return False
        #   for i in range(2, int(n**0.5) + 1):
        #     if n % i == 0:
        #       return False
        #   return True

        # primes = []
        # num = 2
        # while len(primes) < 50:
        #   if is_prime(num):
        #     primes.append(num)
        #   num += 1

        # print(f'{primes=}')

        # sum_of_primes = sum(primes)
        # print(f'{sum_of_primes=}')

        # primes=[2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229]
        # sum_of_primes=5117
        # [END code_execution_chat_return]


if __name__ == "__main__":
    absltest.main()
