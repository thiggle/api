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

    def categorize(self, 
                   prompt, 
                   categories,
                   allow_null_category=False,
                   allow_multiple_classes=False):
        """
        Categorize a prompt into one or more categories.

        :param prompt:  The prompt to categorize.
        :param categories:  The categories to categorize into.
        :param allow_null_category:  Whether to allow the model to return no category. Defaults to False.
        :param allow_multiple_classes:  Whether to allow the model to return multiple categories. Defaults to False.
        """

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

    def regex_completion(self, 
                         prompt, 
                         patterns, 
                         max_new_tokens=5, 
                         stop_after_match=True,
                         temperature=1.0,
                         top_p=1.0,
                         top_k=50,
                         repetition_penalty=1.0,
                         ):
        
        """
        Generate LLM completions that match regex patterns.

        :param prompt:  The prompt to generate from.
        :param patterns:  A single regex pattern or list of regex patterns to match.
        :param max_new_tokens:  The maximum number of tokens to generate. Defaults to 5.
        :param stop_after_match:  Whether to stop generating after a match is found. Defaults to True.
        :param top_k:  The number of highest probability vocabulary tokens to keep for top-k-filtering. Defaults to 50.
        :param top_p:  If set to float < 1, only the most probable tokens with probabilities that add up to top_p or higher are kept for generation. Defaults to 1.0.
        :param temperature:  The value used to modulate the next token probabilities. Defaults to 1.0.
        :param repetition_penalty:  The parameter for repetition penalty. 1.0 means no penalty. See this https://arxiv.org/pdf/1909.05858.pdf for more details. Defaults to 1.0.

        """
        payload = {
            "prompt": prompt,
            "patterns": patterns,
            "max_new_tokens": max_new_tokens,
            "stop_after_match": stop_after_match,
            "temperature": temperature,
            "top_p": top_p,
            "top_k": top_k,
            "repetition_penalty": repetition_penalty
        }

        response = requests.post(
            f'{self.base_url}/v1/completion/regex',
            headers=self.headers,
            data=json.dumps(payload)
        )

        if response.status_code == 200:
            return response.json()
        else:
            response.raise_for_status()

    def cfg_completion(self,
                       prompt,
                       grammar,
                       max_new_tokens=5,
                       stop_after_match=True,
                       temperature=1.0,
                       top_p=1.0,
                       top_k=50,
                       repetition_penalty=1.0,
                       ):
        """
        Generate LLM completions that match CFG grammar rules.

        :param prompt:  The prompt to generate from.
        :param grammar:  A single CFG grammar rule to match.
        :param max_new_tokens:  The maximum number of tokens to generate. Defaults to 5.
        :param stop_after_match:  Whether to stop generating after a match is found. Defaults to True.
        :param top_k:  The number of highest probability vocabulary tokens to keep for top-k-filtering. Defaults to 50.
        :param top_p:  If set to float < 1, only the most probable tokens with probabilities that add up to top_p or higher are kept for generation. Defaults to 1.0.
        :param temperature:  The value used to modulate the next token probabilities. Defaults to 1.0.
        :param repetition_penalty:  The parameter for repetition penalty. 1.0 means no penalty. See this https://arxiv.org/pdf/1909.05858.pdf for more details. Defaults to 1.0.

        """
        payload = {
            "prompt": prompt,
            "grammar": grammar,
            "max_new_tokens": max_new_tokens,
            "stop_after_match": stop_after_match,
            "temperature": temperature,
            "top_p": top_p,
            "top_k": top_k,
            "repetition_penalty": repetition_penalty
        }

        response = requests.post(
            f'{self.base_url}/v1/completion/cfg',
            headers=self.headers,
            data=json.dumps(payload)
        )

        if response.status_code == 200:
            return response.json()
        else:
            response.raise_for_status()

