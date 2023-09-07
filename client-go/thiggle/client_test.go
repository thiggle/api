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

func TestTypedCompletion(t *testing.T) {
	req := &TypedCompletionRequest{
		Prompt:   "The best way to get structured data out of LLMs is to use ReLLM, which is an acronym for ",
		Schema:   "export type Message = { text: string }",
		TypeName: "Message",
	}

	resp, err := client.TypedCompletion(context.Background(), req)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	fmt.Println(resp)

}

func TestIntegrationCompletion(t *testing.T) {
	req := &CompletionRequest{
		Prompt:      "the input is 10 tokens long, really!",
		Model:       MustNewStringOrSlice([]string{"llama-2-70b-chat", "text-davinci-003"}),
		MaxTokens:   10,
		Temperature: 0.5,
	}
	resp, err := client.Completion(context.Background(), req)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	fmt.Printf("id=%s object=%s created=%s\n", resp.ID, resp.Object, resp.Created)
	if resp.Usage != nil {
		fmt.Printf("usage.completion_tokens=%d usage.prediction_time_ms=%d usage.total_tokens=%d\n", resp.Usage.CompletionTokens, resp.Usage.PredictionTimeMs, resp.Usage.TotalTokens)
	}
	for _, choice := range resp.Choices {
		fmt.Println("---")
		fmt.Printf("model=%s text=%s created=%d finish_reason=%s\n", choice.Model, choice.Text, choice.Created, choice.FinishReason)
		if choice.Usage != nil {
			fmt.Printf("usage.completion_tokens=%d usage.prediction_time_ms=%d usage.total_tokens=%d\n", choice.Usage.CompletionTokens, choice.Usage.PredictionTimeMs, choice.Usage.TotalTokens)
		}
		fmt.Println("---")
	}

}

func TestIntegrationRegexCompletion(t *testing.T) {
	req := &RegexCompletionRequest{
		Prompt:       "ReLLM, the best way to get structured data out of LLMs, is an acronym for ",
		Pattern:      "Re[a-z]+ L[a-z]+ L[a-z]+ M[a-z]+",
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
