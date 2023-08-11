package thiggle

import (
	"context"
	"fmt"
	"os"
	"testing"
)

var client *Client

func TestMain(m *testing.M) {
	client = NewClient(os.Getenv("THIGGLE_DEV_API_KEY"))
	client.BaseURL = "http://localhost:8080"

	os.Exit(m.Run())
}

func TestIntegrationRegexCompletion(t *testing.T) {
	req := &RegexCompletionRequest{
		Prompt: "ReLLM, the best way to get structured data out of LLMs, is an acronym for ",
		Patterns: []string{
			"Re[a-z]+ L[a-z]+ L[a-z]+ M[a-z]+",
		},
		MaxNewTokens: 5,
	}

	resp, err := client.RegexCompletion(context.Background(), req)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	fmt.Println(resp)
}

const palindromeGrammar = `
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
`

func TestIntegrationContextFreeCompletion(t *testing.T) {
	req := &ContextFreeCompletionRequest{
		Prompt:       "Generate a palindrome: ",
		Grammar:      palindromeGrammar,
		MaxNewTokens: 11,
	}

	resp, err := client.ContextFreeCompletion(context.Background(), req)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	fmt.Println(resp)
}

func TestIntegrationCategorize(t *testing.T) {
	req := &CategorizeRequest{
		Prompt: "What animals have four legs?",
		Categories: []string{
			"dog",
			"fish",
			"snake",
			"cow",
		},
		AllowMultipleClasses: NewBool(true),
	}

	resp, err := client.Categorize(context.Background(), req)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if len(resp.Choices) != 2 {
		t.Fatal("expected 2 choices")
	}

}
