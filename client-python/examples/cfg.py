import os
import thiggle as tg

# Create an API object with your API key
api = tg.API(os.getenv("THIGGLE_DEV_API_KEY"))
grammar = """
start: palindrome
palindrome: 	letter
		| "a" palindrome "a"
		| "b" palindrome "b"
		| "c" palindrome "c"
		| "d" palindrome "d"
		| "e" palindrome "e"
		| "f" palindrome "f"
		| "g" palindrome "g"
		| "h" palindrome "h"
		| "i" palindrome "i"
		| "j" palindrome "j"
		| "k" palindrome "k"
		| "l" palindrome "l"
		| "m" palindrome "m"
		| "n" palindrome "n"
		| "p" palindrome "p"
		| "q" palindrome "q"
		| "r" palindrome "r"
		| "s" palindrome "s"
		| "t" palindrome "t"
		| "u" palindrome "u"
		| "v" palindrome "v"
		| "w" palindrome "w"
		| "x" palindrome "x"
		| "y" palindrome "y"
		| "z" palindrome "z"

letter: "a".."z"
"""
prompt = "Generate a palindrome: "
response = api.cfg_completion(prompt, grammar, max_new_tokens=15)
print(response)
