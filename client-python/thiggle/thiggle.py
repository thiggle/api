import json

import requests

version = "0.0.1"

class API:
    def __init__(self, api_key):
        self.api_key = api_key
        self.base_url = "https://api.thiggle.com"
        self.headers = {
            'Authorization': f'Bearer {self.api_key}',
            'Content-Type': 'application/json',
            'User-Agent': f'thiggle-client-py/{version}'
        }

    def categorize(self, prompt, categories, allow_null_category=False, allow_multiple_classes=False):
        payload = {
            "prompt": prompt,
            "categories": categories,
            "allow_null_category": allow_null_category,
            "allow_multiple_classes": allow_multiple_classes
        }

        response = requests.post(
            f'{self.base_url}/v1/categorize',
            headers=self.headers,
            data=json.dumps(payload)
        )

        if response.status_code == 200:
            return response.json()
        else:
            response.raise_for_status()
