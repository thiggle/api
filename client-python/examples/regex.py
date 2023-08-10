

import os

import thiggle as tg

# Create an API object with your API key
api = tg.API(os.getenv("THIGGLE_API_KEY"))

prompt = "Thiggle, the specialized and structured LLM API is an acronym for "
patterns = ["T[a-z]+ H[a-z]+ I[a-z]+ G[a-z]+ G[a-z]+ L[a-z]+ E[a-z]+"]

response = api.regex_completion(prompt, patterns, max_new_tokens=15)

print(response)