import os

import thiggle as tg

# Create an API object with your API key
api = tg.API(os.getenv("THIGGLE_API_KEY"))

# Call the categorize method with your prompt and categories
response = api.categorize("What animal barks?", ["dog", "cat", "bird", "fish"])
# Print the response
print(response)