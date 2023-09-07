import os

import thiggle as tg

# Create an API object with your API key
api = tg.API(os.getenv("THIGGLE_API_KEY"))

prompt = "Thiggle, the specialized and structured LLM API is an acronym for "
schema = "export type Message = { text: string; }"
type_name = "Message"

response = api.typed_completion(prompt, schema, type_name, max_new_tokens=15)

print(response)
