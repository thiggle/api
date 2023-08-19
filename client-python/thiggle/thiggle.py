import json

import requests

version = "0.0.5"

class API:
    def __init__(self, api_key):
        self.api_key = api_key
        self.base_url = "https://api.thiggle.com"
        self.headers = {
            'Authorization': f'Bearer {self.api_key}',
            'Content-Type': 'application/json',
            'User-Agent': f'thiggle-client-py/{version}'
        }
    
    def completion(self, 
              prompt,
              model=None,
              max_tokens=None,
              temperature=None,
              top_p=None,
              stop=None,
              presence_penalty=None,
              frequency_penalty=None):
        """
        Generate a completion based on the given prompt.

        :param prompt: The prompt to complete.
        :param model: The model name(s) to use for completion. Can be a string or a list of strings.
        :param max_tokens: Maximum number of tokens in the response.
        :param temperature: Controls randomness in the response. Ranges from 0 to 1.
        :param top_p: Controls the nucleus sampling method. Ranges from 0 to 1.
        :param stop: List of strings that indicates stopping tokens.
        :param presence_penalty: Affects the presence of tokens in the response. Ranges from -2 to 2.
        :param frequency_penalty: Affects the frequency of tokens in the response. Ranges from -2 to 2.
        """

        payload = {
            "prompt": prompt,
            "model": model,
            "max_tokens": max_tokens,
            "temperature": temperature,
            "top_p": top_p,
            "stop": stop,
            "presence_penalty": presence_penalty,
            "frequency_penalty": frequency_penalty
        }

        # Removing None values from the payload
        payload = {k: v for k, v in payload.items() if v is not None}

        response = requests.post(
            f'{self.base_url}/v1/completion',
            headers=self.headers,
            data=json.dumps(payload)
        )

        if response.status_code == 200:
            return response.json()
        else:
            response.raise_for_status()

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
                         pattern, 
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
        :param pattern:  A single regex pattern to match.
        :param max_new_tokens:  The maximum number of tokens to generate. Defaults to 5.
        :param stop_after_match:  Whether to stop generating after a match is found. Defaults to True.
        :param top_k:  The number of highest probability vocabulary tokens to keep for top-k-filtering. Defaults to 50.
        :param top_p:  If set to float < 1, only the most probable tokens with probabilities that add up to top_p or higher are kept for generation. Defaults to 1.0.
        :param temperature:  The value used to modulate the next token probabilities. Defaults to 1.0.
        :param repetition_penalty:  The parameter for repetition penalty. 1.0 means no penalty. See this https://arxiv.org/pdf/1909.05858.pdf for more details. Defaults to 1.0.

        """
        payload = {
            "prompt": prompt,
            "pattern": pattern,
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

