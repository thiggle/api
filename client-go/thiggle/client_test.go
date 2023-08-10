package thiggle

import (
	"context"
	"os"
	"testing"
)

func TestIntegrationRegexCompletion(t *testing.T) {
	client := NewClient(os.Getenv("THIGGLE_API_KEY"))
	client.BaseURL = "http://localhost:8080"

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

	if len(resp.Completion) != 2 {
		t.Fatal("expected 2 choices")
	}
}

func TestIntegrationCategorize(t *testing.T) {
	client := NewClient(os.Getenv("THIGGLE_API_KEY"))
	client.BaseURL = "http://localhost:8080"

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
