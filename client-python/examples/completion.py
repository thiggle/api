import os
import thiggle as tg

api = tg.API(os.getenv("THIGGLE_DEV_API_KEY"))
api.base_url = "http://localhost:8080"
request = {
    "prompt": "What is the meaning of life?",
    "model": ["gpt-4", "llama-2-70b-chat"],
    "max_tokens": 10,
}
response = api.completion(**request)
print(response)